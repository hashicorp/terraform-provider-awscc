# Create the DMS service-linked role
resource "awscc_iam_service_linked_role" "dms" {
  aws_service_name = "dms.amazonaws.com"
  description      = "Service-linked role for AWS DMS"
}

# Required IAM role for DMS VPC access
data "aws_iam_policy_document" "dms_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["dms.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "dms_vpc_policy" {
  statement {
    actions = [
      "ec2:CreateNetworkInterface",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DeleteNetworkInterface",
      "ec2:DescribeSubnets",
      "ec2:DescribeVpcs",
      "ec2:DescribeSecurityGroups",
      "ec2:ModifyNetworkInterfaceAttribute",
      "ec2:DescribeAvailabilityZones"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "dms_vpc_role" {
  role_name                   = "dms-vpc-role"
  assume_role_policy_document = data.aws_iam_policy_document.dms_assume_role.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "dms_vpc_policy" {
  policy_name     = "dms-vpc-policy"
  role_name       = awscc_iam_role.dms_vpc_role.role_name
  policy_document = data.aws_iam_policy_document.dms_vpc_policy.json
}

# Create an DMS instance profile
resource "awscc_dms_instance_profile" "example" {
  instance_profile_identifier = "dms-instance-profile-example"
  instance_profile_name       = "DMS-Instance-Profile-Example"
  description                 = "Example DMS instance profile created with AWSCC"
  network_type                = "IPV4"
  publicly_accessible         = false

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  depends_on = [awscc_iam_role_policy.dms_vpc_policy, awscc_iam_service_linked_role.dms]
}