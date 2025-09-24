# Data source for AWS Account ID
data "aws_caller_identity" "current" {}

# Data source for AWS Region
data "aws_region" "current" {}

# Data source for access point policy
data "aws_iam_policy_document" "access_point_policy" {
  statement {
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "s3outposts:*"
    ]
    resources = [
      "arn:aws:s3-outposts:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:outpost/op-1234567890123/accesspoint/${var.access_point_name}",
      "arn:aws:s3-outposts:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:outpost/op-1234567890123/accesspoint/${var.access_point_name}/*"
    ]
  }
}

variable "access_point_name" {
  type        = string
  description = "Name of the access point"
  default     = "example-ap"
}

# S3 Outposts Access Point
resource "awscc_s3outposts_access_point" "example" {
  name   = var.access_point_name
  bucket = "arn:aws:s3-outposts:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:outpost/op-1234567890123/bucket/example-bucket"
  policy = jsonencode(data.aws_iam_policy_document.access_point_policy.json)

  vpc_configuration = {
    vpc_id = "vpc-1234567890abcdef0"
  }
}