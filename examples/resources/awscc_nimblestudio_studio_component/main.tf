# Studio role setup
data "aws_iam_policy_document" "studio_role" {
  statement {
    effect = "Allow"
    principals {
      type = "Service"
      identifiers = [
        "nimble.amazonaws.com",
        "studio.amazonaws.com"
      ]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "awscc_iam_role" "studio_role" {
  role_name                   = "NimbleStudioStudioRole"
  description                 = "IAM role for Nimble Studio"
  assume_role_policy_document = jsonencode(data.aws_iam_policy_document.studio_role.json)

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# The studio itself
resource "awscc_nimblestudio_studio" "example" {
  admin_role_arn = awscc_iam_role.studio_role.arn
  user_role_arn  = awscc_iam_role.studio_role.arn
  display_name   = "Example Studio"
  studio_name    = "ExampleStudio"
  studio_encryption_configuration = {
    key_type = "AWS_OWNED_KEY"
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Example security group for the studio component
resource "awscc_ec2_security_group" "example" {
  group_description = "Example security group for NimbleStudio Component"
  vpc_id            = "vpc-12345678" # Replace with your VPC ID

  group_name = "nimble-studio-component-sg"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# The studio component
resource "awscc_nimblestudio_studio_component" "example" {
  name      = "ExampleStudioComponent"
  type      = "COMPUTE_FARM"
  studio_id = awscc_nimblestudio_studio.example.id

  configuration = {
    compute_farm_configuration = {
      active_directory_user = "admin"
      endpoint              = "https://nimblestudio-example.com"
    }
  }

  description = "Example Studio Component for testing"

  ec_2_security_group_ids = [awscc_ec2_security_group.example.id]

  initialization_scripts = [{
    launch_profile_protocol_version = "2021-03-31"
    platform                        = "WINDOWS"
    run_context                     = "SYSTEM_INITIALIZATION"
    script                          = "echo 'Hello from initialization script'"
  }]

  script_parameters = [{
    key   = "example_key"
    value = "example_value"
  }]

  subtype = "NIMBLE_ON_DEMAND"

  tags = {
    "Modified By" = "AWSCC"
  }
}