# Copyright IBM Corp. 2021, 2026
# SPDX-License-Identifier: MPL-2.0

terraform {
  backend "remote" {
    organization = "hashicorp-terraform-aws-provider"

    workspaces {
      name = "terraform-provider-awscc-repository"
    }
  }

  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6"
    }
  }

  required_version = "~> 1.15"
}

provider "github" {
  owner = "hashicorp"
}
