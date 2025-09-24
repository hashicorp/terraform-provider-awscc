resource "awscc_elasticloadbalancingv2_target_group" "instance-example" {
  name     = "instance-example"
  port     = 80
  protocol = "HTTP"
  vpc_id   = awscc_ec2_vpc.main.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = [
    {
      key   = "Name"
      value = "main"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}