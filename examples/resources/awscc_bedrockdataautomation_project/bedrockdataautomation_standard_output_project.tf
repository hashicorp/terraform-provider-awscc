resource "awscc_bedrock_data_automation_project" "example" {
  project_name                = "12345647891-example-project"
  override_configuration = {
    document = {
      splitter = {
        state = "ENABLED"
      }
    }
  }
  project_description = "example-description"
  standard_output_configuration = {
      document = {
            extraction = {
                granularity = {
                    types = ["DOCUMENT"]
                }
                bounding_box = {
                    state = "ENABLED"
                }
            }
            generative_field = {
                state = "DISABLED"
            }
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
}