resource "awscc_sagemaker_domain" "example" {
  domain_name = "example"
  auth_mode   = "IAM"
  vpc_id      = awscc_ec2_vpc.main.id
  subnet_ids  = [awscc_ec2_subnet.this.id]

  default_user_settings = {
    execution_role = awscc_iam_role.example.arn
    kernel_gateway_app_settings = {
      custom_image = {
        app_image_config_name = awscc_sagemaker_app_image_config.example.app_image_config_name
        image_name            = awscc_sagemaker_image_version.example.image_name
      }
    }
  }
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}


variable "image_name" {
  type = string
}

resource "awscc_sagemaker_image" "example" {
  image_name     = var.image_name
  image_role_arn = awscc_iam_role.example.arn
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_sagemaker_image_version" "example" {
  image_name = var.image_name
  base_image = "012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest"
  depends_on = [
    awscc_sagemaker_image.example
  ]
}

resource "awscc_sagemaker_app_image_config" "example" {
  app_image_config_name = "example"

  kernel_gateway_image_config = {
    kernel_specs = [{
      name = "example"
    }]
  }
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}