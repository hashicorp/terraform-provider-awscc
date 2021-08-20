terraform {
  required_providers {

    aws = {
      source = "hashicorp/aws"
    }

    awscc = {
      source = "hashicorp/awscc"
    }
  }
}

provider "aws" {}
provider "awscc" {}

resource "awscc_kms_key" "test" {
  key_policy = jsonencode({
    Id = "kms-tf-1"
    Statement = [
      {
        Action = "kms:*"
        Effect = "Allow"
        Principal = {
          AWS = "*"
        }

        Resource = "*"
        Sid      = "Enable IAM User Permissions"
      },
    ]
    Version = "2012-10-17"
  })
}

# resource "awscc_ec2_dhcp_options" "test" {
#   domain_name          = "service.tf"
#   domain_name_servers  = ["127.0.0.1", "10.0.0.2"]
#   ntp_servers          = ["127.0.0.1"]
#   netbios_name_servers = ["127.0.0.1"]
#   netbios_node_type    = 2

#   tags = [
#     {
#       key   = "Name"
#       value = "AWS CC testing"
#     }
#   ]
# }

# resource "awscc_athena_work_group" "test" {
#   name = "awscc-testing"

#   work_group_configuration = {
#     bytes_scanned_cutoff_per_query = 12582912

#     result_configuration = {
#       encryption_configuration = {
#         encryption_option = "SSE_S3"
#       }
#     }
#   }
# }

# resource "awscc_xray_group" "test" {
#   filter_expression = "responsetime > 5"
#   group_name        = "ewbankkit-cc-test"
# }

# data "aws_availability_zones" "available" {
#   state = "available"

#   filter {
#     name   = "opt-in-status"
#     values = ["opt-in-not-required"]
#   }
# }

# resource "aws_iam_role" "test" {
#   name = "ewbankkit-awscc-test"

#   assume_role_policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [{
#     "Action": "sts:AssumeRole",
#     "Principal": {
#       "Service": "lambda.amazonaws.com"
#     },
#     "Effect": "Allow"
#   }]
# }
# EOF
# }

# resource "aws_iam_policy" "test" {
#   name = "ewbankkit-awscc-test"

#   policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Effect": "Allow",
#       "Action": [
#         "kafka:DescribeCluster",
#         "kafka:GetBootstrapBrokers",
#         "ec2:CreateNetworkInterface",
#         "ec2:DeleteNetworkInterface",
#         "ec2:DescribeNetworkInterfaces",
#         "ec2:DescribeSecurityGroups",
#         "ec2:DescribeSubnets",
#         "ec2:DescribeVpcs",
#         "logs:CreateLogGroup",
#         "logs:CreateLogStream",
#         "logs:PutLogEvents"
#       ],
#       "Resource": "*"
#     }
#   ]
# }
# EOF
# }

# resource "aws_iam_policy_attachment" "test" {
#   name       = "ewbankkit-awscc-test"
#   roles      = [aws_iam_role.test.name]
#   policy_arn = aws_iam_policy.test.arn
# }

# resource "aws_vpc" "test" {
#   cidr_block = "192.168.0.0/22"

#   tags = {
#     Name = "ewbankkit-awscc-test"
#   }
# }

# resource "aws_subnet" "test" {
#   count = 2

#   vpc_id            = aws_vpc.test.id
#   cidr_block        = cidrsubnet(aws_vpc.test.cidr_block, 2, count.index)
#   availability_zone = data.aws_availability_zones.available.names[count.index]

#   tags = {
#     Name = "ewbankkit-awscc-test"
#   }
# }

# resource "aws_security_group" "test" {
#   name   = "ewbankkit-awscc-test"
#   vpc_id = aws_vpc.test.id

#   tags = {
#     Name = "ewbankkit-awscc-test"
#   }
# }

# resource "aws_lambda_function" "test" {
#   filename      = "lambdatest.zip"
#   function_name = "ewbankkit-awscc-test"
#   handler       = "exports.example"
#   role          = aws_iam_role.test.arn
#   runtime       = "nodejs12.x"
# }

# resource "awscc_lambda_event_source_mapping" "test" {
#   batch_size        = 100
#   enabled           = true
#   function_name     = aws_lambda_function.test.arn
#   topics            = ["test"]
#   starting_position = "TRIM_HORIZON"

#   self_managed_event_source = {
#     endpoints = {
#       kafka_bootstrap_servers = ["test1:9092", "test2:9092"]
#     }
#   }

#   source_access_configurations = [
#     {
#       type = "VPC_SECURITY_GROUP"
#       uri  = aws_security_group.test.id
#     },
#     {
#       type = "VPC_SUBNET"
#       uri  = "subnet:${aws_subnet.test.0.id}"
#     },
#     {
#       type = "VPC_SUBNET"
#       uri  = "subnet:${aws_subnet.test.1.id}"
#     },
#   ]
# }