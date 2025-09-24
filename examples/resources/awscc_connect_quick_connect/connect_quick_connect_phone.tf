resource "awscc_connect_quick_connect" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "example"
  description  = "example for phone connect type"
  quick_connect_config = {
    quick_connect_type = "PHONE_NUMBER"
    phone_config = {
      phone_number = "+12345678912"
    }
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
