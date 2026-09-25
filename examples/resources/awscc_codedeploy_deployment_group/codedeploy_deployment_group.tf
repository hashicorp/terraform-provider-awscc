resource "aws_codedeploy_app" "example" {
  name             = "example-application"
  compute_platform = "Server"

  tags = {
    Name        = "example-application"
    Environment = "example"
  }
}

resource "aws_iam_role" "codedeploy_service" {
  name = "example-codedeploy-service-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "codedeploy.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name        = "example-codedeploy-service-role"
    Environment = "example"
  }
}

resource "aws_iam_role_policy_attachment" "codedeploy_service" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"
  role       = aws_iam_role.codedeploy_service.name
}

resource "time_sleep" "wait_for_iam_role" {
  depends_on      = [aws_iam_role_policy_attachment.codedeploy_service]
  create_duration = "15s"
}

resource "awscc_codedeploy_deployment_group" "example" {
  depends_on = [time_sleep.wait_for_iam_role]

  application_name      = aws_codedeploy_app.example.name
  service_role_arn      = aws_iam_role.codedeploy_service.arn
  deployment_group_name = "example-deployment-group"

  deployment_config_name = "CodeDeployDefault.OneAtATime"

  ec_2_tag_filters = [
    {
      type  = "KEY_AND_VALUE"
      key   = "Environment"
      value = "example"
    }
  ]

  tags = [
    {
      key   = "Name"
      value = "example-deployment-group"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
