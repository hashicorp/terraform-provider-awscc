resource "awscc_bedrock_data_automation_project" "example" {
  project_name        = "example-project"
  project_description = "example-description"
  standard_output_configuration = {
    document = {
      output_format = {
        text_format = {
          types = ["PLAIN_TEXT"]
        }
        additional_file_format = {
          state = "DISABLED"
        }
      }
    }
  }
  custom_output_configuration = {
    blueprints = [
      {
        blueprint_arn = awscc_bedrock_blueprint.example.blueprint_arn
      }
    ]
  }
  override_configuration = {
    document = {
      splitter = {
        state = "ENABLED"
      }
    }
  }
}