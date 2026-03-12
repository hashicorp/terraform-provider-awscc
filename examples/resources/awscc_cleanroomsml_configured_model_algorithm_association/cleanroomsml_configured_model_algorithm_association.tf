
terraform {
  required_providers {
    awscc = {
      source  = "hashicorp/awscc"
      version = "~> 1.75.0"
    }
  }
}

provider "awscc" {
  region = "us-east-1"
}

resource "awscc_cleanroomsml_configured_model_algorithm_association" "example" {
  configured_model_algorithm_arn = "arn:aws:cleanroomsml:us-east-1:123456789012:configured-model-algorithm/example-algorithm"
  membership_identifier          = "example-membership-id"
  name                          = "example-configured-model-algorithm-association"
  
  tags = {
    Environment = "test"
    Project     = "example"
  }
}
