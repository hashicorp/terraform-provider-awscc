# Get the current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create S3 bucket for artifacts
resource "awscc_s3_bucket" "artifacts" {
  bucket_name = "codepipeline-artifacts-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# Create IAM role for CodePipeline
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["codepipeline.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "codepipeline_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:GetObjectVersion",
      "s3:GetBucketVersioning",
      "s3:PutObject"
    ]
    resources = [
      "${awscc_s3_bucket.artifacts.arn}",
      "${awscc_s3_bucket.artifacts.arn}/*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "codecommit:CancelUploadArchive",
      "codecommit:GetBranch",
      "codecommit:GetCommit",
      "codecommit:GetUploadArchiveStatus",
      "codecommit:UploadArchive"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "codebuild:BatchGetBuilds",
      "codebuild:StartBuild"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "codepipeline_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "Role for CodePipeline service"
  managed_policy_arns         = []
  max_session_duration        = 3600
  path                        = "/service-role/"
  policies = [
    {
      policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.codepipeline_policy.json))
      policy_name     = "codepipeline-policy"
    }
  ]
  role_name = "AWSCodePipelineServiceRole"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create CodePipeline
resource "awscc_codepipeline_pipeline" "example" {
  name     = "example-pipeline"
  role_arn = awscc_iam_role.codepipeline_role.arn

  artifact_store = {
    location = awscc_s3_bucket.artifacts.id
    type     = "S3"
  }

  stages = [
    {
      name = "Source"
      actions = [
        {
          name = "Source"
          action_type_id = {
            category = "Source"
            owner    = "AWS"
            provider = "CodeCommit"
            version  = "1"
          }
          configuration = jsonencode({
            RepositoryName = "example-repo"
            BranchName     = "main"
          })
          output_artifacts = [
            {
              name = "SourceCode"
            }
          ]
          role_arn  = awscc_iam_role.codepipeline_role.arn
          run_order = 1
        }
      ]
    },
    {
      name = "Build"
      actions = [
        {
          name = "Build"
          action_type_id = {
            category = "Build"
            owner    = "AWS"
            provider = "CodeBuild"
            version  = "1"
          }
          configuration = jsonencode({
            ProjectName = "example-project"
          })
          input_artifacts = [
            {
              name = "SourceCode"
            }
          ]
          output_artifacts = [
            {
              name = "BuildOutput"
            }
          ]
          role_arn  = awscc_iam_role.codepipeline_role.arn
          run_order = 1
        }
      ]
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}