---
page_title: "awscc_s3_bucket Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::S3::Bucket resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
  To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to retain the bucket or to delete the bucket. For more information, see DeletionPolicy Attribute https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html.
  You can only delete empty buckets. Deletion fails for buckets that have contents.
---

# awscc_s3_bucket (Resource)

The ``AWS::S3::Bucket`` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
 To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).
  You can only delete empty buckets. Deletion fails for buckets that have contents.
## Example Usage

### Create an S3 bucket 

To create an S3 bucket

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

}
```

### Create an S3 bucket with public access restricted 

To create an S3 bucket with public access restricted

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

}
```

### S3 bucket with default encryption AES256

To create an S3 bucket with server side default encryption AES256

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm = "AES256"
      }
    }]
  }
}
```

### S3 bucket with default encryption KMS

To create an S3 bucket with server side encryption using KMS

```terraform
resource "awscc_kms_key" "example" {
  description         = "S3 KMS key"
  enable_key_rotation = true
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-kms"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = awscc_kms_key.example.arn
      }
    }]
  }
}
```

### S3 bucket with versioning enabled

Creates an S3 bucket with versioning enabled.

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-versioned"
  versioning_configuration = {
    status = "Enabled"
  }

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

}
```

### S3 bucket with ownership controls specified

Creates an S3 bucket with BucketOwnerPreferred specified as the object owner.

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"
  ownership_controls = {
    rules = [{
      object_ownership = "BucketOwnerPreferred"
    }]
  }

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

}
```

### S3 bucket with object expiration lifecycle rules

Creates an S3 bucket with a lifecycle rule to expire non_current version of objects with inputs to classify the current/non-current versions.

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "expire_non_current_version"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        status = "Enabled"
      }
    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}
```

### S3 bucket with object expiration lifecycle rules with a filter based on both prefix and one or more tags

The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with the specified prefix and two tags (with the specific tag keys and values)

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "expire_non_current_version_filtered_by_tags"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        prefix = "logs/"
        tag_filters = [{
          key   = "key1"
          value = "value1"
          },
          {
            key   = "key2"
            value = "value2"
          }
        ]
        status = "Enabled"
      }
    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}
```

### S3 bucket with abort multipart upload lifecycle rule

Creates an S3 bucket with a lifecycle rule to configure the days after which Amazon S3 aborts and incomplete multipart upload.

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "abort_incomplete_upload"
        abort_incomplete_multipart_upload = {
          days_after_initiation = 1
        }
        status = "Enabled"
      }

    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}
```

### Specifying a filter based on object size

Creates an S3 bucket with a lifecycle rule filtered on object size greater than a specified value. Object size values are in bytes. Maximum filter size is 5TB. Some storage classes have minimum object size limitations, for more information, see [Comparing the Amazon S3 storage classes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html#sc-compare).

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {

        id = "expire_non_current_version"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        object_size_greater_than = 500
        status                   = "Enabled"
      }
    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}
```

### Specifying a filter based on object size range and prefix

Creates an S3 bucket with a lifecycle rule based on object size range and a prefix. The `object_size_greater_than` must be less than the `object_size_less_than`.

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {

        id = "expire_non_current_version"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        object_size_greater_than = 500
        status                   = "Enabled"
      }
    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}
```

### Specifying a lifecycle rule to transition objects between storage classes

