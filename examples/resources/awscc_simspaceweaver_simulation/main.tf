data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

locals {
  bucket_name = "simspaceweaver-simulation-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# S3 bucket for simulation schema and snapshots
resource "aws_s3_bucket" "simulation" {
  bucket        = local.bucket_name
  force_destroy = true
}

# Schema file for simulation
resource "aws_s3_object" "schema" {
  bucket = aws_s3_bucket.simulation.id
  key    = "simulation/schema.json"
  content = jsonencode({
    version = "1.0"
    simulation = {
      name = "example"
    }
  })
}

# IAM role for SimSpaceWeaver
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["simspaceweaver.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

# IAM permissions required for SimSpaceWeaver
data "aws_iam_policy_document" "simulation" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::${local.bucket_name}",
      "arn:aws:s3:::${local.bucket_name}/*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/simspaceweaver/*"
    ]
  }
}

resource "aws_iam_role" "simulation" {
  name               = "simspaceweaver-simulation-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json

  inline_policy {
    name   = "simulation-policy"
    policy = data.aws_iam_policy_document.simulation.json
  }
}

# SimSpaceWeaver simulation
resource "awscc_simspaceweaver_simulation" "example" {
  name     = "example-simulation"
  role_arn = aws_iam_role.simulation.arn

  schema_s3_location = {
    bucket_name = aws_s3_bucket.simulation.id
    object_key  = aws_s3_object.schema.key
  }

  snapshot_s3_location = {
    bucket_name = aws_s3_bucket.simulation.id
    object_key  = "simulation/snapshots/"
  }
}