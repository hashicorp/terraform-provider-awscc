resource "awscc_iam_role" "lambda_capacity_provider" {
  role_name = "lambda-capacity-provider-role"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
  ]

  policies = [{
    policy_name = "LambdaCapacityProviderPolicy"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [{
        Effect = "Allow"
        Action = [
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcs",
          "ec2:DescribeInstanceTypeOfferings",
          "ec2:DescribeInstanceTypes",
          "ec2:RunInstances",
          "ec2:TerminateInstances",
          "ec2:DescribeInstances",
          "ec2:CreateTags"
        ]
        Resource = "*"
      }]
    })
  }]

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = [{
    key   = "Name"
    value = "lambda-capacity-provider-vpc"
  }]
}

resource "awscc_ec2_subnet" "example_1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-1a"

  tags = [{
    key   = "Name"
    value = "lambda-capacity-provider-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "example_2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-east-1b"

  tags = [{
    key   = "Name"
    value = "lambda-capacity-provider-subnet-2"
  }]
}

resource "awscc_ec2_security_group" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  group_name        = "lambda-capacity-provider-sg"
  group_description = "Security group for Lambda capacity provider"

  tags = [{
    key   = "Name"
    value = "lambda-capacity-provider-sg"
  }]
}

resource "awscc_lambda_capacity_provider" "example" {
  capacity_provider_name = "example-capacity-provider"

  vpc_config = {
    subnet_ids         = [awscc_ec2_subnet.example_1.id, awscc_ec2_subnet.example_2.id]
    security_group_ids = [awscc_ec2_security_group.example.id]
  }

  permissions_config = {
    capacity_provider_operator_role_arn = awscc_iam_role.lambda_capacity_provider.arn
  }

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
