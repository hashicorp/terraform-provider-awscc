# Get current AWS region
data "aws_region" "current" {}

# Create VPC Resources
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM roles and policies
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "nimble.amazonaws.com",
        "studio.amazonaws.com"
      ]
    }
  }
}

resource "awscc_iam_role" "studio" {
  assume_role_policy_document = jsonencode(data.aws_iam_policy_document.assume_role.json)
  role_name                   = "NimbleStudioStudioRole"
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AWSNimbleStudioStudioServicerole"]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Nimble Studio
resource "awscc_nimblestudio_studio" "example" {
  admin_role_arn = awscc_iam_role.studio.arn
  display_name   = "Example Studio"
  studio_name    = "example-studio"
  user_role_arn  = awscc_iam_role.studio.arn
  studio_encryption_configuration = {
    key_type = "AWS_OWNED_KEY"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Studio Component
resource "awscc_nimblestudio_studio_component" "example" {
  name        = "example-workstation"
  type        = "COMPUTE_FARM"
  studio_id   = awscc_nimblestudio_studio.example.studio_id
  description = "Example workstation component"
  configuration = {
    computeFarmConfiguration = {
      activeDirectoryConfiguration = {
        directoryId = "d-example"
      }
      endpoint = "example.nimblestudio.com"
    }
  }
  initialization_scripts = []
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Launch Profile
resource "awscc_nimblestudio_launch_profile" "example" {
  description                      = "Example Launch Profile"
  ec_2_subnet_ids                  = [awscc_ec2_subnet.example.id]
  launch_profile_protocol_versions = ["2021-03-31"]
  name                             = "example-launch-profile"
  stream_configuration = {
    clipboard_mode                = "ENABLED"
    ec_2_instance_types           = ["g4dn.xlarge"]
    streaming_image_ids           = ["ami-example"]
    max_session_length_in_minutes = 690
    session_storage = {
      mode = ["UPLOAD"]
      root = {
        linux   = "/home/user"
        windows = "C:\\Users\\user"
      }
    }
  }
  studio_component_ids = [awscc_nimblestudio_studio_component.example.studio_component_id]
  studio_id            = awscc_nimblestudio_studio.example.studio_id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}