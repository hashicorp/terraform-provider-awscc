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
  vp_cs = [
    {
      vpc_id     = awscc_ec2_vpc.main.id
      vpc_region = "us-east-1"
    }
  ]
}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}
