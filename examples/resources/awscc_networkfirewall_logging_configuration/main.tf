# Get the current AWS region
data "aws_region" "current" {}

# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create subnets
resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# S3 bucket for logging
resource "awscc_s3_bucket" "logging" {
  bucket_name = "network-firewall-logs-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Bucket policy for Network Firewall logging
resource "aws_s3_bucket_policy" "logging" {
  bucket = awscc_s3_bucket.logging.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowNetworkFirewallLogging"
        Effect = "Allow"
        Principal = {
          Service = "network-firewall.amazonaws.com"
        }
        Action = [
          "s3:PutObject"
        ]
        Resource = [
          "${awscc_s3_bucket.logging.arn}/*"
        ]
      }
    ]
  })
}

# Create Network Firewall Policy
resource "awscc_networkfirewall_firewall_policy" "example" {
  firewall_policy_name = "example-policy"
  firewall_policy = {
    stateless_default_actions          = ["aws:forward_to_sfe"]
    stateless_fragment_default_actions = ["aws:forward_to_sfe"]
    stateful_rule_group_reference      = []
    stateless_rule_group_reference     = []
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Local values for ARNs
locals {
  policy_arn   = "arn:aws:network-firewall:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:firewall-policy/example-policy"
  firewall_arn = "arn:aws:network-firewall:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:firewall/example-firewall"
}

# Create Network Firewall
resource "awscc_networkfirewall_firewall" "example" {
  firewall_name       = "example-firewall"
  firewall_policy_arn = local.policy_arn
  vpc_id              = awscc_ec2_vpc.example.id
  subnet_mappings = [{
    subnet_id = awscc_ec2_subnet.example.id
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Network Firewall Logging Configuration
resource "awscc_networkfirewall_logging_configuration" "example" {
  firewall_arn  = local.firewall_arn
  firewall_name = "example-firewall"
  logging_configuration = {
    log_destination_configs = [
      {
        log_destination = {
          bucketName = awscc_s3_bucket.logging.id
        }
        log_destination_type = "S3"
        log_type             = "ALERT"
      }
    ]
  }
}