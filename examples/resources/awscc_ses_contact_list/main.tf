# Create a simple contact list with tags and topics
resource "awscc_ses_contact_list" "example" {
  contact_list_name = "my-contact-list"
  description       = "Example contact list created with AWSCC provider"

  topics = [
    {
      topic_name                  = "newsletters"
      display_name                = "Monthly Newsletters"
      description                 = "Monthly updates and news"
      default_subscription_status = "OPT_IN"
    },
    {
      topic_name                  = "promotions"
      display_name                = "Special Promotions"
      description                 = "Special offers and promotions"
      default_subscription_status = "OPT_OUT"
    }
  ]

  tags = [
    {
      key   = "Environment"
      value = "Test"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}