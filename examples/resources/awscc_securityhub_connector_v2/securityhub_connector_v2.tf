terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
    awscc = {
      source  = "hashicorp/awscc"
      version = "~> 1.68.0"
    }
  }
}

provider "aws" {
  region = "us-west-2"
}

provider "awscc" {
  region = "us-west-2"
}

# Security Hub Connector V2 resource with Jira Cloud integration
resource "awscc_securityhub_connector_v2" "example" {
  name        = "example-jira-connector"
  description = "Example Security Hub Connector for Jira Cloud integration"
  provider_type = "JiraCloud"
  
  provider_configuration = {
    jira_cloud = {
      hostname      = "example.atlassian.net"
      project_key   = "SEC"
      secret_arn    = "arn:aws:secretsmanager:us-west-2:123456789012:secret:jira-credentials-abcdef"
    }
  }
}

# Output the connector ID
output "connector_id" {
  value = awscc_securityhub_connector_v2.example.id
}

# Output the connector name
output "connector_name" {
  value = awscc_securityhub_connector_v2.example.name
}
