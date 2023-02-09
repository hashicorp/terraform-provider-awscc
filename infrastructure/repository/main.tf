# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  backend "remote" {
    organization = "hashicorp-v2"

    workspaces {
      name = "terraform-provider-awscc-repository"
    }
  }

  required_providers {
    github = {
      source  = "integrations/github"
      version = "4.14.0"
    }
  }

  required_version = ">= 1.0.0"
}

provider "github" {
  owner = "hashicorp"
}
