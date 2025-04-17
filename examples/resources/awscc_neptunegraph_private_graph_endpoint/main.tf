data "aws_region" "current" {}

# Example VPC for Neptune Graph
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "neptune-graph-vpc"
  }]
}

# Example Subnet 1
resource "awscc_ec2_subnet" "example1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "neptune-graph-subnet-1"
  }]
}

# Example Subnet 2
resource "awscc_ec2_subnet" "example2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "neptune-graph-subnet-2"
  }]
}

# Security Group for Neptune Graph
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for Neptune Graph"
  vpc_id            = awscc_ec2_vpc.example.id
  security_group_ingress = [{
    ip_protocol = "tcp"
    from_port   = 8182
    to_port     = 8182
    cidr_ip     = "10.0.0.0/16"
  }]
  tags = [{
    key   = "Name"
    value = "neptune-graph-sg"
  }]
}

# Example Graph resource
resource "awscc_neptunegraph_graph" "example" {
  vector_search_configuration = {
    vector_search_dimension = 1536
  }
  provisioned_memory = 2
}

# Private Graph Endpoint
resource "awscc_neptunegraph_private_graph_endpoint" "example" {
  graph_identifier   = awscc_neptunegraph_graph.example.id
  vpc_id             = awscc_ec2_vpc.example.id
  subnet_ids         = [awscc_ec2_subnet.example1.id, awscc_ec2_subnet.example2.id]
  security_group_ids = [awscc_ec2_security_group.example.id]
}