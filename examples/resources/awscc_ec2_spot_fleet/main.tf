# Get Amazon Linux 2023 AMI
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["al2023-ami-*-x86_64"]
  }
}

# Create a VPC and subnet for the spot fleet instances
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Name"
    value = "spot-fleet-vpc"
  }]
}

resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Name"
    value = "spot-fleet-igw"
  }]
}

resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

resource "awscc_ec2_route_table" "example" {
  vpc_id = awscc_ec2_vpc.example.id

  tags = [{
    key   = "Name"
    value = "spot-fleet-rt"
  }]
}

resource "awscc_ec2_route" "example" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id         = awscc_ec2_route_table.example.id
  gateway_id             = awscc_ec2_internet_gateway.example.id
  depends_on             = [awscc_ec2_vpc_gateway_attachment.example]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true
  tags = [{
    key   = "Name"
    value = "spot-fleet-subnet"
  }]
}

resource "awscc_ec2_subnet_route_table_association" "example" {
  subnet_id      = awscc_ec2_subnet.example.id
  route_table_id = awscc_ec2_route_table.example.id
}

# Create security group
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for spot fleet"
  vpc_id            = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Name"
    value = "spot-fleet-sg"
  }]
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 80
      to_port     = 80
      cidr_ip     = "0.0.0.0/0"
    },
    {
      ip_protocol = "tcp"
      from_port   = 443
      to_port     = 443
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  security_group_egress = [
    {
      ip_protocol = "-1"
      from_port   = -1
      to_port     = -1
      cidr_ip     = "0.0.0.0/0"
    }
  ]
}

# Create IAM role for Spot Fleet
data "aws_iam_policy_document" "spot_fleet_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["spotfleet.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "spot_fleet" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.spot_fleet_assume_role.json))
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AmazonEC2SpotFleetTaggingRole"]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Spot Fleet request
resource "awscc_ec2_spot_fleet" "example" {
  spot_fleet_request_config_data = {
    iam_fleet_role  = awscc_iam_role.spot_fleet.arn
    target_capacity = 1

    # Using launch specifications
    launch_specifications = [
      {
        instance_type     = "t3.micro"
        image_id          = data.aws_ami.amazon_linux.id
        subnet_id         = awscc_ec2_subnet.example.id
        weighted_capacity = 1

        security_groups = [
          {
            group_id = awscc_ec2_security_group.example.id
          }
        ]

        tag_specifications = [
          {
            resource_type = "instance"
            tags = [
              {
                key   = "Name"
                value = "spot-fleet-instance"
              },
              {
                key   = "Modified By"
                value = "AWSCC"
              }
            ]
          }
        ]
      }
    ]

    tag_specifications = [
      {
        resource_type = "spot-fleet-request"
        tags = [
          {
            key   = "Name"
            value = "example-spot-fleet"
          },
          {
            key   = "Modified By"
            value = "AWSCC"
          }
        ]
      }
    ]

    type                           = "maintain"
    allocation_strategy            = "lowestPrice"
    instance_interruption_behavior = "terminate"
  }
}