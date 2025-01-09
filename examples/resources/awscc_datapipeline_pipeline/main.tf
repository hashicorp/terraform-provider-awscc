data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_s3_bucket" "pipeline_logs" {
  bucket_name = "datapipeline-logs-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

# Pipeline execution role
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type = "Service"
      identifiers = [
        "datapipeline.amazonaws.com",
        "elasticmapreduce.amazonaws.com"
      ]
    }
    actions = ["sts:AssumeRole"]
  }
}

# Pipeline role policy
data "aws_iam_policy_document" "pipeline_role_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:*",
      "ec2:*",
      "iam:PassRole",
      "iam:ListRolePolicies",
      "iam:GetRolePolicy",
      "iam:ListInstanceProfiles",
      "rds:Describe*",
      "redshift:DescribeClusters",
      "redshift:DescribeClusterSecurityGroups",
      "sns:*",
      "sqs:*",
      "cloudwatch:*",
      "elasticmapreduce:*"
    ]
    resources = ["*"]
  }
}

# EC2 instance role assume policy
data "aws_iam_policy_document" "ec2_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

# EC2 instance role policy
data "aws_iam_policy_document" "ec2_role_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:*",
      "cloudwatch:*",
      "datapipeline:*",
      "dynamodb:*",
      "ec2:Describe*",
      "elasticmapreduce:AddJobFlowSteps",
      "elasticmapreduce:Describe*",
      "elasticmapreduce:ListInstance*",
      "rds:Describe*",
      "redshift:DescribeClusters",
      "redshift:DescribeClusterSecurityGroups",
      "sdb:*",
      "sns:*",
      "sqs:*"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "pipeline_role" {
  role_name                   = "DataPipelineExampleRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  policies = [
    {
      policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.pipeline_role_policy.json))
      policy_name     = "DataPipelineExamplePolicy"
    }
  ]
}

resource "awscc_iam_role" "ec2_role" {
  role_name                   = "DataPipelineEC2ExampleRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.ec2_assume_role.json))
  policies = [
    {
      policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.ec2_role_policy.json))
      policy_name     = "DataPipelineEC2ExamplePolicy"
    }
  ]
}

# Create the instance profile - name must match the role name
resource "awscc_iam_instance_profile" "ec2_profile" {
  instance_profile_name = awscc_iam_role.ec2_role.role_name
  roles                 = [awscc_iam_role.ec2_role.role_name]
}

resource "awscc_datapipeline_pipeline" "example" {
  name        = "example-pipeline"
  description = "Example Data Pipeline using AWSCC provider"

  parameter_objects = [
    {
      id = "myShellCmd"
      attributes = [
        {
          key          = "type"
          string_value = "String"
        },
        {
          key          = "description"
          string_value = "Shell command to run"
        }
      ]
    }
  ]

  parameter_values = [
    {
      id           = "myShellCmd"
      string_value = "echo 'Hello from Data Pipeline'"
    }
  ]

  pipeline_objects = [
    {
      id   = "Default"
      name = "Default"
      fields = [
        {
          key          = "type"
          string_value = "Default"
        },
        {
          key          = "scheduleType"
          string_value = "ondemand"
        },
        {
          key          = "role"
          string_value = awscc_iam_role.pipeline_role.role_name
        },
        {
          key          = "pipelineLogUri"
          string_value = "s3://${awscc_s3_bucket.pipeline_logs.bucket_name}/logs/"
        }
      ]
    },
    {
      id   = "MyEC2Resource"
      name = "MyEC2Resource"
      fields = [
        {
          key          = "type"
          string_value = "Ec2Resource"
        },
        {
          key          = "terminateAfter"
          string_value = "1 Hour"
        },
        {
          key          = "role"
          string_value = awscc_iam_role.pipeline_role.role_name
        },
        {
          key          = "resourceRole"
          string_value = awscc_iam_role.ec2_role.role_name
        }
      ]
    },
    {
      id   = "ShellCommandActivity"
      name = "ShellCommandActivity"
      fields = [
        {
          key          = "type"
          string_value = "ShellCommandActivity"
        },
        {
          key       = "runsOn"
          ref_value = "MyEC2Resource"
        },
        {
          key          = "command"
          string_value = "#{myShellCmd}"
        }
      ]
    }
  ]

  pipeline_tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]

  # Set to true to activate the pipeline immediately
  activate = false

  depends_on = [
    awscc_s3_bucket.pipeline_logs,
    awscc_iam_instance_profile.ec2_profile
  ]
}