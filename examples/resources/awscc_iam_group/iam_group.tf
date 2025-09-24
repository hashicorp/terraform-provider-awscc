resource "awscc_iam_group" "example" {
  group_name          = "example"
  managed_policy_arns = ["arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"]
  path                = "/"
}