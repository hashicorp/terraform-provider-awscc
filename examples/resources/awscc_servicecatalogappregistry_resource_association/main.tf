# Create an App Registry Application first
resource "awscc_servicecatalogappregistry_application" "example" {
  name        = "my-example-app"
  description = "Example application for resource association"
}

# Create a CloudFormation Stack to associate
resource "aws_cloudformation_stack" "example" {
  name = "example-stack"

  template_body = jsonencode({
    Resources = {
      ExampleLogGroup = {
        Type = "AWS::Logs::LogGroup"
        Properties = {
          LogGroupName = "/aws/example/log-group"
        }
      }
    }
  })

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create the resource association
resource "awscc_servicecatalogappregistry_resource_association" "example" {
  application   = awscc_servicecatalogappregistry_application.example.id
  resource      = aws_cloudformation_stack.example.id
  resource_type = "CFN_STACK"
}