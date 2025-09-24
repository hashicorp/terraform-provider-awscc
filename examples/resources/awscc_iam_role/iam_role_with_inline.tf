data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}


data "aws_iam_policy_document" "sample_inline_1" {
  statement {
    sid       = "AccessS3"
    actions   = ["s3:ListAllMyBuckets", "s3:ListBucket", "s3:HeadBucket"]
    resources = ["*"]
  }
}


data "aws_iam_policy_document" "sample_inline_2" {
  statement {
    sid       = "AccessEC2"
    actions   = ["ec2:Describe*"]
    resources = ["*"]
  }
}



resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  path                        = "/"
  policies = [{
    policy_document = data.aws_iam_policy_document.sample_inline_1.json
    policy_name     = "fist_inline_policy"
    },
    {
      policy_document = data.aws_iam_policy_document.sample_inline_2.json
      policy_name     = "second_inline_policy"
  }]
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}