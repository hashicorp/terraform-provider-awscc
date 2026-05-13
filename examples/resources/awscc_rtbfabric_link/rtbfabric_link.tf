terraform {
  required_version = ">= 1.0"
  required_providers {
    awscc = {
      source  = "hashicorp/awscc"
      version = ">= 1.76.0"
    }
  }
}

resource "awscc_rtbfabric_link" "example" {
  link_name = "example-rtb-fabric-link"
  
  tags = {
    Name        = "example-rtb-fabric-link"
    Environment = "test"
  }
}

output "rtb_fabric_link_arn" {
  description = "The ARN of the RTB Fabric Link"
  value       = awscc_rtbfabric_link.example.arn
}

output "rtb_fabric_link_id" {
  description = "The ID of the RTB Fabric Link"
  value       = awscc_rtbfabric_link.example.id
}
