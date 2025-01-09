data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create S3 bucket for MWAA
resource "aws_s3_bucket" "mwaa" {
  bucket = "mwaa-environment-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

resource "aws_s3_bucket_public_access_block" "mwaa" {
  bucket = aws_s3_bucket.mwaa.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_versioning" "mwaa" {
  bucket = aws_s3_bucket.mwaa.id
  versioning_configuration {
    status = "Enabled"
  }
}

# Create VPC and networking components
resource "aws_vpc" "mwaa" {
  cidr_block           = "10.192.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "mwaa-environment"
  }
}

resource "aws_subnet" "private_1" {
  vpc_id            = aws_vpc.mwaa.id
  cidr_block        = "10.192.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"

  tags = {
    Name = "mwaa-private-1"
  }
}

resource "aws_subnet" "private_2" {
  vpc_id            = aws_vpc.mwaa.id
  cidr_block        = "10.192.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"

  tags = {
    Name = "mwaa-private-2"
  }
}

resource "aws_security_group" "mwaa" {
  name_prefix = "mwaa-environment-"
  vpc_id      = aws_vpc.mwaa.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# IAM role and policies
data "aws_iam_policy_document" "mwaa_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["airflow.amazonaws.com", "airflow-env.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "mwaa_execution" {
  statement {
    effect = "Allow"
    actions = [
      "airflow:PublishMetrics",
      "s3:ListAllMyBuckets",
      "s3:GetBucketLocation"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject*",
      "s3:GetBucket*",
      "s3:List*"
    ]
    resources = [
      aws_s3_bucket.mwaa.arn,
      "${aws_s3_bucket.mwaa.arn}/*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:CreateLogGroup",
      "logs:PutLogEvents",
      "logs:GetLogEvents",
      "logs:GetLogRecord",
      "logs:GetLogGroupFields",
      "logs:GetQueryResults"
    ]
    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:airflow-*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "cloudwatch:PutMetricData"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "sqs:ChangeMessageVisibility",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes",
      "sqs:GetQueueUrl",
      "sqs:ReceiveMessage",
      "sqs:SendMessage"
    ]
    resources = [
      "arn:aws:sqs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:airflow-celery-*"
    ]
  }
}

resource "awscc_iam_role" "mwaa_execution" {
  role_name                   = "mwaa-environment-execution-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.mwaa_assume_role.json))
  managed_policy_arns         = []

  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.mwaa_execution.json))
    policy_name     = "MWAA-Execution-Policy"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# MWAA Environment
resource "awscc_mwaa_environment" "example" {
  name = "mwaa-environment-example"

  airflow_version       = "2.7.2"
  environment_class     = "mw1.small"
  dag_s3_path           = "dags"
  max_workers           = 2
  min_workers           = 1
  webserver_access_mode = "PUBLIC_ONLY"

  execution_role_arn = awscc_iam_role.mwaa_execution.arn
  source_bucket_arn  = aws_s3_bucket.mwaa.arn

  logging_configuration = {
    dag_processing_logs = {
      enabled   = true
      log_level = "INFO"
    }
    scheduler_logs = {
      enabled   = true
      log_level = "INFO"
    }
    task_logs = {
      enabled   = true
      log_level = "INFO"
    }
    webserver_logs = {
      enabled   = true
      log_level = "INFO"
    }
    worker_logs = {
      enabled   = true
      log_level = "INFO"
    }
  }

  network_configuration = {
    security_group_ids = [aws_security_group.mwaa.id]
    subnet_ids         = [aws_subnet.private_1.id, aws_subnet.private_2.id]
  }

  tags = jsonencode([{
    key   = "Modified By"
    value = "AWSCC"
  }])
}