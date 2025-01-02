data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["entityresolution.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "id_namespace" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:GetObjectAttributes",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::example-bucket",
      "arn:aws:s3:::example-bucket/*"
    ]
  }
}

resource "awscc_iam_role" "id_namespace" {
  role_name                   = "AWSSCCEntityResolutionIdNamespaceRole"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "id_namespace" {
  policy_name     = "AWSSCCEntityResolutionIdNamespacePolicy"
  role_name       = awscc_iam_role.id_namespace.role_name
  policy_document = data.aws_iam_policy_document.id_namespace.json
}

resource "awscc_entityresolution_id_namespace" "example" {
  id_namespace_name = "example-namespace"
  description       = "Example IdNamespace created with AWSCC provider"
  type              = "SOURCE"
  role_arn          = awscc_iam_role.id_namespace.arn

  input_source_config = [{
    input_source_arn = "arn:aws:glue:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:table/example-database/example-table"
    schema_name      = "schema1"
  }]

  id_mapping_workflow_properties = [{
    id_mapping_type = "RULE_BASED"
    rule_based_properties = {
      attribute_matching_model = "ONE_TO_ONE"
      rule_definition_types    = ["SOURCE"]
      record_matching_models   = ["ONE_SOURCE_TO_ONE_TARGET"]
      rules = [{
        rule_name     = "exact-match"
        matching_keys = ["id", "email"]
      }]
    }
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}