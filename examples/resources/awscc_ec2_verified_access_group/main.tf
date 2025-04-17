# Create Verified Access Instance first (required for the group)
resource "awscc_ec2_verified_access_instance" "example" {
  description = "Example Verified Access Instance"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Trust Provider
resource "aws_verifiedaccess_trust_provider" "example" {
  policy_reference_name = "ExampleTrustProvider"
  trust_provider_type   = "user"

  tags = {
    Modified_By = "AWSCC"
  }

  user_trust_provider_type = "iam-identity-center"
}

# Associate the trust provider with the instance
resource "aws_verifiedaccess_instance_trust_provider_attachment" "example" {
  verifiedaccess_instance_id       = awscc_ec2_verified_access_instance.example.verified_access_instance_id
  verifiedaccess_trust_provider_id = aws_verifiedaccess_trust_provider.example.id
}

# Create policy document
data "aws_iam_policy_document" "verified_access_policy" {
  statement {
    effect = "Allow"
    condition {
      test     = "StringEquals"
      variable = "aws:PrincipalTag/Department"
      values   = ["Engineering"]
    }
    principals {
      type        = "*"
      identifiers = ["*"]
    }
  }
}

# Create Verified Access Group
resource "awscc_ec2_verified_access_group" "example" {
  verified_access_instance_id = awscc_ec2_verified_access_instance.example.verified_access_instance_id
  description                 = "Example Verified Access Group with Policy"
  policy_enabled              = true
  policy_document             = jsonencode(jsondecode(data.aws_iam_policy_document.verified_access_policy.json))

  depends_on = [aws_verifiedaccess_instance_trust_provider_attachment.example]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}