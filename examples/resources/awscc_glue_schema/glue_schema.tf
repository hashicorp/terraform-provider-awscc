resource "awscc_glue_schema" "example" {
  name        = "example"
  description = "Glue schema example"
  registry = {
    arn = awscc_glue_registry.example.arn
  }

  data_format       = "AVRO"
  compatibility     = "NONE"
  schema_definition = "{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_glue_registry" "example" {
  name        = "example-registry"
  description = "Glue registry example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
