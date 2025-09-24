resource "awscc_ssm_parameter" "example" {
  name  = "command"
  type  = "String"
  value = "date"
  tier  = "Advanced"
  policies = jsonencode([
    {
      "Type" : "Expiration",
      "Version" : "1.0",
      "Attributes" : {
        "Timestamp" : "2024-05-13T00:00:00.000Z"
      }
    },
    {
      "Type" : "ExpirationNotification",
      "Version" : "1.0",
      "Attributes" : {
        "Before" : "5",
        "Unit" : "Days"
      }
    },
    {
      "Type" : "NoChangeNotification",
      "Version" : "1.0",
      "Attributes" : {
        "After" : "60",
        "Unit" : "Days"
      }
    }
  ])
  description     = "SSM Parameter for running date command."
  allowed_pattern = "^[a-zA-Z]{1,10}$"
}