
resource "aws_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "example-mediaconnect-vpc"
  }
}

resource "aws_subnet" "example" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "example-mediaconnect-subnet"
  }
}

resource "aws_security_group" "example" {
  name_prefix = "mediaconnect-router-"
  vpc_id      = aws_vpc.example.id

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "example-mediaconnect-sg"
  }
}

resource "awscc_mediaconnect_router_network_interface" "example" {
  name        = "example-router-eni"
  region_name = "us-west-2"

  configuration = {
    vpc = {
      subnet_id          = aws_subnet.example.id
      security_group_ids = [aws_security_group.example.id]
    }
  }

  tags = [
    {
      key   = "Name"
      value = "example-router-network-interface"
    }
  ]
}

resource "awscc_mediaconnect_router_output" "example" {
  name              = "example-router-output"
  maximum_bitrate   = 100000000
  routing_scope     = "REGIONAL"
  tier              = "OUTPUT_100"
  region_name       = "us-west-2"
  availability_zone = "us-west-2a"

  configuration = {
    standard = {
      network_interface_arn = awscc_mediaconnect_router_network_interface.example.arn
      protocol              = "RTP"
      protocol_configuration = {
        rtp = {
          destination_address = "225.1.1.1"
          destination_port    = 5004
        }
      }
    }
  }

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-router-output"
    }
  ]
}
