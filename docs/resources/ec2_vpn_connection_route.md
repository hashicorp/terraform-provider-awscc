---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_vpn_connection_route Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::VPNConnectionRoute
---

# awscc_ec2_vpn_connection_route (Resource)

Resource Type definition for AWS::EC2::VPNConnectionRoute



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_cidr_block` (String) The CIDR block associated with the local subnet of the customer network.
- `vpn_connection_id` (String) The ID of the VPN connection.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_vpn_connection_route.example <resource ID>
```
