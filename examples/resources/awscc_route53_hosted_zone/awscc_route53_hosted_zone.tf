resource "awscc_route53_hosted_zone" "this" {
  name = "this.com"
  hosted_zone_tags = [
    {
      key   = "Name"
      value = "this"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
