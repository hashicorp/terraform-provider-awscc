resource "awscc_elasticbeanstalk_application" "example-app" {
  application_name = "example-app"
  description      = "example-app"
}

resource "awscc_elasticbeanstalk_environment" "example-env" {
  application_name    = awscc_elasticbeanstalk_application.example-app.application_name
  solution_stack_name = "64bit Amazon Linux 2023 v4.0.3 running Python 3.11"
  option_settings = [{
    namespace   = "aws:autoscaling:launchconfiguration"
    option_name = "IamInstanceProfile"
    value       = "example-aws-elasticbeanstalk-ec2-role"
  }]
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
