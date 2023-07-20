resource "awscc_iam_service_linked_role" "autoscaling" {
  aws_service_name = "autoscaling.amazonaws.com"
  description      = "service linked role for AWS Auto Scaling"
  custom_suffix    = "TestSuffix"
}