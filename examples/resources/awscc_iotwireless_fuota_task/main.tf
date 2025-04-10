# Create IAM role for FUOTA task
data "aws_iam_policy_document" "fuota_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iotwireless.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "fuota_policy" {
  statement {
    actions = [
      "s3:GetObject"
    ]
    resources = [
      "arn:aws:s3:::example-bucket/firmware/*"
    ]
  }
}

resource "awscc_iam_role" "fuota_role" {
  role_name                   = "IoTWirelessFUOTARole"
  assume_role_policy_document = data.aws_iam_policy_document.fuota_assume_role.json
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "fuota_role_policy" {
  policy_name     = "IoTWirelessFUOTAPolicy"
  role_name       = awscc_iam_role.fuota_role.role_name
  policy_document = data.aws_iam_policy_document.fuota_policy.json
}

# Create FUOTA task
resource "awscc_iotwireless_fuota_task" "example" {
  name                  = "example-fuota-task"
  description           = "Example FUOTA task for firmware updates"
  firmware_update_image = "s3://example-bucket/firmware/update.bin"
  firmware_update_role  = awscc_iam_role.fuota_role.arn

  lo_ra_wan = {
    rf_region = "US915"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}