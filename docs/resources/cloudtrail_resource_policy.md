---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudtrail_resource_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::CloudTrail::ResourcePolicy
---

# awscc_cloudtrail_resource_policy (Resource)

Resource Type definition for AWS::CloudTrail::ResourcePolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_arn` (String) The ARN of the AWS CloudTrail resource to which the policy applies.
- `resource_policy` (String) A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cloudtrail_resource_policy.example <resource ID>
```
