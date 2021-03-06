---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_vpc_endpoint Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::VPCEndpoint
---

# awscc_ec2_vpc_endpoint (Resource)

Resource Type definition for AWS::EC2::VPCEndpoint



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `service_name` (String)
- `vpc_id` (String)

### Optional

- `policy_document` (Map of String)
- `private_dns_enabled` (Boolean)
- `route_table_ids` (List of String)
- `security_group_ids` (List of String)
- `subnet_ids` (List of String)
- `vpc_endpoint_type` (String)

### Read-Only

- `creation_timestamp` (String)
- `dns_entries` (List of String)
- `id` (String) The ID of this resource.
- `network_interface_ids` (List of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_vpc_endpoint.example <resource ID>
```
