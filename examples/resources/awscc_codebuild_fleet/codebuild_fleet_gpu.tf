resource "awscc_codebuild_fleet" "example" {
  name             = "example"
  base_capacity    = 1
  compute_type     = "BUILD_GENERAL1_LARGE"
  environment_type = "LINUX_GPU_CONTAINER"
}