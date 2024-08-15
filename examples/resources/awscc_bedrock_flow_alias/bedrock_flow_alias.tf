resource "awscc_bedrock_flow_alias" "example" {
  name        = "example"
  flow_arn    = "arn:aws:bedrock:us-east-1:123456789012:flow/ASERNAQ7HX"
  description = "example"
  routing_configuration = [
    {
      flow_version = awscc_bedrock_flow_version.example.version
    }
  ]
}

resource "awscc_bedrock_flow_version" "example" {
  flow_arn    = "arn:aws:bedrock:us-east-1:123456789012:flow/ASERNAQ7HX"
  description = "example"
}
