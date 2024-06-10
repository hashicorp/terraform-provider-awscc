resource "awscc_sagemaker_user_profile" "example" {
  domain_id         = awscc_sagemaker_domain.example.id
  user_profile_name = "example"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}