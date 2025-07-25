---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_shield_proactive_engagement Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
---

# awscc_shield_proactive_engagement (Resource)

Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `emergency_contact_list` (Attributes List) A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.
To enable proactive engagement, the contact list must include at least one phone number. (see [below for nested schema](#nestedatt--emergency_contact_list))
- `proactive_engagement_status` (String) If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
If `DISABLED`, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.

### Read-Only

- `account_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--emergency_contact_list"></a>
### Nested Schema for `emergency_contact_list`

Required:

- `email_address` (String) The email address for the contact.

Optional:

- `contact_notes` (String) Additional notes regarding the contact.
- `phone_number` (String) The phone number for the contact

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_shield_proactive_engagement.example
  id = "account_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_shield_proactive_engagement.example "account_id"
```
