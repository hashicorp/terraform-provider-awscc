resource "awscc_bedrock_blueprint" "example" {
  blueprint_name = "example-blueprint"
  type           = "DOCUMENT"
  schema = jsonencode({
    "$schema" : "http://json-schema.org/draft-07/schema#",
    "description" : "default",
    "class" : "default",
    "type" : "object",
    "definitions" : {},
    "properties" : {
      "gross_pay_this_period" : {
        "type" : "number",
        "inferenceType" : "explicit",
        "instruction" : "The gross pay for this pay period from the Earnings table"
      },
      "net_pay" : {
        "type" : "number",
        "inferenceType" : "explicit",
        "instruction" : "The net pay for this pay period from the bottom of the document"
      }
    }
  })
}