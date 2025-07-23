# Create the IAM role for MediaPackage SPEKE
data "aws_iam_policy_document" "mediapackage_speke_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["mediapackage.amazonaws.com"]
    }
  }
}

# Create the SPEKE role
resource "awscc_iam_role" "mediapackage_speke_role" {
  role_name                   = "MediaPackageSpekeRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.mediapackage_speke_assume_role.json))
}

# First create a packaging group
resource "awscc_mediapackage_packaging_group" "example" {
  packaging_group_id = "example-packaging-group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Then create the packaging configuration
resource "awscc_mediapackage_packaging_configuration" "example" {
  packaging_configuration_id = "example-packaging-config"
  packaging_group_id         = awscc_mediapackage_packaging_group.example.id

  # Example using DASH package configuration
  dash_package = {
    segment_duration_seconds   = 30
    include_iframe_only_stream = true

    # Add manifest configuration
    dash_manifests = [{
      manifest_name           = "index"
      min_buffer_time_seconds = 30
      profile                 = "NONE"
      stream_selection = {
        min_video_bits_per_second = 1000000
        max_video_bits_per_second = 8000000
        stream_order              = "ORIGINAL"
      }
    }]

    # Add encryption configuration
    encryption = {
      speke_key_provider = {
        role_arn   = awscc_iam_role.mediapackage_speke_role.arn
        system_ids = ["edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"]
        url        = "https://example.com/speke"
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}