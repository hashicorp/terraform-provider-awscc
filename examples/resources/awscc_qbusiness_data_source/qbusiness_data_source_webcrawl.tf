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
