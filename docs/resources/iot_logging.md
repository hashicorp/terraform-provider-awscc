---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_logging Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.
---

# awscc_iot_logging (Resource)

Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
- `default_log_level` (String) The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
- `role_arn` (String) The ARN of the role that allows IoT to write to Cloudwatch logs.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iot_logging.example <resource ID>
```
