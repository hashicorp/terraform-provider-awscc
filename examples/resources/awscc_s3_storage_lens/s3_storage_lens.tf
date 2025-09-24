resource "awscc_s3_storage_lens" "example" {

  storage_lens_configuration = {
    is_enabled = true
    id         = "example"
    account_level = {
      activity_metrics = {
        is_enabled = true
      }
      advanced_cost_optimization_metrics = {
        is_enabled = true
      }
      advanced_data_protection_metrics = {
        is_enabled = true
      }
      detailed_status_codes_metrics = {
        is_enabled = true
      }

      bucket_level = {
        activity_metrics = {
          is_enabled = true
        }
        advanced_cost_optimization_metrics = {
          is_enabled = true
        }
        advanced_data_protection_metrics = {
          is_enabled = true
        }
        detailed_status_codes_metrics = {
          is_enabled = true
        }
        prefix_level = {
          storage_metrics = {
            is_enabled = true
            selection_criteria = {
              delimiter                    = "/"
              max_depth                    = 5
              min_storage_bytes_percentage = 1.23
            }
          }
        }
      }
    }
    exclude = {
      buckets = [
        "arn:aws:s3:::example_bucket_1",
        "arn:aws:s3:::example_bucket_2",
      ]
    }
    data_export = {
      s3_bucket_destination = {
        account_id            = data.aws_caller_identity.current.account_id
        arn                   = "arn:aws:s3:::destination_bucket"
        output_schema_version = "V_1"
        format                = "CSV"
        prefix                = "/"
        cloudwatch_metrics = {
          is_enabled = true
        }
      }
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

data "aws_caller_identity" "current" {}
