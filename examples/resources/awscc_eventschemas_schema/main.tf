# Create the schema registry first
resource "awscc_eventschemas_registry" "example" {
  registry_name = "example-registry"
  description   = "Example registry for schema"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the schema
resource "awscc_eventschemas_schema" "example" {
  content = jsonencode({
    openapi = "3.0.0"
    info = {
      title   = "ExampleSchema"
      version = "1.0.0"
    }
    paths = {}
    components = {
      schemas = {
        ExampleEvent = {
          type = "object"
          properties = {
            id = {
              type = "string"
            }
            message = {
              type = "string"
            }
            timestamp = {
              type   = "string"
              format = "date-time"
            }
          }
          required = ["id", "message", "timestamp"]
        }
      }
    }
  })
  registry_name = awscc_eventschemas_registry.example.registry_name
  type          = "OpenApi3"
  schema_name   = "example-schema"
  description   = "Example schema for EventBridge events"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}