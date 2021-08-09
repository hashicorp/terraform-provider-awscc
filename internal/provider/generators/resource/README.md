# Terraform Resource Schema Generator

Generate a Terraform resource schema from a CloudFormation resource schema.

This tool

* Parses a CloudFormation resource type schema
* Generates Go code for the schema targeting the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework)
