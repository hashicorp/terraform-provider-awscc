# Use existing VPC resources
data "aws_vpc" "example" {
  default = true
}

data "aws_subnets" "example" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.example.id]
  }
  filter {
    name   = "default-for-az"
    values = ["true"]
  }
}

# Create Redshift Cluster Subnet Group
resource "awscc_redshift_cluster_subnet_group" "example" {
  description = "Example Redshift cluster subnet group"
  subnet_ids  = data.aws_subnets.example.ids

  tags = [
    {
      key   = "Name"
      value = "example-subnet-group"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
