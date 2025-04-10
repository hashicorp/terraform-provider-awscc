# Create an IAM role that can be associated with the ACM certificate
data "aws_iam_policy_document" "ec2_nitro_enclaves_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "acm_access" {
  statement {
    actions = [
      "acm:GetCertificateRequest",
      "acm:ExportCertificate",
    ]

    resources = ["*"]

    effect = "Allow"
  }
}

resource "awscc_iam_role" "nitro_enclave_role" {
  role_name = "nitro-enclave-role"
  assume_role_policy_document = jsonencode(
    jsondecode(data.aws_iam_policy_document.ec2_nitro_enclaves_assume_role.json)
  )

  tags = [{
    key   = "Created By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "nitro_enclave_policy" {
  role_name   = awscc_iam_role.nitro_enclave_role.role_name
  policy_name = "nitro-enclave-acm-access"
  policy_document = jsonencode(
    jsondecode(data.aws_iam_policy_document.acm_access.json)
  )
}

# Create a self-signed certificate for testing
resource "aws_acm_certificate" "cert" {
  domain_name       = "example.com"
  validation_method = "EMAIL"

  tags = {
    "Created By" = "AWSCC"
  }

  lifecycle {
    create_before_destroy = true
  }
}

# Associate the IAM role with the ACM certificate
resource "awscc_ec2_enclave_certificate_iam_role_association" "example" {
  certificate_arn = aws_acm_certificate.cert.arn
  role_arn        = awscc_iam_role.nitro_enclave_role.arn
}