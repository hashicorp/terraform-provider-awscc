---
page_title: "awscc_qbusiness_data_source Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::QBusiness::DataSource Resource Type
---

# awscc_qbusiness_data_source (Resource)

Definition of AWS::QBusiness::DataSource Resource Type

## Example Usage

### QBusiness data source of type S3 with the IAM role specifications for S3 access.

```terraform
resource "awscc_qbusiness_data_source" "exaple" {
  application_id = awscc_qbusiness_application.example.application_id
  display_name   = "example_q_data_source"
  index_id       = awscc_qbusiness_index.example.index_id
  role_arn       = awscc_iam_role.ds.arn
  configuration = jsonencode(
    {
      type     = "S3"
      version  = "1.0.0"
      syncMode = "FORCED_FULL_CRAWL"
      connectionConfiguration = {
        repositoryEndpointMetadata = {
          BucketName = var.bucket_name
        }
      }
      additionalProperties = {
        inclusionPrefixes = ["docs/"]
      }
      repositoryConfigurations = {
        document = {
          fieldMappings = [
            {
              dataSourceFieldName = "s3_document_id"
              indexFieldType      = "STRING"
              indexFieldName      = "s3_document_id"
            }
          ]
        }
      }
    }
  )
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "example" {
  role_name   = "QBusiness-DataSource-Role"
  description = "QBusiness Data source role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowsAmazonQToAssumeRoleForServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "qbusiness.amazonaws.com"
        }
        Action = [
          "sts:AssumeRole"
        ]
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "sample_iam_role_policy"
  role_name   = awscc_iam_role.example.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "s3:GetObject"
        Resource = "arn:aws:s3:::${var.bucket_name}/*"
      },
      {
        Effect   = "Allow"
        Action   = "s3:ListBucket"
        Resource = "arn:aws:s3:::${var.bucket_name}"
      },
      {
        Effect = "Allow"
        Action = [
          "qbusiness:BatchPutDocument",
          "qbusiness:BatchDeleteDocument"
        ]
        Resource = "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}/index/${awscc_qbusiness_index.example.id}"
      },
      {
        Effect = "Allow"
        Action = ["qbusiness:PutGroup",
          "qbusiness:CreateUser",
          "qbusiness:DeleteGroup",
          "qbusiness:UpdateUser",
        "qbusiness:ListGroups"]
        Resource = [
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}",
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}/index/${awscc_qbusiness_index.example.id}",
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}/index/${awscc_qbusiness_index.example.id}/data-source/*"
        ]
      }
    ]
  })
}

variable "bucket_name" {
  type        = string
  description = "Name of the bucket to be used as the data source input"
}

data "aws_caller_identity" "current" {}

data "aws_region" "current" {}
```

### QBusiness data source of type WebCrawler

