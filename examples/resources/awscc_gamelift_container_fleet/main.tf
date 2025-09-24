# Create IAM role for GameLift fleet
data "aws_iam_policy_document" "fleet_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["gamelift.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "fleet_role" {
  role_name                   = "gamelift-fleet-role"
  assume_role_policy_document = data.aws_iam_policy_document.fleet_assume_role.json
  description                 = "IAM role for GameLift Fleet"
  policies = [{
    policy_name = "fleet-policy"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow"
          Action = [
            "s3:GetObject",
            "s3:ListBucket",
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:PutLogEvents",
            "logs:DescribeLogGroups",
            "logs:DescribeLogStreams"
          ]
          Resource = "*"
        }
      ]
    })
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create GameLift Container Fleet
resource "awscc_gamelift_container_fleet" "example" {
  fleet_role_arn = awscc_iam_role.fleet_role.arn
  description    = "Example GameLift Container Fleet"

  instance_type = "c5.large"
  billing_type  = "ON_DEMAND"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}