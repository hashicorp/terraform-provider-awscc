resource "awscc_sagemaker_app_image_config" "example" {
  app_image_config_name = "example"

  kernel_gateway_image_config = {
    kernel_specs = [{
      name = "example"
    }]
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
