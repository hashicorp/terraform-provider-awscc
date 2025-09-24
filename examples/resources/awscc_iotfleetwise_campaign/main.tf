# Get current AWS Account ID
data "aws_caller_identity" "current" {}

# Get current AWS Region
data "aws_region" "current" {}

# Create a signal catalog first
resource "awscc_iotfleetwise_signal_catalog" "example" {
  name        = "ExampleSignalCatalog"
  description = "Example signal catalog for FleetWise"
  nodes = [
    {
      branch = {
        fully_qualified_name = "Vehicle"
      }
    },
    {
      sensor = {
        data_type            = "DOUBLE"
        fully_qualified_name = "Vehicle.EngineSensor"
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a fleet
resource "awscc_iotfleetwise_fleet" "example" {
  fleet_id           = "examplefleet"
  description        = "Example fleet for FleetWise"
  signal_catalog_arn = awscc_iotfleetwise_signal_catalog.example.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create S3 bucket for data collection
resource "awscc_s3_bucket" "fleetwise_data" {
  bucket_name = "example-fleetwise-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Allow FleetWise access to S3
resource "awscc_s3_bucket_policy" "fleetwise_access" {
  bucket = awscc_s3_bucket.fleetwise_data.id
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowFleetWiseAccess"
        Effect = "Allow"
        Principal = {
          Service = "iotfleetwise.amazonaws.com"
        }
        Action = [
          "s3:PutObject",
          "s3:ListBucket"
        ]
        Resource = [
          awscc_s3_bucket.fleetwise_data.arn,
          "${awscc_s3_bucket.fleetwise_data.arn}/*"
        ]
      }
    ]
  })
}

# Create the campaign
resource "awscc_iotfleetwise_campaign" "example" {
  name               = "ExampleCampaign"
  signal_catalog_arn = awscc_iotfleetwise_signal_catalog.example.arn
  target_arn         = awscc_iotfleetwise_fleet.example.arn

  collection_scheme = {
    time_based_collection_scheme = {
      period_ms = 60000 # Collect data every minute
    }
  }

  signals_to_collect = [
    {
      name                         = "Vehicle.EngineSensor"
      max_sample_count             = 100
      minimum_sampling_interval_ms = 1000
    }
  ]

  data_destination_configs = [
    {
      s3_config = {
        bucket_arn  = awscc_s3_bucket.fleetwise_data.arn
        prefix      = "data/"
        data_format = "JSON"
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}