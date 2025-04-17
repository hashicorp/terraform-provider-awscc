# Create a Cognito User Pool first
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = false
    require_uppercase = true
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "email"
    required            = true
    string_attribute_constraints {
      max_length = "2048"
      min_length = "0"
    }
  }

  mfa_configuration = "OFF"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a client app using AWSCC provider
resource "awscc_cognito_user_pool_client" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_name  = "example-client"

  generate_secret               = false
  prevent_user_existence_errors = "ENABLED"
  callback_ur_ls                = ["https://example.com/callback"]
  allowed_o_auth_flows          = ["code"]
  allowed_o_auth_scopes         = ["openid"]
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]
  allowed_o_auth_flows_user_pool_client = true
}

# Create branding configuration
resource "awscc_cognito_managed_login_branding" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_id    = split("|", awscc_cognito_user_pool_client.example.id)[1]

  # Note: Settings format depends on the specific Cognito UI version you're using
  settings = jsonencode({})

  assets = [{
    category   = "PAGE_HEADER_LOGO"
    extension  = "PNG"
    bytes      = "iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAIAAAD/gAIDAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA2ZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMC1jMDYwIDYxLjEzNDc3NywgMjAxMC8wMi8xMi0xNzozMjowMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDpGNzdGMTE3NDA3MjA2ODExOEMxNEE3NkIxRDhEMzU5RSIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDpBQzU0OTI2RjY5MjAxMUUyQjM1OUE4QzQwMEM2QjM0MCIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDpBQzU0OTI2RTY5MjAxMUUyQjM1OUE4QzQwMEM2QjM0MCIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ1M1IE1hY2ludG9zaCI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjA1ODAxMTc0MDcyMDY4MTE4QzE0QTc2QjFEOEQzNTlFIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOkY3N0YxMTc0MDcyMDY4MTE4QzE0QTc2QjFEOEQzNTlFIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+SG1RYgAAABl0RVh0U29mdHdhcmUAQWRvYmUgSW1hZ2VSZWFkeXHJZTwAAADqSURBVHjaYmRmYaSN4GVkYKKxASxMtDWAlYW2BrCy0tYANjbaGsDORlsDeNhpawAvO20N4OOgrQF8XLQ1gJ+btgYI8dLWAGFe2hogykdbA8T4aGuAOB9tDZDgo60BUvy0NUCWn7YGSAvQ1gAZAdoaICtIWwPkBGlrgIIQbQ1QFqatAWrCtDVAU4S2BmiJ0tYAHVHaGqArSlsDDMVpa4CROG0NMBanrQFmUrQ1wFKKtgbYStHWADsp2hrgJEVbA1ylqG0AP1MDCwsLjdMGAwM9i0pmNnoVlexspL0RebhobQAvUwAAAgwA3q4UUqwxJWsAAAAASUVORK5CYII="
    color_mode = "LIGHT"
  }]
}