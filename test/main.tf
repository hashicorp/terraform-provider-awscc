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

# resource "awscc_ecs_task_definition" "test" {
#   container_definitions = [
#     {
#       command = ["sleep", "10"]
#       cpu     = 10
#       docker_labels = {
#         "env" = "testing"
#       }
#       entry_point = ["/"]
#       environment = [
#         {
#           name  = "VARNAME"
#           value = "VARVAL"
#         }
#       ]
#       essential = true
#       image     = "jenkins/jenkins:lts"
#       memory    = 128
#       name      = "jenkins"
#       port_mappings = [
#         {
#           container_port = 80
#           host_port      = 8080
#         }
#       ]
#     }
#   ]
# }
