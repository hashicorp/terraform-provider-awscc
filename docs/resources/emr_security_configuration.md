---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_emr_security_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Use a SecurityConfiguration resource to configure data encryption, Kerberos authentication, and Amazon S3 authorization for EMRFS.
---

# awscc_emr_security_configuration (Resource)

Use a SecurityConfiguration resource to configure data encryption, Kerberos authentication, and Amazon S3 authorization for EMRFS.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `security_configuration` (String) The security configuration details in JSON format.

### Optional

- `name` (String) The name of the security configuration.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_emr_security_configuration.example <resource ID>
```
