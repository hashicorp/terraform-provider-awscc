---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_config_organization_conformance_pack Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Config::OrganizationConformancePack.
---

# awscc_config_organization_conformance_pack (Resource)

Resource schema for AWS::Config::OrganizationConformancePack.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_conformance_pack_name` (String) The name of the organization conformance pack.

### Optional

- `conformance_pack_input_parameters` (Attributes List) A list of ConformancePackInputParameter objects. (see [below for nested schema](#nestedatt--conformance_pack_input_parameters))
- `delivery_s3_bucket` (String) AWS Config stores intermediate files while processing conformance pack template.
- `delivery_s3_key_prefix` (String) The prefix for the delivery S3 bucket.
- `excluded_accounts` (List of String) A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.
- `template_body` (String) A string containing full conformance pack template body.
- `template_s3_uri` (String) Location of file containing the template body.

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--conformance_pack_input_parameters"></a>
### Nested Schema for `conformance_pack_input_parameters`

Optional:

- `parameter_name` (String)
- `parameter_value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_config_organization_conformance_pack.example
  id = "organization_conformance_pack_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_config_organization_conformance_pack.example "organization_conformance_pack_name"
```
