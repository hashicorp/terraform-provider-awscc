resource "awscc_iam_service_linked_role" "elasticbeanstalk" {
  aws_service_name = "elasticbeanstalk.amazonaws.com"
  description      = "service linked role for AWS Elastic Beanstalk"
}