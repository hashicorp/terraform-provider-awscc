data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "datasync-fsx-ontap-example"
  }
}

resource "aws_internet_gateway" "example" {
  vpc_id = aws_vpc.example.id
  tags = {
    Name = "datasync-fsx-ontap-igw"
  }
}

resource "aws_subnet" "example" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = {
    Name = "datasync-fsx-ontap-subnet"
  }
}

resource "aws_route_table" "example" {
  vpc_id = aws_vpc.example.id
  tags = {
    Name = "datasync-fsx-ontap-rt"
  }
}

resource "aws_route_table_association" "example" {
  subnet_id      = aws_subnet.example.id
  route_table_id = aws_route_table.example.id
}

resource "aws_security_group" "fsx_ontap" {
  name        = "datasync-fsx-ontap-sg"
  description = "Security group for FSx ONTAP DataSync location"
  vpc_id      = aws_vpc.example.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [aws_vpc.example.cidr_block]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "datasync-fsx-ontap-sg"
  }
}

resource "aws_fsx_ontap_file_system" "example" {
  storage_capacity    = 1024
  subnet_ids          = [aws_subnet.example.id]
  deployment_type     = "SINGLE_AZ_1"
  throughput_capacity = 128
  preferred_subnet_id = aws_subnet.example.id
  security_group_ids  = [aws_security_group.fsx_ontap.id]
  route_table_ids     = [aws_route_table.example.id]
  fsx_admin_password  = "Password123!"

  tags = {
    Name = "datasync-fsx-ontap-fs"
  }
}

resource "aws_fsx_ontap_storage_virtual_machine" "example" {
  file_system_id = aws_fsx_ontap_file_system.example.id
  name           = "datasyncexample"
}

# DataSync Location FSx ONTAP
resource "awscc_datasync_location_fsx_ontap" "example" {
  security_group_arns         = ["arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:security-group/${aws_security_group.fsx_ontap.id}"]
  storage_virtual_machine_arn = aws_fsx_ontap_storage_virtual_machine.example.arn
  subdirectory                = "/share"
  protocol = {
    nfs = {
      mount_options = {
        version = "NFS3"
      }
    }
  }
  tags = [{
    key   = "Environment"
    value = "Example"
  }]
}