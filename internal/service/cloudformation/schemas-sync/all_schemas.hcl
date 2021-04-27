meta_schema {
  source      = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-schema/v2.0.4/src/main/resources/schema/provider.definition.schema.v1.json"
  destination = "provider.definition.schema.v1.json"
}

resource_schema "aws_logs_log_group" {
  source      = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-providers-logs/8b9229f78b832800b7fb6c1165bcb3893f44b856/aws-logs-loggroup/aws-logs-loggroup.json"
  destination = "aws_logs_log_group.cf-resource-schema.json"
}

resource_schema "aws_athena_workgroup" {
  source      = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-providers-athena/a2b8f560891e1256fd1ba965e9ffe3dcddde1855/workgroup/aws-athena-workgroup.json"
  destination = "aws_athena_workgroup.cf-resource-schema.json"
}
