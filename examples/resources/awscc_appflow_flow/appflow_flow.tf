resource "awscc_appflow_flow" "example" {
  flow_name = "example"

  source_flow_config = {
    connector_type = "S3"
    source_connector_properties = {
      s3 = {
        bucket_name   = var.source_bucket_name
        bucket_prefix = "example"
      }
    }
  }

  destination_flow_config_list = [{
    connector_type = "S3"
    destination_connector_properties = {
      s3 = {
        bucket_name = var.destination_bucket_name

        s3_output_format_config = {
          prefix_config = {
            prefix_type = "PATH"
          }
        }
      }
    }
    }
  ]

  tasks = [
    {
      source_fields     = ["ExampleField"]
      destination_field = "exampleField"
      task_type         = "Map"

      connector_operator = {
        s3 = "NO_OP"
      }
    }
  ]

  trigger_config = {
    trigger_type = "OnDemand"
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}

variable "source_bucket_name" {
  type = string
}

variable "destination_bucket_name" {
  type = string
}
