resource "awscc_connect_phone_number" "example" {
  target_arn   = aws_connect_instance.example.arn
  description  = "example"
  country_code = "US"
  type         = "DID"
}