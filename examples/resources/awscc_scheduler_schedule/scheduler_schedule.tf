# Create an IAM role for the scheduler
resource "aws_iam_role" "scheduler_role" {
  name = "example-scheduler-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "scheduler.amazonaws.com"
        },
        Effect = "Allow",
        Sid    = ""
      }
    ]
  })
}

# Create an IAM role for the Lambda function
resource "aws_iam_role" "lambda_role" {
  name = "example-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Effect = "Allow",
        Sid    = ""
      }
    ]
  })
}

# Basic CloudWatch Logs policy for Lambda
resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Create a Lambda function to target
resource "aws_lambda_function" "example" {
  filename      = "function.zip"
  function_name = "example-function"
  role          = aws_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"
}

# Attach Lambda execution permission to the scheduler role
resource "aws_iam_role_policy" "lambda_execution" {
  name = "lambda-execution-policy"
  role = aws_iam_role.scheduler_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "lambda:InvokeFunction"
        ],
        Resource = aws_lambda_function.example.arn,
        Effect   = "Allow"
      }
    ]
  })
}

# EventBridge Scheduler Schedule using the AWSCC provider
resource "awscc_scheduler_schedule" "example_schedule" {
  name                         = "example-schedule"
  description                  = "Example schedule created with Terraform"
  schedule_expression          = "cron(0 12 * * ? *)" # Daily at 12:00 PM
  schedule_expression_timezone = "America/Los_Angeles"

  flexible_time_window = {
    mode                      = "FLEXIBLE"
    maximum_window_in_minutes = 30 # Allow 30-minute execution window
  }

  target = {
    arn      = aws_lambda_function.example.arn
    role_arn = aws_iam_role.scheduler_role.arn
    input = jsonencode({
      message = "Scheduled execution from EventBridge Scheduler"
    })
    retry_policy = {
      maximum_retry_attempts       = 3
      maximum_event_age_in_seconds = 3600 # 1 hour
    }
  }

  state = "ENABLED" # Schedule is active
}
