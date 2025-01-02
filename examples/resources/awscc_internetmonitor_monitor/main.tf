# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for the monitor logs
resource "awscc_s3_bucket" "monitor_logs" {
  bucket_name = "internetmonitor-logs-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a bucket policy to allow Internet Monitor to write logs
# Note: We still need aws_iam_policy_document data source as it's a helper for policy generation
data "aws_iam_policy_document" "monitor_bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["internetmonitor.amazonaws.com"]
    }
    actions = [
      "s3:PutObject"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.monitor_logs.bucket_name}/*"
    ]
  }
}

resource "awscc_s3_bucket_policy" "monitor_logs" {
  bucket          = awscc_s3_bucket.monitor_logs.bucket_name
  policy_document = data.aws_iam_policy_document.monitor_bucket_policy.json
}

# Create the Internet Monitor monitor
resource "awscc_internetmonitor_monitor" "example" {
  monitor_name = "example-monitor"

  # Monitor 100 city networks
  max_city_networks_to_monitor = 100

  # Monitor 20% of traffic
  traffic_percentage_to_monitor = 20

  # Define health events configuration
  health_events_config = {
    availability_score_threshold = 95
    performance_score_threshold  = 95

    availability_local_health_events_config = {
      health_score_threshold = 90
      min_traffic_impact     = 0.2
      status                 = "ENABLED"
    }

    performance_local_health_events_config = {
      health_score_threshold = 90
      min_traffic_impact     = 0.2
      status                 = "ENABLED"
    }
  }

  # Configure log delivery to S3
  internet_measurements_log_delivery = {
    s3_config = {
      bucket_name         = awscc_s3_bucket.monitor_logs.bucket_name
      bucket_prefix       = "internet-monitor-logs"
      log_delivery_status = "ENABLED"
    }
  }

  # Add example resources to monitor
  # Only setting monitor name and log delivery for example
  resources = []

  # Add tags
  tags = [{
    key   = "Environment"
    value = "Production"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}