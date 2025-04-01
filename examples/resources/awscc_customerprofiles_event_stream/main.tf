# Create a Kinesis stream as the destination
resource "aws_kinesis_stream" "profile_events" {
  name             = "customer-profile-events"
  shard_count      = 1
  retention_period = 24

  tags = {
    Environment = "test"
  }
}

# Create an IAM role and policy for Customer Profiles to access Kinesis
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["profile.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "profile_kinesis_role" {
  name               = "CustomerProfilesKinesisAccess"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
  description        = "IAM role for Customer Profiles to access Kinesis stream"

  tags = {
    "Modified By" = "AWSCC"
  }
}

data "aws_iam_policy_document" "kinesis_access" {
  statement {
    effect = "Allow"
    actions = [
      "kinesis:PutRecord",
      "kinesis:PutRecords",
      "kinesis:DescribeStream"
    ]
    resources = [aws_kinesis_stream.profile_events.arn]
  }
}

resource "aws_iam_policy" "profile_kinesis_policy" {
  name        = "CustomerProfilesKinesisAccess"
  description = "Policy for Customer Profiles to access Kinesis stream"
  policy      = data.aws_iam_policy_document.kinesis_access.json

  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_iam_role_policy_attachment" "profile_kinesis" {
  policy_arn = aws_iam_policy.profile_kinesis_policy.arn
  role       = aws_iam_role.profile_kinesis_role.name
}

# Create the CustomerProfiles Domain (required for event stream)
resource "awscc_customerprofiles_domain" "example" {
  domain_name             = "example-domain"
  default_expiration_days = 366

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create the Event Stream
resource "awscc_customerprofiles_event_stream" "example" {
  domain_name       = awscc_customerprofiles_domain.example.domain_name
  event_stream_name = "example-event-stream"
  uri               = aws_kinesis_stream.profile_events.arn

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}