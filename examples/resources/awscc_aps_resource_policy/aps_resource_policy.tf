# Create APS workspace
resource "awscc_aps_workspace" "example" {
  alias = "example-workspace"
  tags = [
    {
      key   = "Name"
      value = "example-workspace"
    },
    {
      key   = "Environment"
      value = "dev"
    }
  ]
}

# Create a resource policy for the APS workspace
resource "awscc_aps_resource_policy" "example" {
  workspace_arn = awscc_aps_workspace.example.arn

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::123456789012:root"
        }
        Action = [
          "aps:QueryMetrics",
          "aps:GetLabels",
          "aps:GetMetricMetadata",
          "aps:GetSeries"
        ]
        Resource = awscc_aps_workspace.example.arn
      }
    ]
  })
}
