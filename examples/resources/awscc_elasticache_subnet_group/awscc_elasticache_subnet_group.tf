resource "awscc_ec2_vpc" "this" {
  cidr_block = "10.0.0.0/16"
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_ec2_subnet" "this" {
  vpc_id            = resource.awscc_ec2_vpc.this.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-1a"

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_elasticache_subnet_group" "this" {
  cache_subnet_group_name = "example-cache-group"
  description             = "example awscc cache subnet"
  subnet_ids              = [awscc_ec2_subnet.this.id]
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
