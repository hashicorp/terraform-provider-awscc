resource "awscc_sagemaker_image" "example" {
  image_name     = "example"
  image_role_arn = aws_iam_role.test.arn
}