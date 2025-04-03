resource "awscc_bedrockdataautomation_project" "example" {
  project_name = "example-project"
  custom_output_configuration = {
    blueprints = [
      {
        blueprint_arn = "arn:aws:bedrock:us-west-2:<AWS_ACCOUNT_ID>:blueprint/<BLUEPRINT_ID>"
      }
    ]
  }
  kms_encryption_context = ""
  kms_key_id             = ""
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
      outputFormat = {
        textFormat = {
          types = ["PLAIN_TEXT"]
        }
        additionalFileFormat = {
          state = "ENABLED"
        }
      }
    }
  }
  tags = ""
}