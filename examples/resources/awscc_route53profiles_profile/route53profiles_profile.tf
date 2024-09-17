resource "awscc_route53profiles_profile" "example" {
  name = "example"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}