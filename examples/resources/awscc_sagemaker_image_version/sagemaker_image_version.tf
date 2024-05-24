resource "awscc_sagemaker_image_version" "example" {
  image_name = "example"
  base_image = "012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest"
}

resource "awscc_sagemaker_image" "example" {
  image_name     = awscc_sagemaker_image.example.image_name
  image_role_arn = awscc_iam_role.example.arn
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}