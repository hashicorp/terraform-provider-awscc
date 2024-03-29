---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3_access_grant Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::S3::AccessGrant
---

# awscc_s3_access_grant (Data Source)

Data Source schema for AWS::S3::AccessGrant



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `access_grant_arn` (String) The Amazon Resource Name (ARN) of the specified access grant.
- `access_grant_id` (String) The ID assigned to this access grant.
- `access_grants_location_configuration` (Attributes) The configuration options of the grant location, which is the S3 path to the data to which you are granting access. (see [below for nested schema](#nestedatt--access_grants_location_configuration))
- `access_grants_location_id` (String) The custom S3 location to be accessed by the grantee
- `application_arn` (String) The ARN of the application grantees will use to access the location
- `grant_scope` (String) The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix.
- `grantee` (Attributes) The principal who will be granted permission to access S3. (see [below for nested schema](#nestedatt--grantee))
- `permission` (String) The level of access to be afforded to the grantee
- `s3_prefix_type` (String) The type of S3SubPrefix.
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--access_grants_location_configuration"></a>
### Nested Schema for `access_grants_location_configuration`

Read-Only:

- `s3_sub_prefix` (String) The S3 sub prefix of a registered location in your S3 Access Grants instance


<a id="nestedatt--grantee"></a>
### Nested Schema for `grantee`

Read-Only:

- `grantee_identifier` (String) The unique identifier of the Grantee
- `grantee_type` (String) Configures the transfer acceleration state for an Amazon S3 bucket.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
