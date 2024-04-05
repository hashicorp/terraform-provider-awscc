#Example with lambda function
resource "aws_greengrassv2_component_version" "MyGreengrassComponentVersion" {
  lambda_function {
    component_lambda_version = "arn:aws:lambda:<region>:<account-id>:function:<LambdaFunctionName>:<version>"
    component_name           = "MyLambdaComponent"
    component_publisher      = "MyCompany"
  }

  tags = {
    Environment = "Production"
    Project     = "Greengrass-awscc-Project"
  }
}
