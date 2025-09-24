resource "awscc_qbusiness_web_experience" "example" {
  application_id              = awscc_qbusiness_application.example.application_id
  role_arn                    = awscc_iam_role.example.arn
  sample_prompts_control_mode = "ENABLED"
  subtitle                    = "Drop a file and ask questions"
  title                       = "Sample Amazon Q Business App"
  welcome_message             = "Welcome, please enter your questions"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "example" {
  role_name   = "Amazon-QBusiness-WebExperience-Role"
  description = "Grants permissions to AWS Services and Resources used or managed by Amazon Q Business"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "QBusinessTrustPolicy"
        Effect = "Allow"
        Principal = {
          Service = "application.qbusiness.amazonaws.com"
        }
        Action = [
          "sts:AssumeRole",
          "sts:SetContext"
        ]
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
          ArnEquals = {
            "aws:SourceArn" = awscc_qbusiness_application.example.application_arn
          }
        }
      }
    ]
  })
  policies = [
    {
      policy_name = "qbusiness_policy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Sid    = "QBusinessConversationPermission"
            Effect = "Allow"
            Action = [
              "qbusiness:Chat",
              "qbusiness:ChatSync",
              "qbusiness:ListMessages",
              "qbusiness:ListConversations",
              "qbusiness:DeleteConversation",
              "qbusiness:PutFeedback",
              "qbusiness:GetWebExperience",
              "qbusiness:GetApplication",
              "qbusiness:ListPlugins",
              "qbusiness:GetChatControlsConfiguration"
            ]
            Resource = awscc_qbusiness_application.example.application_arn
          }
        ]
      })
      tags = [
        {
          key   = "Modified By"
          value = "AWSCC"
        }
      ]
    }
  ]
}

data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

data "aws_partition" "current" {}
