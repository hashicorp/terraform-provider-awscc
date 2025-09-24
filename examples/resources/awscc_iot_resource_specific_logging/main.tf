# Create IAM role for IoT logging
resource "awscc_iam_role" "iot_logging" {
  role_name = "iot-logging-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iot.amazonaws.com"
        }
      }
    ]
  })
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSIoTLogging"
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# First create an IoT thing group as target
resource "awscc_iot_thing_group" "example" {
  thing_group_name = "example-thing-group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Set up IoT logging options
resource "aws_iot_logging_options" "example" {
  default_log_level = "ERROR"
  role_arn          = awscc_iam_role.iot_logging.arn
}

# Create resource specific logging for the thing group
resource "awscc_iot_resource_specific_logging" "example" {
  target_type = "THING_GROUP"
  target_name = awscc_iot_thing_group.example.thing_group_name
  log_level   = "INFO"
  depends_on  = [aws_iot_logging_options.example]
}