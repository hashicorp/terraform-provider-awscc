resource "awscc_sagemaker_app_image_config" "example" {
  app_image_config_name = "example"

  kernel_gateway_image_config = {
    kernel_specs = [{
      name = "example"
    }]
    file_system_config = {
      default_gid = 100
      default_uid = 1000
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