```terraform
# Get current AWS account ID and region for reference
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_ssoadmin_instances" "example" {}


# Create an Amazon Q Business application
resource "awscc_qbusiness_application" "example" {
  display_name = "webcrawler-example"
  description  = "Amazon Q Business application with web crawler data source"
  role_arn     = awscc_iam_role.example.arn

  identity_center_instance_arn = data.aws_ssoadmin_instances.example.arns[0]
  attachments_configuration = {
    attachments_control_mode = "ENABLED"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an IAM role for Amazon Q Business service
resource "awscc_iam_role" "example" {
  role_name   = "QBusiness-DataSource-Role"
  description = "QBusiness Data source role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowsAmazonQToAssumeRoleForServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "qbusiness.amazonaws.com"
        }
        Action = [
          "sts:AssumeRole"
        ]
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "sample_iam_role_policy"
  role_name   = awscc_iam_role.example.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "qbusiness:BatchPutDocument",
          "qbusiness:BatchDeleteDocument"
        ]
        Resource = "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.application_id}/index/${awscc_qbusiness_index.example.index_id}"
      },
      {
        Effect = "Allow"
        Action = ["qbusiness:PutGroup",
          "qbusiness:CreateUser",
          "qbusiness:DeleteGroup",
          "qbusiness:UpdateUser",
        "qbusiness:ListGroups"]
        Resource = [
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.application_id}",
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.application_id}/index/${awscc_qbusiness_index.example.index_id}",
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.application_id}/index/${awscc_qbusiness_index.example.index_id}/data-source/*"
        ]
      }
    ]
  })
}

# Create an index for the data source
resource "awscc_qbusiness_index" "example" {
  application_id = awscc_qbusiness_application.example.id
  display_name   = "webcrawler-example-Index"
  description    = "Index for web crawler data source"

  document_attribute_configurations = [
    {
      name        = "source_uri"
      type        = "STRING"
      search      = "ENABLED"
      displayable = true
    },
    {
      name        = "last_updated_at"
      type        = "DATE"
      search      = "ENABLED"
      displayable = true
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Amazon Q Business data source with web crawler configuration
resource "awscc_qbusiness_data_source" "web_crawler" {
  application_id = awscc_qbusiness_application.example.application_id
  index_id       = awscc_qbusiness_index.example.index_id
  display_name   = "webcrawler-ds"
  description    = "Web crawler data source for indexing website content"
  role_arn       = awscc_iam_role.example.arn

  configuration = jsonencode({
    type     = "WEBCRAWLERV2",
    syncMode = "FULL_CRAWL",
    connectionConfiguration = {
      repositoryEndpointMetadata = {
        seedUrlConnections = [
          { seedUrl = "https://en.wikipedia.org/wiki/AWS" }
        ]
      }
    }
    additionalProperties = {
      rateLimit      = "300"
      maxFileSize    = "50"
      crawlDepth     = "2"
      crawlSubDomain = false
      crawlAllDomain = false
      maxLinksPerUrl = "100"
      honorRobots    = false
    }
    repositoryConfigurations = {
      webPage = {
        fieldMappings = [
          {
            dataSourceFieldName = "webs_id"
            indexFieldType      = "LONG"
            indexFieldName      = "webs_id"
          }
        ]
      }
    }
    }
  )
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String)
- `configuration` (String)
- `display_name` (String)
- `index_id` (String)

### Optional

- `description` (String)
- `document_enrichment_configuration` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration))
- `media_extraction_configuration` (Attributes) (see [below for nested schema](#nestedatt--media_extraction_configuration))
- `role_arn` (String)
- `sync_schedule` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `vpc_configuration` (Attributes) (see [below for nested schema](#nestedatt--vpc_configuration))

### Read-Only

- `created_at` (String)
- `data_source_arn` (String)
- `data_source_id` (String)
- `id` (String) Uniquely identifies the resource.
- `status` (String)
- `type` (String)
- `updated_at` (String)

<a id="nestedatt--document_enrichment_configuration"></a>
### Nested Schema for `document_enrichment_configuration`

Optional:

- `inline_configurations` (Attributes List) (see [below for nested schema](#nestedatt--document_enrichment_configuration--inline_configurations))
- `post_extraction_hook_configuration` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--post_extraction_hook_configuration))
- `pre_extraction_hook_configuration` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--pre_extraction_hook_configuration))

<a id="nestedatt--document_enrichment_configuration--inline_configurations"></a>
### Nested Schema for `document_enrichment_configuration.inline_configurations`

Optional:

- `condition` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--inline_configurations--condition))
- `document_content_operator` (String)
- `target` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--inline_configurations--target))

<a id="nestedatt--document_enrichment_configuration--inline_configurations--condition"></a>
### Nested Schema for `document_enrichment_configuration.inline_configurations.condition`

Optional:

- `key` (String)
- `operator` (String)
- `value` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--inline_configurations--condition--value))

<a id="nestedatt--document_enrichment_configuration--inline_configurations--condition--value"></a>
### Nested Schema for `document_enrichment_configuration.inline_configurations.condition.value`

Optional:

- `date_value` (String)
- `long_value` (Number)
- `string_list_value` (List of String)
- `string_value` (String)



<a id="nestedatt--document_enrichment_configuration--inline_configurations--target"></a>
### Nested Schema for `document_enrichment_configuration.inline_configurations.target`

Optional:

- `attribute_value_operator` (String)
- `key` (String)
- `value` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--inline_configurations--target--value))

<a id="nestedatt--document_enrichment_configuration--inline_configurations--target--value"></a>
### Nested Schema for `document_enrichment_configuration.inline_configurations.target.value`

Optional:

- `date_value` (String)
- `long_value` (Number)
- `string_list_value` (List of String)
- `string_value` (String)




<a id="nestedatt--document_enrichment_configuration--post_extraction_hook_configuration"></a>
### Nested Schema for `document_enrichment_configuration.post_extraction_hook_configuration`

Optional:

- `invocation_condition` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--post_extraction_hook_configuration--invocation_condition))
- `lambda_arn` (String)
- `role_arn` (String)
- `s3_bucket_name` (String)

<a id="nestedatt--document_enrichment_configuration--post_extraction_hook_configuration--invocation_condition"></a>
### Nested Schema for `document_enrichment_configuration.post_extraction_hook_configuration.invocation_condition`

Optional:

- `key` (String)
- `operator` (String)
- `value` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--post_extraction_hook_configuration--invocation_condition--value))

<a id="nestedatt--document_enrichment_configuration--post_extraction_hook_configuration--invocation_condition--value"></a>
### Nested Schema for `document_enrichment_configuration.post_extraction_hook_configuration.invocation_condition.value`

Optional:

- `date_value` (String)
- `long_value` (Number)
- `string_list_value` (List of String)
- `string_value` (String)




<a id="nestedatt--document_enrichment_configuration--pre_extraction_hook_configuration"></a>
### Nested Schema for `document_enrichment_configuration.pre_extraction_hook_configuration`

Optional:

- `invocation_condition` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--pre_extraction_hook_configuration--invocation_condition))
- `lambda_arn` (String)
- `role_arn` (String)
- `s3_bucket_name` (String)

<a id="nestedatt--document_enrichment_configuration--pre_extraction_hook_configuration--invocation_condition"></a>
### Nested Schema for `document_enrichment_configuration.pre_extraction_hook_configuration.invocation_condition`

Optional:

- `key` (String)
- `operator` (String)
- `value` (Attributes) (see [below for nested schema](#nestedatt--document_enrichment_configuration--pre_extraction_hook_configuration--invocation_condition--value))

<a id="nestedatt--document_enrichment_configuration--pre_extraction_hook_configuration--invocation_condition--value"></a>
### Nested Schema for `document_enrichment_configuration.pre_extraction_hook_configuration.invocation_condition.value`

Optional:

- `date_value` (String)
- `long_value` (Number)
- `string_list_value` (List of String)
- `string_value` (String)





<a id="nestedatt--media_extraction_configuration"></a>
### Nested Schema for `media_extraction_configuration`

Optional:

- `image_extraction_configuration` (Attributes) (see [below for nested schema](#nestedatt--media_extraction_configuration--image_extraction_configuration))

<a id="nestedatt--media_extraction_configuration--image_extraction_configuration"></a>
### Nested Schema for `media_extraction_configuration.image_extraction_configuration`

Optional:

- `image_extraction_status` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--vpc_configuration"></a>
### Nested Schema for `vpc_configuration`

Optional:

- `security_group_ids` (List of String)
- `subnet_ids` (List of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_qbusiness_data_source.example "application_id|data_source_id|index_id"
```
