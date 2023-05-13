# resource "awscc_iam_role" "main" {
#   description = "AWS IAM role for lambda function"
#   assume_role_policy_document = jsonencode({
#     Version = "2012-10-17"
#     Statement = [
#       {
#         Action = "sts:AssumeRole"
#         Effect = "Allow"
#         Sid    = ""
#         Principal = {
#           Service = "lambda.amazonaws.com"
#         }
#       },
#     ]
#   })
# }

# data "archive_file" "main" {
#   type        = "zip"
#   source_file = "main.py"
#   output_path = "lambda_function_payload.zip"
# }

# # EFS file system
# resource "awscc_efs_file_system" "efs_for_lambda" {
#   file_system_tags = [{
#     key   = "Name"
#     value = "efs_for_lambda"
#   }]
# }

# resource "awscc_ec2_vpc" "main" {
#   cidr_block = "10.0.0.0/16"
# }

# resource "awscc_ec2_subnet" "main" {
#   vpc_id     = awscc_ec2_vpc.main.id
#   cidr_block = "10.0.1.0/24"
#   tags = [{
#     key   = "Name"
#     value = "main"
#   }]
# }


# resource "awscc_security_group" "sg_for_lambda" {
#   name        = "allow_tls"
#   description = "Allow TLS inbound traffic"
#   vpc_id      = aws_vpc.main.id

#   ingress {
#     description      = "TLS from VPC"
#     from_port        = 443
#     to_port          = 443
#     protocol         = "tcp"
#     cidr_blocks      = [aws_vpc.main.cidr_block]
#     ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
#   }

#   egress {
#     from_port        = 0
#     to_port          = 0
#     protocol         = "-1"
#     cidr_blocks      = ["0.0.0.0/0"]
#     ipv6_cidr_blocks = ["::/0"]
#   }

#   tags = {
#     Name = "allow_tls"
#   }
# }

# # Mount target connects the file system to the subnet
# resource "awscc_efs_mount_target" "alpha" {
#   file_system_id  = awscc_efs_file_system.efs_for_lambda.id
#   subnet_id       = awscc_ec2_subnet.main.id
#   security_groups = [awscc_security_group.sg_for_lambda.id]
# }

# # EFS access point used by lambda file system
# resource "awscc_efs_access_point" "access_point_for_lambda" {
#   file_system_id = awscc_efs_file_system.efs_for_lambda.id

#   root_directory {
#     path = "/lambda"
#     creation_info {
#       owner_gid   = 1000
#       owner_uid   = 1000
#       permissions = "777"
#     }
#   }

#   posix_user {
#     gid = 1000
#     uid = 1000
#   }
# }


# resource "awscc_lambda_function" "main" {
#   function_name = "lambda_function_name"
#   description   = "AWS Lambda function"
#   code = {
#     zip_file = data.archive_file.main.output_path
#   }
#   file_system_config = {
#     # EFS file system access point ARN
#     arn = awscc_efs_access_point.access_point_for_lambda.arn

#     # Local mount path inside the lambda function. Must start with '/mnt/'.
#     local_mount_path = "/mnt/efs"
#   }

#   vpc_config = {
#     # Every subnet should be able to reach an EFS mount target in the same Availability Zone. Cross-AZ mounts are not permitted.
#     subnet_ids         = [awscc_ec2_subnet.subnet_for_lambda.id]
#     security_group_ids = [awscc_security_group.sg_for_lambda.id]
#   }

#   handler       = "main.lambda_handler"
#   runtime       = "python3.10"
#   timeout       = "300"
#   memory_size   = "128"
#   role          = awscc_iam_role.main.arn
#   architectures = ["arm64"]
#   environment = {
#     variables = {
#       MY_KEY_1 = "MY_VALUE_1"
#       MY_KEY_2 = "MY_VALUE_2"
#     }
#   }
# }


