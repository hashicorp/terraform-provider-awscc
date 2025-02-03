data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# S3 bucket to store the layer code
resource "awscc_s3_bucket" "lambda_layer" {
  bucket_name = "lambda-layer-example-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  # Add required tags per guidelines
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload a sample Python package to S3
resource "aws_s3_object" "layer_code" {
  bucket = awscc_s3_bucket.lambda_layer.id
  key    = "python-layer.zip"
  source = "${path.module}/python-layer.zip"

  # Using etag to force update when the file changes
  etag = filemd5("${path.module}/python-layer.zip")
}

# Create Lambda layer version
resource "awscc_lambda_layer_version" "example" {
  layer_name  = "example-layer"
  description = "Example Lambda layer created with AWSCC provider"

  content = {
    s3_bucket = awscc_s3_bucket.lambda_layer.id
    s3_key    = aws_s3_object.layer_code.key
  }

  compatible_runtimes      = ["python3.8", "python3.9"]
  compatible_architectures = ["x86_64"]
  license_info             = "MIT"
}