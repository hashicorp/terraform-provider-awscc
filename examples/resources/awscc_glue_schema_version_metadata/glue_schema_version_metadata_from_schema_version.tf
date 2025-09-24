resource "awscc_glue_schema_version_metadata" "example" {
  schema_version_id = awscc_glue_schema_version.example.version_id
  key               = "test-key"
  value             = "test-value"
}


resource "awscc_glue_schema_version" "example" {
  schema = {
    schema_arn = awscc_glue_schema.example.arn
  }
  schema_definition = jsonencode(
    {
      type = "record",
      name = "r1",
      fields = [
        {
          name = "f1",
          type = "int"
        },
        { name = "f3",
          type = "string"
        }
      ]
    }
  )
}

resource "awscc_glue_schema" "example" {
  name = "example"
  registry = {
    arn = awscc_glue_registry.example.arn
  }
  data_format   = "AVRO"
  compatibility = "NONE"
  schema_definition = jsonencode(
    {
      type = "record",
      name = "r1",
      fields = [
        {
          name = "f1",
          type = "int"
        },
        {
          name = "f2",
          type = "string"
        }
      ]
    }
  )
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}