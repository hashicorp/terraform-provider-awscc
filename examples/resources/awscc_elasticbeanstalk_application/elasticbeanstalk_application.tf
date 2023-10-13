resource "awscc_elasticbeanstalk_application" "example" {
  application_name = "SampleAWSElasticBeanstalkApplication"
  description      = "AWS Elastic Beanstalk PHP Sample Application"
  resource_lifecycle_config = {
    service_role = awscc_iam_role.elasticbeanstalk_servicerole.arn
    version_lifecycle_config = {
      max_count_rule = {
        enabled               = true
        delete_source_from_s3 = false
        max_count             = 50
      }
    }
  }
}

resource "awscc_iam_role" "elasticbeanstalk_servicerole" {
  role_name           = "elasticbeanstalk-custom-service-role"
  description         = "This is a service role for ElasticBeanstalk"
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSElasticBeanstalkEnhancedHealth", "arn:aws:iam::aws:policy/AWSElasticBeanstalkManagedUpdatesCustomerRolePolicy"]
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "elasticbeanstalk.amazonaws.com"
        }
      }
    ]

  })
}