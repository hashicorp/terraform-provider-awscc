# AWS CloudControl FSx S3 Access Point Attachment Configuration

# Create a VPC for the resources
resource "aws_vpc" "fsx_vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "fsx-vpc"
  }
}

# Create subnets in the VPC
resource "aws_subnet" "fsx_subnet_az1" {
  vpc_id            = aws_vpc.fsx_vpc.id
  cidr_block        = "10.0.3.0/24"
  availability_zone = "us-west-2a"
  tags = {
    Name = "fsx-subnet-az1"
  }
}

resource "aws_subnet" "fsx_subnet_az2" {
  vpc_id            = aws_vpc.fsx_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"
  tags = {
    Name = "fsx-subnet-az2"
  }
}

# Create an internet gateway
resource "aws_internet_gateway" "fsx_igw" {
  vpc_id = aws_vpc.fsx_vpc.id
  tags = {
    Name = "fsx-igw"
  }
}

# Create a route table
resource "aws_route_table" "fsx_route_table" {
  vpc_id = aws_vpc.fsx_vpc.id
  tags = {
    Name = "fsx-route-table"
  }
}

# Create a route to the internet
resource "aws_route" "fsx_internet_route" {
  route_table_id         = aws_route_table.fsx_route_table.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.fsx_igw.id
}

# Associate the route table with the subnets
resource "aws_route_table_association" "fsx_route_assoc_az1" {
  subnet_id      = aws_subnet.fsx_subnet_az1.id
  route_table_id = aws_route_table.fsx_route_table.id
}

resource "aws_route_table_association" "fsx_route_assoc_az2" {
  subnet_id      = aws_subnet.fsx_subnet_az2.id
  route_table_id = aws_route_table.fsx_route_table.id
}

# Create a security group for FSx
resource "aws_security_group" "fsx_sg" {
  name        = "fsx-security-group"
  description = "Security group for FSx"
  vpc_id      = aws_vpc.fsx_vpc.id

  ingress {
    from_port   = 2049
    to_port     = 2049
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]
    description = "Allow NFSv4 traffic"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "fsx-security-group"
  }
}

# Create FSx OpenZFS file system with MULTI_AZ_1 deployment type
resource "aws_fsx_openzfs_file_system" "example" {
  subnet_ids          = [aws_subnet.fsx_subnet_az1.id, aws_subnet.fsx_subnet_az2.id]
  preferred_subnet_id = aws_subnet.fsx_subnet_az1.id
  security_group_ids  = [aws_security_group.fsx_sg.id]
  deployment_type     = "MULTI_AZ_1"
  storage_capacity    = 64
  throughput_capacity = 160 # Minimum for MULTI_AZ_1

  root_volume_configuration {
    data_compression_type = "NONE"
    nfs_exports {
      client_configurations {
        clients = "10.0.0.0/16"
        options = ["rw", "async"]
      }
    }
  }

  tags = {
    Name = "example-openzfs"
  }
}

# Create FSx OpenZFS volume
resource "aws_fsx_openzfs_volume" "example" {
  name                             = "example-volume"
  parent_volume_id                 = aws_fsx_openzfs_file_system.example.root_volume_id
  volume_type                      = "OPENZFS"
  storage_capacity_reservation_gib = 32
  storage_capacity_quota_gib       = 32
  data_compression_type            = "NONE"

  nfs_exports {
    client_configurations {
      clients = "10.0.0.0/16"
      options = ["rw", "async"]
    }
  }

  tags = {
    Name = "example-volume"
  }
}

# Example FSx OpenZFS S3 Access Point Attachment
resource "awscc_fsx_s3_access_point_attachment" "example" {
  name = "example-s3-access-point"
  type = "OPENZFS"

  open_zfs_configuration = {
    file_system_identity = {
      type = "POSIX"
      posix_user = {
        uid = 1000
        gid = 1000
      }
    }
    volume_id = aws_fsx_openzfs_volume.example.id
  }

  s3_access_point = {
    vpc_configuration = {
      vpc_id = aws_vpc.fsx_vpc.id
    }
  }
}
