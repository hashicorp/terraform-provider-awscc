resource "awscc_datasync_location_s3" "example" {
  s3_bucket_arn = "arn:aws:s3:::example-bucket"
  s3_config = {
    bucket_access_role_arn = awscc_iam_role.example.arn
  }
  s3_storage_class = "STANDARD"
  subdirectory     = "/docs"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "example" {
  role_name = "AWSDataSyncS3BucketAccess-example-bucket"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "datasync.amazonaws.com"
        }
        Action = "sts:AssumeRole"
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          },
          ArnLike = {
            "aws:SourceArn" = "arn:aws:datasync:us-east-1:${data.aws_caller_identity.current.account_id}:*"
          }
        }
      }
    ]
  })
  policies = [{
    policy_name = "s2_ds_inline"
    policy_document = jsonencode(
      {
        "Version" : "2012-10-17",
        "Statement" : [
          {
            "Sid" : "AWSDataSyncS3BucketPermissions",
            "Effect" : "Allow",
            "Action" : [
              "s3:GetBucketLocation",
              "s3:ListBucket",
              "s3:ListBucketMultipartUploads"
            ],
            "Resource" : "arn:aws:s3:::example-bucket",
            "Condition" : {
              "StringEquals" : {
                "aws:ResourceAccount" : data.aws_caller_identity.current.account_id
              }
            }
          },
          {
            "Sid" : "AWSDataSyncS3ObjectPermissions",
            "Effect" : "Allow",
            "Action" : [
              "s3:AbortMultipartUpload",
              "s3:DeleteObject",
              "s3:GetObject",
              "s3:GetObjectTagging",
              "s3:GetObjectVersion",
              "s3:GetObjectVersionTagging",
              "s3:ListMultipartUploadParts",
              "s3:PutObject",
              "s3:PutObjectTagging"
            ],
            "Resource" : "arn:aws:s3:::example-bucket/*",
            "Condition" : {
              "StringEquals" : {
                "aws:ResourceAccount" : data.aws_caller_identity.current.account_id
              }
            }
          }
        ]
      }
    )
  }]

}

data "aws_caller_identity" "current" {}
