terraform {
  required_providers {
    awscc = {
      source = "hashicorp/awscc"
      version = "0.61.0"
    }
  }
}

provider "awscc" {
  region  = "ap-southeast-2"
  profile = "default"
}