data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for Lambda
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lambda_permissions" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "iot:CreateCertificateFromCsr",
      "iot:DeleteCertificate",
      "iot:DescribeCertificate",
      "iot:RegisterCertificate",
      "iot:UpdateCertificate"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "lambda_role" {
  role_name                   = "iot-certificate-provider-role"
  assume_role_policy_document = data.aws_iam_policy_document.lambda_assume_role.json
  path                        = "/"

  policies = [
    {
      policy_name     = "IoTCertificateProviderPolicy"
      policy_document = data.aws_iam_policy_document.lambda_permissions.json
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Allow IoT service to invoke the Lambda function
resource "awscc_lambda_permission" "allow_iot_invoke" {
  action        = "lambda:InvokeFunction"
  function_name = "iot-certificate-provider"
  principal     = "iot.amazonaws.com"
  source_arn    = "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
}

# Sample Lambda function for certificate provider
resource "aws_lambda_function" "cert_provider" {
  filename      = "lambda_function.zip"
  function_name = "iot-certificate-provider"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a sample Lambda function code
resource "local_file" "lambda_function" {
  filename = "index.js"
  content  = <<-EOT
exports.handler = async (event) => {
    console.log('Received event:', JSON.stringify(event));
    return {
        statusCode: 200,
        body: JSON.stringify('Hello from Certificate Provider!')
    };
};
EOT
}

# Create zip file for Lambda
data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = local_file.lambda_function.filename
  output_path = "lambda_function.zip"
}

# IoT Certificate Provider
resource "awscc_iot_certificate_provider" "example" {
  certificate_provider_name      = "example-provider"
  lambda_function_arn            = aws_lambda_function.cert_provider.arn
  account_default_for_operations = ["CreateCertificateFromCsr"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}