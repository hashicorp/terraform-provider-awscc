data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a basic S3 bucket
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-${data.aws_caller_identity.current.account_id}-${formatdate("YYYYMMDD", timestamp())}"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}

# Create an S3 access point for the bucket
resource "awscc_s3_access_point" "example" {
  bucket = awscc_s3_bucket.example.id
  name   = "example-ap"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  vpc_configuration = {
    vpc_id = null
  }
}

# Create Lambda function for object transformation
resource "aws_lambda_function" "example" {
  filename      = "lambda_function.zip"
  function_name = "s3-object-lambda-function"
  role          = awscc_iam_role.lambda.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"

  source_code_hash = filebase64sha256("lambda_function.zip")
}

# Create IAM role for Lambda
resource "awscc_iam_role" "lambda" {
  role_name = "s3-object-lambda-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM policy for Lambda
data "aws_iam_policy_document" "lambda_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "s3-object-lambda:WriteGetObjectResponse"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role_policy" "lambda" {
  policy_name     = "s3-object-lambda-policy"
  role_name       = awscc_iam_role.lambda.role_name
  policy_document = data.aws_iam_policy_document.lambda_policy.json
}

# Create S3 Object Lambda Access Point
resource "awscc_s3objectlambda_access_point" "example" {
  name = "example-s3-object-lambda-ap"
  object_lambda_configuration = {
    supporting_access_point    = awscc_s3_access_point.example.arn
    cloudwatch_metrics_enabled = true
    allowed_features           = ["GetObject-Range", "GetObject-PartNumber"]
    transformation_configurations = [{
      actions = ["GetObject"]
      content_transformation = {
        aws_lambda = {
          function_arn = aws_lambda_function.example.arn
          function_payload = jsonencode({
            type = "example"
          })
        }
      }
    }]
  }
}