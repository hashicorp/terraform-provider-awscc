# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Policy document for IoT Wireless Partner Account
data "aws_iam_policy_document" "sidewalk" {
  statement {
    effect = "Allow"
    actions = [
      "iotwireless:CreateWirelessDevice",
      "iotwireless:UpdateWirelessDevice",
      "iotwireless:DeleteWirelessDevice"
    ]
    resources = [
      "arn:aws:iotwireless:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:WirelessDevice/*"
    ]
  }
}

resource "awscc_iotwireless_partner_account" "example" {
  sidewalk = {
    app_server_private_key = "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
    policy_document        = jsonencode(jsondecode(data.aws_iam_policy_document.sidewalk.json))
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}