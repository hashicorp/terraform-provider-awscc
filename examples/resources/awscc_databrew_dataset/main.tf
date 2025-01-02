data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for the DataBrew dataset
resource "awscc_s3_bucket" "databrew_input" {
  bucket_name = "databrew-input-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload a sample CSV file to the bucket
resource "aws_s3_object" "sample_data" {
  bucket  = awscc_s3_bucket.databrew_input.id
  key     = "input/sample.csv"
  content = "id,name,age\n1,John,30\n2,Jane,25"
}

# Create the DataBrew dataset
resource "awscc_databrew_dataset" "example" {
  name = "example-dataset"

  input = {
    s3_input_definition = {
      bucket = awscc_s3_bucket.databrew_input.id
      key    = aws_s3_object.sample_data.key
    }
  }

  format = "CSV"
  format_options = {
    csv = {
      delimiter  = ","
      header_row = true
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}