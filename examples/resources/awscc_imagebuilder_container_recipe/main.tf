# Required data sources for dynamic values
data "aws_caller_identity" "current" {}

# IAM policy for component access
data "aws_iam_policy_document" "component_policy" {
  statement {
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "imagebuilder:GetComponent",
      "imagebuilder:ListComponents"
    ]
    resources = ["*"]
  }
}

resource "awscc_ecr_repository" "example" {
  repository_name = "example-container-repo"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_imagebuilder_component" "example" {
  name        = "example-component"
  platform    = "Linux"
  version     = "1.0.0"
  data        = <<EOF
schemaVersion: 1.0
phases:
  - name: build
    steps:
      - name: Example
        action: ExecuteBash
        inputs:
          commands:
            - echo "Hello from component"
EOF
  description = "Example component for container recipe"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_imagebuilder_container_recipe" "example" {
  name        = "example-container-recipe"
  version     = "1.0.0"
  description = "Example container recipe"

  components = [{
    component_arn = awscc_imagebuilder_component.example.arn
  }]

  container_type = "DOCKER"
  parent_image   = "amazonlinux:2"

  target_repository = {
    service         = "ECR"
    repository_name = awscc_ecr_repository.example.repository_name
  }

  dockerfile_template_data = <<-EOT
    FROM {{{ imagebuilder:parentImage }}}
    {{{ imagebuilder:components }}}
  EOT

  working_directory = "/tmp"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}