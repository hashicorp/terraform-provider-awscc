defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

resource_schema "aws_logs_log_group" {
  cloudformation_type_name = "AWS::Logs::LogGroup"
}