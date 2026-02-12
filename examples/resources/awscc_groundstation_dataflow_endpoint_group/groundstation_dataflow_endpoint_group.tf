data "aws_caller_identity" "current" {}

resource "aws_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "groundstation-vpc"
  }
}

resource "aws_subnet" "example" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "groundstation-subnet"
  }
}

resource "aws_security_group" "example" {
  name_prefix = "groundstation-sg"
  vpc_id      = aws_vpc.example.id

  tags = {
    Name = "groundstation-sg"
  }
}

resource "aws_iam_role" "groundstation_role" {
  name = "groundstation-dataflow-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "groundstation.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "groundstation-dataflow-role"
  }
}

resource "aws_iam_role_policy_attachment" "groundstation_policy" {
  role       = aws_iam_role.groundstation_role.name
  policy_arn = "arn:aws:iam::aws:policy/AWSGroundStationAgentInstancePolicy"
}

resource "awscc_groundstation_dataflow_endpoint_group" "example" {
  endpoint_details = [
    {
      aws_ground_station_agent_endpoint = {
        name           = "example-endpoint-1"
        egress_address = {
          socket_address = {
            name = "192.168.1.100"
            port = 8080
          }
        }
        ingress_address = {
          socket_address = {
            name = "192.168.1.101"
            port_range = {
              minimum = 8080
              maximum = 8085
            }
          }
        }
        agent_status = "SUCCESS"
        audit_results = "HEALTHY"
      }
      security_details = {
        role_arn           = aws_iam_role.groundstation_role.arn
        security_group_ids = [aws_security_group.example.id]
        subnet_ids         = [aws_subnet.example.id]
      }
    }
  ]

  tags = {
    Name = "example-dataflow-endpoint-group"
  }
}