resource "awscc_ec2_eip" "main" {
  domain      = "vpc"
  instance_id = var.instance_id
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}