Creates an S3 bucket with a lifecycle rule which moves non current versions of objects to different storage classes based on predefined days.

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "non_current_version_transitions"

        noncurrent_version_expiration_in_days = 90
        noncurrent_version_transitions = [
          {
            transition_in_days = 30
            storage_class      = "STANDARD_IA"
          },
          {
            transition_in_days = 60
            storage_class      = "INTELLIGENT_TIERING"
          }
        ]
        status = "Enabled"
      }
    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `accelerate_configuration` (Attributes) Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*. (see [below for nested schema](#nestedatt--accelerate_configuration))
- `access_control` (String) This is a legacy property, and it is not recommended for most use cases. A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.
  A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide*.
  S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy with any value other than Private. Use cases requiring ACLs are uncommon.
  The majority of access control configurations can be successfully and more easily achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3 Server Access Logs buckets and more, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) in the *Amazon S3 User Guide*.
- `analytics_configurations` (Attributes List) Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket. (see [below for nested schema](#nestedatt--analytics_configurations))
- `bucket_encryption` (Attributes) Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*. (see [below for nested schema](#nestedatt--bucket_encryption))
- `bucket_name` (String) A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) in the *Amazon S3 User Guide*. 
  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
- `cors_configuration` (Attributes) Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*. (see [below for nested schema](#nestedatt--cors_configuration))
- `intelligent_tiering_configurations` (Attributes List) Defines how Amazon S3 handles Intelligent-Tiering storage. (see [below for nested schema](#nestedatt--intelligent_tiering_configurations))
- `inventory_configurations` (Attributes List) Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*. (see [below for nested schema](#nestedatt--inventory_configurations))
- `lifecycle_configuration` (Attributes) Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*. (see [below for nested schema](#nestedatt--lifecycle_configuration))
- `logging_configuration` (Attributes) Settings that define where logs are stored. (see [below for nested schema](#nestedatt--logging_configuration))
- `metadata_table_configuration` (Attributes) The metadata table configuration of an S3 general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) and [Setting up permissions for configuring metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html). (see [below for nested schema](#nestedatt--metadata_table_configuration))
- `metrics_configurations` (Attributes List) Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html). (see [below for nested schema](#nestedatt--metrics_configurations))
- `notification_configuration` (Attributes) Configuration that defines how Amazon S3 handles bucket notifications. (see [below for nested schema](#nestedatt--notification_configuration))
- `object_lock_configuration` (Attributes) This operation is not supported for directory buckets.
  Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). 
   +  The ``DefaultRetention`` settings require both a mode and a period.
  +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
  +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html). (see [below for nested schema](#nestedatt--object_lock_configuration))
- `object_lock_enabled` (Boolean) Indicates whether this bucket has an Object Lock configuration enabled. Enable ``ObjectLockEnabled`` when you apply ``ObjectLockConfiguration`` to a bucket.
- `ownership_controls` (Attributes) Configuration that defines how Amazon S3 handles Object Ownership rules. (see [below for nested schema](#nestedatt--ownership_controls))
- `public_access_block_configuration` (Attributes) Configuration that defines how Amazon S3 handles public access. (see [below for nested schema](#nestedatt--public_access_block_configuration))
- `replication_configuration` (Attributes) Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the ``VersioningConfiguration`` property.
 Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist. (see [below for nested schema](#nestedatt--replication_configuration))
- `tags` (Attributes List) An arbitrary set of tags (key-value pairs) for this S3 bucket. (see [below for nested schema](#nestedatt--tags))
- `versioning_configuration` (Attributes) Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
  When you enable versioning on a bucket for the first time, it might take a short amount of time for the change to be fully propagated. We recommend that you wait for 15 minutes after enabling versioning before issuing write operations (``PUT`` or ``DELETE``) on objects in the bucket. (see [below for nested schema](#nestedatt--versioning_configuration))
- `website_configuration` (Attributes) Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html). (see [below for nested schema](#nestedatt--website_configuration))

### Read-Only

- `arn` (String) the Amazon Resource Name (ARN) of the specified bucket.
- `domain_name` (String)
- `dual_stack_domain_name` (String)
- `id` (String) Uniquely identifies the resource.
- `regional_domain_name` (String)
- `website_url` (String)

<a id="nestedatt--accelerate_configuration"></a>
### Nested Schema for `accelerate_configuration`

Optional:

- `acceleration_status` (String) Specifies the transfer acceleration status of the bucket.


<a id="nestedatt--analytics_configurations"></a>
### Nested Schema for `analytics_configurations`

Optional:

- `id` (String) The ID that identifies the analytics configuration.
- `prefix` (String) The prefix that an object must have to be included in the analytics results.
- `storage_class_analysis` (Attributes) Contains data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes. (see [below for nested schema](#nestedatt--analytics_configurations--storage_class_analysis))
- `tag_filters` (Attributes List) The tags to use when evaluating an analytics filter.
 The analytics only includes objects that meet the filter's criteria. If no filter is specified, all of the contents of the bucket are included in the analysis. (see [below for nested schema](#nestedatt--analytics_configurations--tag_filters))

<a id="nestedatt--analytics_configurations--storage_class_analysis"></a>
### Nested Schema for `analytics_configurations.storage_class_analysis`

Optional:

- `data_export` (Attributes) Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported. (see [below for nested schema](#nestedatt--analytics_configurations--storage_class_analysis--data_export))

<a id="nestedatt--analytics_configurations--storage_class_analysis--data_export"></a>
### Nested Schema for `analytics_configurations.storage_class_analysis.data_export`

Optional:

- `destination` (Attributes) The place to store the data for an analysis. (see [below for nested schema](#nestedatt--analytics_configurations--storage_class_analysis--data_export--destination))
- `output_schema_version` (String) The version of the output schema to use when exporting data. Must be ``V_1``.

<a id="nestedatt--analytics_configurations--storage_class_analysis--data_export--destination"></a>
### Nested Schema for `analytics_configurations.storage_class_analysis.data_export.destination`

Optional:

- `bucket_account_id` (String) The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not validated before exporting data.
   Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.
- `bucket_arn` (String) The Amazon Resource Name (ARN) of the bucket to which data is exported.
- `format` (String) Specifies the file format used when exporting data to Amazon S3.
  *Allowed values*: ``CSV`` | ``ORC`` | ``Parquet``
- `prefix` (String) The prefix to use when exporting data. The prefix is prepended to all results.




<a id="nestedatt--analytics_configurations--tag_filters"></a>
### Nested Schema for `analytics_configurations.tag_filters`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.



<a id="nestedatt--bucket_encryption"></a>
### Nested Schema for `bucket_encryption`

Optional:

- `server_side_encryption_configuration` (Attributes List) Specifies the default server-side-encryption configuration. (see [below for nested schema](#nestedatt--bucket_encryption--server_side_encryption_configuration))

<a id="nestedatt--bucket_encryption--server_side_encryption_configuration"></a>
### Nested Schema for `bucket_encryption.server_side_encryption_configuration`

Optional:

- `bucket_key_enabled` (Boolean) Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Setting the ``BucketKeyEnabled`` element to ``true`` causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.
 For more information, see [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) in the *Amazon S3 User Guide*.
- `server_side_encryption_by_default` (Attributes) Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. (see [below for nested schema](#nestedatt--bucket_encryption--server_side_encryption_configuration--server_side_encryption_by_default))

<a id="nestedatt--bucket_encryption--server_side_encryption_configuration--server_side_encryption_by_default"></a>
### Nested Schema for `bucket_encryption.server_side_encryption_configuration.server_side_encryption_by_default`

Optional:

- `kms_master_key_id` (String) AWS Key Management Service (KMS) customer managed key ID to use for the default encryption. 
   +   *General purpose buckets* - This parameter is allowed if and only if ``SSEAlgorithm`` is set to ``aws:kms`` or ``aws:kms:dsse``.
  +   *Directory buckets* - This parameter is allowed if and only if ``SSEAlgorithm`` is set to ``aws:kms``.
  
  You can specify the key ID, key alias, or the Amazon Resource Name (ARN) of the KMS key.
  +  Key ID: ``1234abcd-12ab-34cd-56ef-1234567890ab`` 
  +  Key ARN: ``arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`` 
  +  Key Alias: ``alias/alias-name`` 
  
 If you are using encryption with cross-account or AWS service operations, you must use a fully qualified KMS key ARN. For more information, see [Using encryption for cross-account operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html#bucket-encryption-update-bucket-policy).
   +   *General purpose buckets* - If you're specifying a customer managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key alias instead, then KMS resolves the key within the requester?s account. This behavior can result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket owner. Also, if you use a key ID, you can run into a LogDestination undeliverable error when creating a VPC flow log. 
  +   *Directory buckets* - When you specify an [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk) for encryption in your directory bucket, only use the key ID or key ARN. The key alias format of the KMS key isn't supported.
  
   Amazon S3 only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *Key Management Service Developer Guide*.
- `sse_algorithm` (String) Server-side encryption algorithm to use for the default encryption.
  For directory buckets, there are only two supported values for server-side encryption: ``AES256`` and ``aws:kms``.




<a id="nestedatt--cors_configuration"></a>
### Nested Schema for `cors_configuration`

Optional:

- `cors_rules` (Attributes List) A set of origins and methods (cross-origin access that you want to allow). You can add up to 100 rules to the configuration. (see [below for nested schema](#nestedatt--cors_configuration--cors_rules))

<a id="nestedatt--cors_configuration--cors_rules"></a>
### Nested Schema for `cors_configuration.cors_rules`

Optional:

- `allowed_headers` (List of String) Headers that are specified in the ``Access-Control-Request-Headers`` header. These headers are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns any requested headers that are allowed.
- `allowed_methods` (List of String) An HTTP method that you allow the origin to run.
  *Allowed values*: ``GET`` | ``PUT`` | ``HEAD`` | ``POST`` | ``DELETE``
- `allowed_origins` (List of String) One or more origins you want customers to be able to access the bucket from.
- `exposed_headers` (List of String) One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript ``XMLHttpRequest`` object).
- `id` (String) A unique identifier for this rule. The value must be no more than 255 characters.
- `max_age` (Number) The time in seconds that your browser is to cache the preflight response for the specified resource.



<a id="nestedatt--intelligent_tiering_configurations"></a>
### Nested Schema for `intelligent_tiering_configurations`

Optional:

- `id` (String) The ID used to identify the S3 Intelligent-Tiering configuration.
- `prefix` (String) An object key name prefix that identifies the subset of objects to which the rule applies.
- `status` (String) Specifies the status of the configuration.
- `tag_filters` (Attributes List) A container for a key-value pair. (see [below for nested schema](#nestedatt--intelligent_tiering_configurations--tag_filters))
- `tierings` (Attributes List) Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration. At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: ``ARCHIVE_ACCESS`` and ``DEEP_ARCHIVE_ACCESS``.
  You only need Intelligent Tiering Configuration enabled on a bucket if you want to automatically move objects stored in the Intelligent-Tiering storage class to Archive Access or Deep Archive Access tiers. (see [below for nested schema](#nestedatt--intelligent_tiering_configurations--tierings))

<a id="nestedatt--intelligent_tiering_configurations--tag_filters"></a>
### Nested Schema for `intelligent_tiering_configurations.tag_filters`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.


<a id="nestedatt--intelligent_tiering_configurations--tierings"></a>
### Nested Schema for `intelligent_tiering_configurations.tierings`

Optional:

- `access_tier` (String) S3 Intelligent-Tiering access tier. See [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) for a list of access tiers in the S3 Intelligent-Tiering storage class.
- `days` (Number) The number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier. The minimum number of days specified for Archive Access tier must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to 2 years (730 days).



<a id="nestedatt--inventory_configurations"></a>
### Nested Schema for `inventory_configurations`

Optional:

- `destination` (Attributes) Contains information about where to publish the inventory results. (see [below for nested schema](#nestedatt--inventory_configurations--destination))
- `enabled` (Boolean) Specifies whether the inventory is enabled or disabled. If set to ``True``, an inventory list is generated. If set to ``False``, no inventory list is generated.
- `id` (String) The ID used to identify the inventory configuration.
- `included_object_versions` (String) Object versions to include in the inventory list. If set to ``All``, the list includes all the object versions, which adds the version-related fields ``VersionId``, ``IsLatest``, and ``DeleteMarker`` to the list. If set to ``Current``, the list does not contain these version-related fields.
- `optional_fields` (List of String) Contains the optional fields that are included in the inventory results.
- `prefix` (String) Specifies the inventory filter prefix.
- `schedule_frequency` (String) Specifies the schedule for generating inventory results.

<a id="nestedatt--inventory_configurations--destination"></a>
### Nested Schema for `inventory_configurations.destination`

Optional:

- `bucket_account_id` (String) The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not validated before exporting data.
   Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.
- `bucket_arn` (String) The Amazon Resource Name (ARN) of the bucket to which data is exported.
- `format` (String) Specifies the file format used when exporting data to Amazon S3.
  *Allowed values*: ``CSV`` | ``ORC`` | ``Parquet``
- `prefix` (String) The prefix to use when exporting data. The prefix is prepended to all results.



<a id="nestedatt--lifecycle_configuration"></a>
### Nested Schema for `lifecycle_configuration`

Optional:

- `rules` (Attributes List) A lifecycle rule for individual objects in an Amazon S3 bucket. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules))
- `transition_default_minimum_object_size` (String) Indicates which default minimum object size behavior is applied to the lifecycle configuration.
  This parameter applies to general purpose buckets only. It isn't supported for directory bucket lifecycle configurations.
   +   ``all_storage_classes_128K`` - Objects smaller than 128 KB will not transition to any storage class by default.
  +   ``varies_by_storage_class`` - Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes will prevent transitions smaller than 128 KB. 
  
 To customize the minimum object size for any transition you can add a filter that specifies a custom ``ObjectSizeGreaterThan`` or ``ObjectSizeLessThan`` in the body of your transition rule. Custom filters always take precedence over the default transition behavior.

<a id="nestedatt--lifecycle_configuration--rules"></a>
### Nested Schema for `lifecycle_configuration.rules`

Optional:

- `abort_incomplete_multipart_upload` (Attributes) Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3 bucket. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--abort_incomplete_multipart_upload))
- `expiration_date` (String) Indicates when objects are deleted from Amazon S3 and Amazon S3 Glacier. The date value must be in ISO 8601 format. The time is always midnight UTC. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.
- `expiration_in_days` (Number) Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon S3 Glacier. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.
- `expired_object_delete_marker` (Boolean) Indicates whether Amazon S3 will remove a delete marker without any noncurrent versions. If set to true, the delete marker will be removed if there are no noncurrent versions. This cannot be specified with ``ExpirationInDays``, ``ExpirationDate``, or ``TagFilters``.
- `id` (String) Unique identifier for the rule. The value can't be longer than 255 characters.
- `noncurrent_version_expiration` (Attributes) Specifies when noncurrent object versions expire. Upon expiration, S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that S3 delete noncurrent object versions at a specific period in the object's lifetime. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--noncurrent_version_expiration))
- `noncurrent_version_expiration_in_days` (Number) (Deprecated.) For buckets with versioning enabled (or suspended), specifies the time, in days, between when a new version of the object is uploaded to the bucket and when old versions of the object expire. When object versions expire, Amazon S3 permanently deletes them. If you specify a transition and expiration time, the expiration time must be later than the transition time.
- `noncurrent_version_transition` (Attributes) (Deprecated.) For buckets with versioning enabled (or suspended), specifies when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the ``NoncurrentVersionTransitions`` property. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--noncurrent_version_transition))
- `noncurrent_version_transitions` (Attributes List) For buckets with versioning enabled (or suspended), one or more transition rules that specify when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the ``NoncurrentVersionTransition`` property. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--noncurrent_version_transitions))
- `object_size_greater_than` (String) Specifies the minimum object size in bytes for this rule to apply to. Objects must be larger than this value in bytes. For more information about size based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide*.
- `object_size_less_than` (String) Specifies the maximum object size in bytes for this rule to apply to. Objects must be smaller than this value in bytes. For more information about sized based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide*.
- `prefix` (String) Object key prefix that identifies one or more objects to which this rule applies.
  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
- `status` (String) If ``Enabled``, the rule is currently being applied. If ``Disabled``, the rule is not currently being applied.
- `tag_filters` (Attributes List) Tags to use to identify a subset of objects to which the lifecycle rule applies. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--tag_filters))
- `transition` (Attributes) (Deprecated.) Specifies when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the ``Transitions`` property. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--transition))
- `transitions` (Attributes List) One or more transition rules that specify when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the ``Transition`` property. (see [below for nested schema](#nestedatt--lifecycle_configuration--rules--transitions))

<a id="nestedatt--lifecycle_configuration--rules--abort_incomplete_multipart_upload"></a>
### Nested Schema for `lifecycle_configuration.rules.abort_incomplete_multipart_upload`

Optional:

- `days_after_initiation` (Number) Specifies the number of days after which Amazon S3 stops an incomplete multipart upload.


<a id="nestedatt--lifecycle_configuration--rules--noncurrent_version_expiration"></a>
### Nested Schema for `lifecycle_configuration.rules.noncurrent_version_expiration`

Optional:

- `newer_noncurrent_versions` (Number) Specifies how many noncurrent versions S3 will retain. If there are this many more recent noncurrent versions, S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide*.
- `noncurrent_days` (Number) Specifies the number of days an object is noncurrent before S3 can perform the associated action. For information about the noncurrent days calculations, see [How Amazon S3 Calculates When an Object Became Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide*.


<a id="nestedatt--lifecycle_configuration--rules--noncurrent_version_transition"></a>
### Nested Schema for `lifecycle_configuration.rules.noncurrent_version_transition`

Optional:

- `newer_noncurrent_versions` (Number) Specifies how many noncurrent versions S3 will retain. If there are this many more recent noncurrent versions, S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide*.
- `storage_class` (String) The class of storage used to store the object.
- `transition_in_days` (Number) Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action. For information about the noncurrent days calculations, see [How Amazon S3 Calculates How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide*.


<a id="nestedatt--lifecycle_configuration--rules--noncurrent_version_transitions"></a>
### Nested Schema for `lifecycle_configuration.rules.noncurrent_version_transitions`

Optional:

- `newer_noncurrent_versions` (Number) Specifies how many noncurrent versions S3 will retain. If there are this many more recent noncurrent versions, S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide*.
- `storage_class` (String) The class of storage used to store the object.
- `transition_in_days` (Number) Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action. For information about the noncurrent days calculations, see [How Amazon S3 Calculates How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide*.


<a id="nestedatt--lifecycle_configuration--rules--tag_filters"></a>
### Nested Schema for `lifecycle_configuration.rules.tag_filters`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.


<a id="nestedatt--lifecycle_configuration--rules--transition"></a>
### Nested Schema for `lifecycle_configuration.rules.transition`

Optional:

- `storage_class` (String) The storage class to which you want the object to transition.
- `transition_date` (String) Indicates when objects are transitioned to the specified storage class. The date value must be in ISO 8601 format. The time is always midnight UTC.
- `transition_in_days` (Number) Indicates the number of days after creation when objects are transitioned to the specified storage class. If the specified storage class is ``INTELLIGENT_TIERING``, ``GLACIER_IR``, ``GLACIER``, or ``DEEP_ARCHIVE``, valid values are ``0`` or positive integers. If the specified storage class is ``STANDARD_IA`` or ``ONEZONE_IA``, valid values are positive integers greater than ``30``. Be aware that some storage classes have a minimum storage duration and that you're charged for transitioning objects before their minimum storage duration. For more information, see [Constraints and considerations for transitions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html#lifecycle-configuration-constraints) in the *Amazon S3 User Guide*.


<a id="nestedatt--lifecycle_configuration--rules--transitions"></a>
### Nested Schema for `lifecycle_configuration.rules.transitions`

Optional:

- `storage_class` (String) The storage class to which you want the object to transition.
- `transition_date` (String) Indicates when objects are transitioned to the specified storage class. The date value must be in ISO 8601 format. The time is always midnight UTC.
- `transition_in_days` (Number) Indicates the number of days after creation when objects are transitioned to the specified storage class. If the specified storage class is ``INTELLIGENT_TIERING``, ``GLACIER_IR``, ``GLACIER``, or ``DEEP_ARCHIVE``, valid values are ``0`` or positive integers. If the specified storage class is ``STANDARD_IA`` or ``ONEZONE_IA``, valid values are positive integers greater than ``30``. Be aware that some storage classes have a minimum storage duration and that you're charged for transitioning objects before their minimum storage duration. For more information, see [Constraints and considerations for transitions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html#lifecycle-configuration-constraints) in the *Amazon S3 User Guide*.




<a id="nestedatt--logging_configuration"></a>
### Nested Schema for `logging_configuration`

Optional:

- `destination_bucket_name` (String) The name of the bucket where Amazon S3 should store server access log files. You can store log files in any bucket that you own. By default, logs are stored in the bucket where the ``LoggingConfiguration`` property is defined.
- `log_file_prefix` (String) A prefix for all log object keys. If you store log files from multiple Amazon S3 buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket.
- `target_object_key_format` (Attributes) Amazon S3 key format for log objects. Only one format, either PartitionedPrefix or SimplePrefix, is allowed. (see [below for nested schema](#nestedatt--logging_configuration--target_object_key_format))

<a id="nestedatt--logging_configuration--target_object_key_format"></a>
### Nested Schema for `logging_configuration.target_object_key_format`

Optional:

- `partitioned_prefix` (Attributes) Amazon S3 keys for log objects are partitioned in the following format:
  ``[DestinationPrefix][SourceAccountId]/[SourceRegion]/[SourceBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]`` 
 PartitionedPrefix defaults to EventTime delivery when server access logs are delivered. (see [below for nested schema](#nestedatt--logging_configuration--target_object_key_format--partitioned_prefix))
- `simple_prefix` (String) This format defaults the prefix to the given log file prefix for delivering server access log file.

<a id="nestedatt--logging_configuration--target_object_key_format--partitioned_prefix"></a>
### Nested Schema for `logging_configuration.target_object_key_format.partitioned_prefix`

Optional:

- `partition_date_source` (String) Specifies the partition date source for the partitioned prefix. ``PartitionDateSource`` can be ``EventTime`` or ``DeliveryTime``.
 For ``DeliveryTime``, the time in the log file names corresponds to the delivery time for the log files. 
  For ``EventTime``, The logs delivered are for a specific day only. The year, month, and day correspond to the day on which the event occurred, and the hour, minutes and seconds are set to 00 in the key.




<a id="nestedatt--metadata_table_configuration"></a>
### Nested Schema for `metadata_table_configuration`

Optional:

- `s3_tables_destination` (Attributes) The destination information for the metadata table configuration. The destination table bucket must be in the same Region and AWS-account as the general purpose bucket. The specified metadata table name must be unique within the ``aws_s3_metadata`` namespace in the destination table bucket. (see [below for nested schema](#nestedatt--metadata_table_configuration--s3_tables_destination))

<a id="nestedatt--metadata_table_configuration--s3_tables_destination"></a>
### Nested Schema for `metadata_table_configuration.s3_tables_destination`

Optional:

- `table_bucket_arn` (String) The Amazon Resource Name (ARN) for the table bucket that's specified as the destination in the metadata table configuration. The destination table bucket must be in the same Region and AWS-account as the general purpose bucket.
- `table_name` (String) The name for the metadata table in your metadata table configuration. The specified metadata table name must be unique within the ``aws_s3_metadata`` namespace in the destination table bucket.

Read-Only:

- `table_arn` (String) The Amazon Resource Name (ARN) for the metadata table in the metadata table configuration. The specified metadata table name must be unique within the ``aws_s3_metadata`` namespace in the destination table bucket.
- `table_namespace` (String) The table bucket namespace for the metadata table in your metadata table configuration. This value is always ``aws_s3_metadata``.



<a id="nestedatt--metrics_configurations"></a>
### Nested Schema for `metrics_configurations`

Optional:

- `access_point_arn` (String) The access point that was used while performing operations on the object. The metrics configuration only includes objects that meet the filter's criteria.
- `id` (String) The ID used to identify the metrics configuration. This can be any value you choose that helps you identify your metrics configuration.
- `prefix` (String) The prefix that an object must have to be included in the metrics results.
- `tag_filters` (Attributes List) Specifies a list of tag filters to use as a metrics configuration filter. The metrics configuration includes only objects that meet the filter's criteria. (see [below for nested schema](#nestedatt--metrics_configurations--tag_filters))

<a id="nestedatt--metrics_configurations--tag_filters"></a>
### Nested Schema for `metrics_configurations.tag_filters`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.



<a id="nestedatt--notification_configuration"></a>
### Nested Schema for `notification_configuration`

Optional:

- `event_bridge_configuration` (Attributes) Enables delivery of events to Amazon EventBridge. (see [below for nested schema](#nestedatt--notification_configuration--event_bridge_configuration))
- `lambda_configurations` (Attributes List) Describes the LAMlong functions to invoke and the events for which to invoke them. (see [below for nested schema](#nestedatt--notification_configuration--lambda_configurations))
- `queue_configurations` (Attributes List) The Amazon Simple Queue Service queues to publish messages to and the events for which to publish messages. (see [below for nested schema](#nestedatt--notification_configuration--queue_configurations))
- `topic_configurations` (Attributes List) The topic to which notifications are sent and the events for which notifications are generated. (see [below for nested schema](#nestedatt--notification_configuration--topic_configurations))

<a id="nestedatt--notification_configuration--event_bridge_configuration"></a>
### Nested Schema for `notification_configuration.event_bridge_configuration`

Optional:

- `event_bridge_enabled` (Boolean) Enables delivery of events to Amazon EventBridge.


<a id="nestedatt--notification_configuration--lambda_configurations"></a>
### Nested Schema for `notification_configuration.lambda_configurations`

Optional:

- `event` (String) The Amazon S3 bucket event for which to invoke the LAMlong function. For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
- `filter` (Attributes) The filtering rules that determine which objects invoke the AWS Lambda function. For example, you can create a filter so that only image files with a ``.jpg`` extension invoke the function when they are added to the Amazon S3 bucket. (see [below for nested schema](#nestedatt--notification_configuration--lambda_configurations--filter))
- `function` (String) The Amazon Resource Name (ARN) of the LAMlong function that Amazon S3 invokes when the specified event type occurs.

<a id="nestedatt--notification_configuration--lambda_configurations--filter"></a>
### Nested Schema for `notification_configuration.lambda_configurations.filter`

Optional:

- `s3_key` (Attributes) A container for object key name prefix and suffix filtering rules. (see [below for nested schema](#nestedatt--notification_configuration--lambda_configurations--filter--s3_key))

<a id="nestedatt--notification_configuration--lambda_configurations--filter--s3_key"></a>
### Nested Schema for `notification_configuration.lambda_configurations.filter.s3_key`

Optional:

- `rules` (Attributes Set) A list of containers for the key-value pair that defines the criteria for the filter rule. (see [below for nested schema](#nestedatt--notification_configuration--lambda_configurations--filter--s3_key--rules))

<a id="nestedatt--notification_configuration--lambda_configurations--filter--s3_key--rules"></a>
### Nested Schema for `notification_configuration.lambda_configurations.filter.s3_key.rules`

Optional:

- `name` (String) The object key name prefix or suffix identifying one or more objects to which the filtering rule applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
- `value` (String) The value that the filter searches for in object key names.





<a id="nestedatt--notification_configuration--queue_configurations"></a>
### Nested Schema for `notification_configuration.queue_configurations`

Optional:

- `event` (String) The Amazon S3 bucket event about which you want to publish messages to Amazon SQS. For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
- `filter` (Attributes) The filtering rules that determine which objects trigger notifications. For example, you can create a filter so that Amazon S3 sends notifications only when image files with a ``.jpg`` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/user-guide/notification-how-to-filtering.html) in the *Amazon S3 User Guide*. (see [below for nested schema](#nestedatt--notification_configuration--queue_configurations--filter))
- `queue` (String) The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type. FIFO queues are not allowed when enabling an SQS queue as the event notification destination.

<a id="nestedatt--notification_configuration--queue_configurations--filter"></a>
### Nested Schema for `notification_configuration.queue_configurations.filter`

Optional:

- `s3_key` (Attributes) A container for object key name prefix and suffix filtering rules. (see [below for nested schema](#nestedatt--notification_configuration--queue_configurations--filter--s3_key))

<a id="nestedatt--notification_configuration--queue_configurations--filter--s3_key"></a>
### Nested Schema for `notification_configuration.queue_configurations.filter.s3_key`

Optional:

- `rules` (Attributes Set) A list of containers for the key-value pair that defines the criteria for the filter rule. (see [below for nested schema](#nestedatt--notification_configuration--queue_configurations--filter--s3_key--rules))

<a id="nestedatt--notification_configuration--queue_configurations--filter--s3_key--rules"></a>
### Nested Schema for `notification_configuration.queue_configurations.filter.s3_key.rules`

Optional:

- `name` (String) The object key name prefix or suffix identifying one or more objects to which the filtering rule applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
- `value` (String) The value that the filter searches for in object key names.





<a id="nestedatt--notification_configuration--topic_configurations"></a>
### Nested Schema for `notification_configuration.topic_configurations`

Optional:

- `event` (String) The Amazon S3 bucket event about which to send notifications. For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
- `filter` (Attributes) The filtering rules that determine for which objects to send notifications. For example, you can create a filter so that Amazon S3 sends notifications only when image files with a ``.jpg`` extension are added to the bucket. (see [below for nested schema](#nestedatt--notification_configuration--topic_configurations--filter))
- `topic` (String) The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.

<a id="nestedatt--notification_configuration--topic_configurations--filter"></a>
### Nested Schema for `notification_configuration.topic_configurations.filter`

Optional:

- `s3_key` (Attributes) A container for object key name prefix and suffix filtering rules. (see [below for nested schema](#nestedatt--notification_configuration--topic_configurations--filter--s3_key))

<a id="nestedatt--notification_configuration--topic_configurations--filter--s3_key"></a>
### Nested Schema for `notification_configuration.topic_configurations.filter.s3_key`

Optional:

- `rules` (Attributes Set) A list of containers for the key-value pair that defines the criteria for the filter rule. (see [below for nested schema](#nestedatt--notification_configuration--topic_configurations--filter--s3_key--rules))

<a id="nestedatt--notification_configuration--topic_configurations--filter--s3_key--rules"></a>
### Nested Schema for `notification_configuration.topic_configurations.filter.s3_key.rules`

Optional:

- `name` (String) The object key name prefix or suffix identifying one or more objects to which the filtering rule applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
- `value` (String) The value that the filter searches for in object key names.






<a id="nestedatt--object_lock_configuration"></a>
### Nested Schema for `object_lock_configuration`

Optional:

- `object_lock_enabled` (String) Indicates whether this bucket has an Object Lock configuration enabled. Enable ``ObjectLockEnabled`` when you apply ``ObjectLockConfiguration`` to a bucket.
- `rule` (Attributes) Specifies the Object Lock rule for the specified object. Enable this rule when you apply ``ObjectLockConfiguration`` to a bucket. If Object Lock is turned on, bucket settings require both ``Mode`` and a period of either ``Days`` or ``Years``. You cannot specify ``Days`` and ``Years`` at the same time. For more information, see [ObjectLockRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html) and [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html). (see [below for nested schema](#nestedatt--object_lock_configuration--rule))

<a id="nestedatt--object_lock_configuration--rule"></a>
### Nested Schema for `object_lock_configuration.rule`

Optional:

- `default_retention` (Attributes) The default Object Lock retention mode and period that you want to apply to new objects placed in the specified bucket. If Object Lock is turned on, bucket settings require both ``Mode`` and a period of either ``Days`` or ``Years``. You cannot specify ``Days`` and ``Years`` at the same time. For more information about allowable values for mode and period, see [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html). (see [below for nested schema](#nestedatt--object_lock_configuration--rule--default_retention))

<a id="nestedatt--object_lock_configuration--rule--default_retention"></a>
### Nested Schema for `object_lock_configuration.rule.default_retention`

Optional:

- `days` (Number) The number of days that you want to specify for the default retention period. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
- `mode` (String) The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
- `years` (Number) The number of years that you want to specify for the default retention period. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.




<a id="nestedatt--ownership_controls"></a>
### Nested Schema for `ownership_controls`

Optional:

- `rules` (Attributes List) Specifies the container element for Object Ownership rules. (see [below for nested schema](#nestedatt--ownership_controls--rules))

<a id="nestedatt--ownership_controls--rules"></a>
### Nested Schema for `ownership_controls.rules`

Optional:

- `object_ownership` (String) Specifies an object ownership rule.



<a id="nestedatt--public_access_block_configuration"></a>
### Nested Schema for `public_access_block_configuration`

Optional:

- `block_public_acls` (Boolean) Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket. Setting this element to ``TRUE`` causes the following behavior:
  +  PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.
  +  PUT Object calls fail if the request includes a public ACL.
  +  PUT Bucket calls fail if the request includes a public ACL.
  
 Enabling this setting doesn't affect existing policies or ACLs.
- `block_public_policy` (Boolean) Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to ``TRUE`` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. 
 Enabling this setting doesn't affect existing bucket policies.
- `ignore_public_acls` (Boolean) Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting this element to ``TRUE`` causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.
 Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
- `restrict_public_buckets` (Boolean) Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to ``TRUE`` restricts access to this bucket to only AWS-service principals and authorized users within this account if the bucket has a public policy.
 Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.


<a id="nestedatt--replication_configuration"></a>
### Nested Schema for `replication_configuration`

Optional:

- `role` (String) The Amazon Resource Name (ARN) of the IAMlong (IAM) role that Amazon S3 assumes when replicating objects. For more information, see [How to Set Up Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-how-setup.html) in the *Amazon S3 User Guide*.
- `rules` (Attributes List) A container for one or more replication rules. A replication configuration must have at least one rule and can contain a maximum of 1,000 rules. (see [below for nested schema](#nestedatt--replication_configuration--rules))

<a id="nestedatt--replication_configuration--rules"></a>
### Nested Schema for `replication_configuration.rules`

Optional:

- `delete_marker_replication` (Attributes) Specifies whether Amazon S3 replicates delete markers. If you specify a ``Filter`` in your replication configuration, you must also include a ``DeleteMarkerReplication`` element. If your ``Filter`` includes a ``Tag`` element, the ``DeleteMarkerReplication`` ``Status`` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config). 
 For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html). 
  If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations). (see [below for nested schema](#nestedatt--replication_configuration--rules--delete_marker_replication))
- `destination` (Attributes) A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC). (see [below for nested schema](#nestedatt--replication_configuration--rules--destination))
- `filter` (Attributes) A filter that identifies the subset of objects to which the replication rule applies. A ``Filter`` must specify exactly one ``Prefix``, ``TagFilter``, or an ``And`` child element. The use of the filter field indicates that this is a V2 replication configuration. This field isn't supported in a V1 replication configuration.
  V1 replication configuration only supports filtering by key prefix. To filter using a V1 replication configuration, add the ``Prefix`` directly as a child element of the ``Rule`` element. (see [below for nested schema](#nestedatt--replication_configuration--rules--filter))
- `id` (String) A unique identifier for the rule. The maximum value is 255 characters. If you don't specify a value, AWS CloudFormation generates a random ID. When using a V2 replication configuration this property is capitalized as "ID".
- `prefix` (String) An object key name prefix that identifies the object or objects to which the rule applies. The maximum prefix length is 1,024 characters. To include all objects in a bucket, specify an empty string. To filter using a V1 replication configuration, add the ``Prefix`` directly as a child element of the ``Rule`` element.
  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
- `priority` (Number) The priority indicates which rule has precedence whenever two or more replication rules conflict. Amazon S3 will attempt to replicate objects according to all replication rules. However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority. The higher the number, the higher the priority. 
 For more information, see [Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the *Amazon S3 User Guide*.
- `source_selection_criteria` (Attributes) A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects. (see [below for nested schema](#nestedatt--replication_configuration--rules--source_selection_criteria))
- `status` (String) Specifies whether the rule is enabled.

<a id="nestedatt--replication_configuration--rules--delete_marker_replication"></a>
### Nested Schema for `replication_configuration.rules.delete_marker_replication`

Optional:

- `status` (String) Indicates whether to replicate delete markers. Disabled by default.


<a id="nestedatt--replication_configuration--rules--destination"></a>
### Nested Schema for `replication_configuration.rules.destination`

Optional:

- `access_control_translation` (Attributes) Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS-account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS-account that owns the source object. (see [below for nested schema](#nestedatt--replication_configuration--rules--destination--access_control_translation))
- `account` (String) Destination bucket owner account ID. In a cross-account scenario, if you direct Amazon S3 to change replica ownership to the AWS-account that owns the destination bucket by specifying the ``AccessControlTranslation`` property, this is the account ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional Configuration: Change Replica Owner](https://docs.aws.amazon.com/AmazonS3/latest/dev/crr-change-owner.html) in the *Amazon S3 User Guide*.
 If you specify the ``AccessControlTranslation`` property, the ``Account`` property is required.
- `bucket` (String) The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.
- `encryption_configuration` (Attributes) Specifies encryption-related information. (see [below for nested schema](#nestedatt--replication_configuration--rules--destination--encryption_configuration))
- `metrics` (Attributes) A container specifying replication metrics-related settings enabling replication metrics and events. (see [below for nested schema](#nestedatt--replication_configuration--rules--destination--metrics))
- `replication_time` (Attributes) A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a ``Metrics`` block. (see [below for nested schema](#nestedatt--replication_configuration--rules--destination--replication_time))
- `storage_class` (String) The storage class to use when replicating objects, such as S3 Standard or reduced redundancy. By default, Amazon S3 uses the storage class of the source object to create the object replica. 
 For valid values, see the ``StorageClass`` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference*.

<a id="nestedatt--replication_configuration--rules--destination--access_control_translation"></a>
### Nested Schema for `replication_configuration.rules.destination.access_control_translation`

Optional:

- `owner` (String) Specifies the replica ownership. For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) in the *Amazon S3 API Reference*.


<a id="nestedatt--replication_configuration--rules--destination--encryption_configuration"></a>
### Nested Schema for `replication_configuration.rules.destination.encryption_configuration`

Optional:

- `replica_kms_key_id` (String) Specifies the ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket. Amazon S3 uses this key to encrypt replica objects. Amazon S3 only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *Key Management Service Developer Guide*.


<a id="nestedatt--replication_configuration--rules--destination--metrics"></a>
### Nested Schema for `replication_configuration.rules.destination.metrics`

Optional:

- `event_threshold` (Attributes) A container specifying the time threshold for emitting the ``s3:Replication:OperationMissedThreshold`` event. (see [below for nested schema](#nestedatt--replication_configuration--rules--destination--metrics--event_threshold))
- `status` (String) Specifies whether the replication metrics are enabled.

<a id="nestedatt--replication_configuration--rules--destination--metrics--event_threshold"></a>
### Nested Schema for `replication_configuration.rules.destination.metrics.event_threshold`

Optional:

- `minutes` (Number) Contains an integer specifying time in minutes. 
  Valid value: 15



<a id="nestedatt--replication_configuration--rules--destination--replication_time"></a>
### Nested Schema for `replication_configuration.rules.destination.replication_time`

Optional:

- `status` (String) Specifies whether the replication time is enabled.
- `time` (Attributes) A container specifying the time by which replication should be complete for all objects and operations on objects. (see [below for nested schema](#nestedatt--replication_configuration--rules--destination--replication_time--time))

<a id="nestedatt--replication_configuration--rules--destination--replication_time--time"></a>
### Nested Schema for `replication_configuration.rules.destination.replication_time.time`

Optional:

- `minutes` (Number) Contains an integer specifying time in minutes. 
  Valid value: 15




<a id="nestedatt--replication_configuration--rules--filter"></a>
### Nested Schema for `replication_configuration.rules.filter`

Optional:

- `and` (Attributes) A container for specifying rule filters. The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter. For example: 
  +  If you specify both a ``Prefix`` and a ``TagFilter``, wrap these filters in an ``And`` tag.
  +  If you specify a filter based on multiple tags, wrap the ``TagFilter`` elements in an ``And`` tag. (see [below for nested schema](#nestedatt--replication_configuration--rules--filter--and))
- `prefix` (String) An object key name prefix that identifies the subset of objects to which the rule applies.
  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
- `tag_filter` (Attributes) A container for specifying a tag key and value. 
 The rule applies only to objects that have the tag in their tag set. (see [below for nested schema](#nestedatt--replication_configuration--rules--filter--tag_filter))

<a id="nestedatt--replication_configuration--rules--filter--and"></a>
### Nested Schema for `replication_configuration.rules.filter.and`

Optional:

- `prefix` (String) An object key name prefix that identifies the subset of objects to which the rule applies.
- `tag_filters` (Attributes List) An array of tags containing key and value pairs. (see [below for nested schema](#nestedatt--replication_configuration--rules--filter--and--tag_filters))

<a id="nestedatt--replication_configuration--rules--filter--and--tag_filters"></a>
### Nested Schema for `replication_configuration.rules.filter.and.tag_filters`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.



<a id="nestedatt--replication_configuration--rules--filter--tag_filter"></a>
### Nested Schema for `replication_configuration.rules.filter.tag_filter`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.



<a id="nestedatt--replication_configuration--rules--source_selection_criteria"></a>
### Nested Schema for `replication_configuration.rules.source_selection_criteria`

Optional:

- `replica_modifications` (Attributes) A filter that you can specify for selection for modifications on replicas. (see [below for nested schema](#nestedatt--replication_configuration--rules--source_selection_criteria--replica_modifications))
- `sse_kms_encrypted_objects` (Attributes) A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS. (see [below for nested schema](#nestedatt--replication_configuration--rules--source_selection_criteria--sse_kms_encrypted_objects))

<a id="nestedatt--replication_configuration--rules--source_selection_criteria--replica_modifications"></a>
### Nested Schema for `replication_configuration.rules.source_selection_criteria.replica_modifications`

Optional:

- `status` (String) Specifies whether Amazon S3 replicates modifications on replicas.
  *Allowed values*: ``Enabled`` | ``Disabled``


<a id="nestedatt--replication_configuration--rules--source_selection_criteria--sse_kms_encrypted_objects"></a>
### Nested Schema for `replication_configuration.rules.source_selection_criteria.sse_kms_encrypted_objects`

Optional:

- `status` (String) Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.





<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) Name of the object key.
- `value` (String) Value of the tag.


<a id="nestedatt--versioning_configuration"></a>
### Nested Schema for `versioning_configuration`

Optional:

- `status` (String) The versioning state of the bucket.


<a id="nestedatt--website_configuration"></a>
### Nested Schema for `website_configuration`

Optional:

- `error_document` (String) The name of the error document for the website.
- `index_document` (String) The name of the index document for the website.
- `redirect_all_requests_to` (Attributes) The redirect behavior for every request to this bucket's website endpoint.
  If you specify this property, you can't specify any other property. (see [below for nested schema](#nestedatt--website_configuration--redirect_all_requests_to))
- `routing_rules` (Attributes List) Rules that define when a redirect is applied and the redirect behavior. (see [below for nested schema](#nestedatt--website_configuration--routing_rules))

<a id="nestedatt--website_configuration--redirect_all_requests_to"></a>
### Nested Schema for `website_configuration.redirect_all_requests_to`

Optional:

- `host_name` (String) Name of the host where requests are redirected.
- `protocol` (String) Protocol to use when redirecting requests. The default is the protocol that is used in the original request.


<a id="nestedatt--website_configuration--routing_rules"></a>
### Nested Schema for `website_configuration.routing_rules`

Optional:

- `redirect_rule` (Attributes) Container for redirect information. You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return. (see [below for nested schema](#nestedatt--website_configuration--routing_rules--redirect_rule))
- `routing_rule_condition` (Attributes) A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the ``/docs`` folder, redirect to the ``/documents`` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error. (see [below for nested schema](#nestedatt--website_configuration--routing_rules--routing_rule_condition))

<a id="nestedatt--website_configuration--routing_rules--redirect_rule"></a>
### Nested Schema for `website_configuration.routing_rules.redirect_rule`

Optional:

- `host_name` (String) The host name to use in the redirect request.
- `http_redirect_code` (String) The HTTP redirect code to use on the response. Not required if one of the siblings is present.
- `protocol` (String) Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
- `replace_key_prefix_with` (String) The object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix ``docs/`` (objects in the ``docs/`` folder) to ``documents/``, you can set a condition block with ``KeyPrefixEquals`` set to ``docs/`` and in the Redirect set ``ReplaceKeyPrefixWith`` to ``/documents``. Not required if one of the siblings is present. Can be present only if ``ReplaceKeyWith`` is not provided.
  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
- `replace_key_with` (String) The specific object key to use in the redirect request. For example, redirect request to ``error.html``. Not required if one of the siblings is present. Can be present only if ``ReplaceKeyPrefixWith`` is not provided.
  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).


<a id="nestedatt--website_configuration--routing_rules--routing_rule_condition"></a>
### Nested Schema for `website_configuration.routing_rules.routing_rule_condition`

Optional:

- `http_error_code_returned_equals` (String) The HTTP error code when the redirect is applied. In the event of an error, if the error code equals this value, then the specified redirect is applied.
 Required when parent element ``Condition`` is specified and sibling ``KeyPrefixEquals`` is not specified. If both are specified, then both must be true for the redirect to be applied.
- `key_prefix_equals` (String) The object key name prefix when the redirect is applied. For example, to redirect requests for ``ExamplePage.html``, the key prefix will be ``ExamplePage.html``. To redirect request for all pages with the prefix ``docs/``, the key prefix will be ``/docs``, which identifies all objects in the docs/ folder.
 Required when the parent element ``Condition`` is specified and sibling ``HttpErrorCodeReturnedEquals`` is not specified. If both conditions are specified, both must be true for the redirect to be applied.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_s3_bucket.example
  id = "bucket_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_s3_bucket.example "bucket_name"
```