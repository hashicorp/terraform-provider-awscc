#Example with lambda function
resource "aws_greengrassv2_component_version" "MyGreengrassComponentVersion" {
  lambda_function {
    lambda_arn        = "arn:aws:lambda:<region>:<account-id>:function:<LambdaFunctionName>:<version>"
    component_name    = "MyLambdaComponent"
    component_version = "1.0.0"
  }

  tags = {
    Environment = "Production"
    Project     = "Greengrass-awscc-Project"
  }
}
