# Pinpoint In-App Template
resource "awscc_pinpoint_in_app_template" "example" {
  template_name        = "example-template"
  template_description = "Example in-app template for notifications"
  layout               = "MOBILE_FEED"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  content = [{
    background_color = "#FFFFFF"
    body_config = {
      alignment  = "CENTER"
      body       = "Welcome to our app! Check out our latest features."
      text_color = "#000000"
    }
    header_config = {
      alignment  = "CENTER"
      header     = "Welcome Message"
      text_color = "#000000"
    }
    image_url = "https://example.com/welcome-image.png"
    primary_btn = {
      default_config = {
        background_color = "#4CAF50"
        border_radius    = 8
        button_action    = "DEEP_LINK"
        link             = "myapp://features"
        text             = "Explore Features"
        text_color       = "#FFFFFF"
      }
      android = {
        button_action = "DEEP_LINK"
        link          = "myapp://features"
      }
      ios = {
        button_action = "DEEP_LINK"
        link          = "myapp://features"
      }
      web = {
        button_action = "DEEP_LINK"
        link          = "myapp://features"
      }
    }
    secondary_btn = {
      default_config = {
        background_color = "#808080"
        border_radius    = 8
        button_action    = "CLOSE"
        text             = "Maybe Later"
        text_color       = "#FFFFFF"
      }
    }
  }]

  custom_config = jsonencode({
    "DaysToDisplay" : 7,
    "DisplayTimes" : 3
  })
}