# Get the current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an S3 bucket to store the app block files
resource "awscc_s3_bucket" "app_block" {
  bucket_name = "appstream-app-block-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the app block
resource "awscc_appstream_app_block" "example" {
  name        = "example-app-block"
  description = "Example AppStream App Block"

  source_s3_location = {
    s3_bucket = awscc_s3_bucket.app_block.id
    s3_key    = "app-block/source/myapp.zip"
  }

  setup_script_details = {
    script_s3_location = {
      s3_bucket = awscc_s3_bucket.app_block.id
      s3_key    = "app-block/scripts/setup.ps1"
    }
    executable_path    = ".\\setup.ps1"
    timeout_in_seconds = 300
  }

  post_setup_script_details = {
    script_s3_location = {
      s3_bucket = awscc_s3_bucket.app_block.id
      s3_key    = "app-block/scripts/post-setup.ps1"
    }
    executable_path    = ".\\post-setup.ps1"
    timeout_in_seconds = 180
  }

  packaging_type = "APPSTREAM2"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}