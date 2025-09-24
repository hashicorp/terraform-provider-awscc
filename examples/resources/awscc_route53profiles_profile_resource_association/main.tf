# This is just a demonstration configuration
# In practice, you need to have a valid AWS Route53 Profile and resource ARN
resource "awscc_route53profiles_profile_resource_association" "example" {
  name       = "example-association"
  profile_id = "12345678-1234-1234-1234-123456789012" # Replace with actual profile ID
  # Replace with actual Route53 domain ARN
  resource_arn = "arn:aws:route53:us-east-1:123456789012:hostedzone/Z123456789ABC"
  # Optional resource properties as JSON string
  resource_properties = jsonencode({
    "Property1" = "Value1"
    "Property2" = "Value2"
  })
}