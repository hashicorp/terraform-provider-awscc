# Example Personalize Schema for user interactions
resource "awscc_personalize_schema" "example" {
  name = "example-user-interactions-schema"
  schema = jsonencode({
    type      = "record"
    name      = "Interactions"
    namespace = "com.amazonaws.personalize.schema"
    fields = [
      {
        name = "USER_ID"
        type = "string"
      },
      {
        name = "ITEM_ID"
        type = "string"
      },
      {
        name = "TIMESTAMP"
        type = "long"
      },
      {
        name = "EVENT_TYPE"
        type = "string"
      }
    ]
    version = "1.0"
  })
}