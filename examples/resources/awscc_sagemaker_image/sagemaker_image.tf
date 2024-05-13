resource "awscc_sagemaker_image" "example" {
  image_name     = "example"
  image_role_arn = awscc_iam_role.main.arn
  tags = [
    {
      key   = "Name"
      value = "example"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
