---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pcaconnectorad_template_group_access_control_entry Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::PCAConnectorAD::TemplateGroupAccessControlEntry
---

# awscc_pcaconnectorad_template_group_access_control_entry (Data Source)

Data Source schema for AWS::PCAConnectorAD::TemplateGroupAccessControlEntry



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `access_rights` (Attributes) (see [below for nested schema](#nestedatt--access_rights))
- `group_display_name` (String)
- `group_security_identifier` (String)
- `template_arn` (String)

<a id="nestedatt--access_rights"></a>
### Nested Schema for `access_rights`

Read-Only:

- `auto_enroll` (String)
- `enroll` (String)
