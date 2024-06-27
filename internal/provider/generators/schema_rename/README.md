# AWS CloudFormation Resource Schema Renamer

Find any reference to the **AWS CloudFormation** in the resource schema attributes / descriptions.

This tool

* Reads the schema files under `./internal/service/cloudformation/schemas/`
* Checks whether any reference to **AWS CloudFormation** exists in the schema attributes or description.
* Replace reference of **AWS CloudFormation** with **Terraform**

## Allowlist

This tool will skips any CloudFormation schema file with prefix `AWS_CloudFormation_` which denotes all AWS CloudFormation resources (stack, hooks, etc).
