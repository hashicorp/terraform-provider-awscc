terraform {
  required_providers {
    awscc = {
      source  = "hashicorp/awscc"
      version = "~> 1.76.0"
    }
  }
}

resource "awscc_route53globalresolver_global_resolver" "example" {
  name = "example-global-resolver"
}
