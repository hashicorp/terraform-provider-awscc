# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create the DataZone connection
resource "awscc_datazone_connection" "example" {
  domain_identifier      = "dzd_example123" # Replace with your domain ID
  environment_identifier = "dze_example123" # Replace with your environment ID
  name                   = "example-connection"
  description            = "Example DataZone connection"

  aws_location = {
    aws_account_id = data.aws_caller_identity.current.account_id
    aws_region     = data.aws_region.current.name
  }

  props = {
    glue_properties = {
      glue_connection_input = {
        name            = "example-glue-connection"
        description     = "Example Glue connection"
        connection_type = "JDBC"
        connection_properties = {
          "JDBC_CONNECTION_URL" = "jdbc:postgresql://example-host:5432/example-db"
          "SECRET_ID"           = "example-secret"
        }
      }
    }
  }
}