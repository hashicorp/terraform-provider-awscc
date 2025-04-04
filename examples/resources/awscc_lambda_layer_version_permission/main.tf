resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "aws_s3_bucket" "layer" {
  bucket = "lambda-layer-example-${random_string.suffix.result}"
}

resource "local_file" "nodejs_file" {
  filename = "${path.module}/nodejs/example.js"
  content  = "exports.handler = async () => { return 'Hello from Layer!'; };"
}

data "archive_file" "layer" {
  type        = "zip"
  output_path = "${path.module}/layer.zip"
  source_dir  = "${path.module}/nodejs"
  depends_on  = [local_file.nodejs_file]
}

resource "aws_s3_object" "layer" {
  bucket = aws_s3_bucket.layer.id
  key    = "layer.zip"
  source = data.archive_file.layer.output_path
}

resource "aws_lambda_layer_version" "example" {
  layer_name          = "example-layer"
  description         = "Example Layer created for layer version permission test"
  compatible_runtimes = ["nodejs18.x"]
  filename            = data.archive_file.layer.output_path
  compatible_architectures = ["x86_64"]
}

resource "awscc_lambda_layer_version_permission" "example" {
  action            = "lambda:GetLayerVersion"
  layer_version_arn = aws_lambda_layer_version.example.arn
  principal         = "*"
  organization_id = "o-xxxxxxxxxx"
}