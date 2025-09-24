resource "awscc_codebuild_fleet" "example" {
  name             = "example"
  base_capacity    = 1
  compute_type     = "BUILD_GENERAL1_SMALL"
  environment_type = "LINUX_CONTAINER"
}