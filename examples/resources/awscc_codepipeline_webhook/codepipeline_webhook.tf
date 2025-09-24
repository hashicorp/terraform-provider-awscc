# Create S3 bucket for artifacts
resource "awscc_s3_bucket" "artifact_store" {
  bucket_name = "example-pipeline-artifacts"
}

# Create IAM role for CodePipeline using standard AWS provider
resource "aws_iam_role" "codepipeline_role" {
  name = "example-codepipeline-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "codepipeline.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "codepipeline_policy" {
  name = "codepipeline-policy"
  role = aws_iam_role.codepipeline_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:GetObjectVersion",
          "s3:GetBucketVersioning",
          "s3:PutObjectAcl",
          "s3:PutObject",
        ]
        Resource = ["*"]
      },
      {
        Effect = "Allow"
        Action = [
          "codebuild:BatchGetBuilds",
          "codebuild:StartBuild"
        ]
        Resource = ["*"]
      }
    ]
  })
}

# Create a CodePipeline with proper structure
resource "awscc_codepipeline_pipeline" "example" {
  name     = "example-pipeline"
  role_arn = aws_iam_role.codepipeline_role.arn

  artifact_store = {
    type     = "S3"
    location = awscc_s3_bucket.artifact_store.id
  }

  stages = [
    {
      name = "Source"
      actions = [
        {
          name      = "Source"
          namespace = "SourceVariables"
          action_type_id = {
            category = "Source"
            owner    = "AWS"
            provider = "S3"
            version  = "1"
          }
          configuration = jsonencode({
            S3Bucket             = awscc_s3_bucket.artifact_store.id
            S3ObjectKey          = "source.zip"
            PollForSourceChanges = false
          })
          output_artifacts = [
            {
              name = "SourceCode"
            }
          ]
        }
      ]
    },
    {
      name = "Approval"
      actions = [
        {
          name = "Approval"
          action_type_id = {
            category = "Approval"
            owner    = "AWS"
            provider = "Manual"
            version  = "1"
          }
          configuration = jsonencode({
            CustomData = "Please review before proceeding"
          })
        }
      ]
    }
  ]
}

resource "awscc_codepipeline_webhook" "example" {
  name = "example-webhook"

  authentication = "GITHUB_HMAC"
  authentication_configuration = {
    secret_token = "examplesecrettokenshouldbereplaced"
  }

  filters = [
    {
      json_path    = "$.ref"
      match_equals = "refs/heads/main"
    }
  ]

  target_pipeline         = awscc_codepipeline_pipeline.example.name
  target_action           = "Source"
  target_pipeline_version = 1

  register_with_third_party = false
}
