resource "awscc_smsvoice_phone_number" "example" {
  iso_country_code    = "US"
  number_type         = "TOLL_FREE"
  number_capabilities = ["SMS"]
  mandatory_keywords = {
    help = {
      keyword = "HELP"
      message = "Reply HELP for help"
    }
    stop = {
      keyword = "STOP"
      message = "Reply STOP to unsubscribe"
    }
  }

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-phone-number"
    }
  ]
}