data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for IoT Wireless Destination
resource "awscc_iam_role" "iot_destination_role" {
  role_name = "IoTWirelessDestinationRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iotwireless.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM policy for IoT Wireless Destination
data "aws_iam_policy_document" "iot_destination_policy" {
  statement {
    effect = "Allow"
    actions = [
      "iot:Connect",
      "iot:Publish",
      "iot:Subscribe",
      "iot:Receive"
    ]
    resources = [
      "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:topic/*",
      "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:topic-filter/*"
    ]
  }
}

resource "awscc_iam_role_policy" "iot_destination_policy" {
  policy_name = "IoTWirelessDestinationPolicy"
  role_name   = awscc_iam_role.iot_destination_role.role_name
  policy_document = jsonencode(
    jsondecode(data.aws_iam_policy_document.iot_destination_policy.json)
  )
}

# IoT Wireless Destination
resource "awscc_iotwireless_destination" "example" {
  name            = "example-destination"
  description     = "Example IoT Wireless Destination"
  expression_type = "RuleName"
  expression      = "IoTWirelessRule"
  role_arn        = awscc_iam_role.iot_destination_role.arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}