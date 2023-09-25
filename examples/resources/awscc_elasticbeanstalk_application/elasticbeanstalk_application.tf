resource "awscc_elasticbeanstalk_application" "example_elasticbeanstalk_application" {
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
