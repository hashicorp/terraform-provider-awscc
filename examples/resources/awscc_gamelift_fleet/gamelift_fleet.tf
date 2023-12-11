resource "awscc_gamelift_fleet" "example" {
  name              = "ExampleFleet"
  build_id          = awscc_gamelift_build.example.id
  compute_type      = "EC2"
  ec2_instance_type = "c5.large"
  fleet_type        = "ON_DEMAND"

  runtime_configuration = {
    server_processes = [
      {
        concurrent_executions = 1
        launch_path           = "/local/game/path-to-your-file"
        parameters            = "yourParameterKey:yourParameterValue yourParameterKey2:YourParameterValue2"
      },
      {
        concurrent_executions = 1
        launch_path           = "/local/game/path-to-your-file"
        parameters            = "yourParameterKey:yourParameterValue yourParameterKey2:YourParameterValue2"
      },
    ]
  }
}

resource "awscc_gamelift_build" "example" {
  name             = "ExampleBuild"
  version          = "1.0"
  operating_system = "AMAZON_LINUX_2"

  storage_location = {
    bucket   = "your-s3-bucket"
    key      = "your-s3-key"
    role_arn = awscc_iam_role.example.arn
  }
}

resource "awscc_iam_role" "example" {
  role_name                   = "gamelift-s3-access"
  description                 = "This IAM role grants Amazon GameLift access to the S3 bucket containing build files"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  managed_policy_arns         = [aws_iam_policy.example.arn]
  max_session_duration        = 7200
  path                        = "/"
  tags = [
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

data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["gamelift.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "example" {
  name = "gamelift-s3-access-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["s3:*"]
        Resource = "*"
      },
    ]
  })
}
