# Create KMS key for IoT encryption
resource "awscc_kms_key" "iot_encryption" {
  description = "KMS key for IoT encryption configuration"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "*"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow IoT service"
        Effect = "Allow"
        Principal = {
          Service = "iot.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:DescribeKey",
          "kms:Encrypt",
          "kms:GenerateDataKey*",
          "kms:ReEncrypt*"
        ]
        Resource = "*"
      }
    ]
  })
}

# Create IAM role for IoT encryption
resource "awscc_iam_role" "iot_encryption" {
  role_name = "example-iot-encryption-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "iot.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
  policies = [
    {
      policy_name = "IoTEncryptionPolicy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "kms:Decrypt",
              "kms:DescribeKey",
              "kms:Encrypt",
              "kms:GenerateDataKey*",
              "kms:ReEncrypt*"
            ]
            Resource = "*"
          }
        ]
      })
    }
  ]
}

resource "awscc_iot_encryption_configuration" "example" {
  encryption_type     = "CUSTOMER_MANAGED_KMS_KEY"
  kms_key_arn         = awscc_kms_key.iot_encryption.arn
  kms_access_role_arn = awscc_iam_role.iot_encryption.arn
}