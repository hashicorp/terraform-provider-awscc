resource "awscc_bedrock_automated_reasoning_policy" "example" {
  name        = "example-automated-reasoning-policy"
  description = "Example automated reasoning policy for access control validation"

  policy_definition = {
    version = "1.0"

    types = [
      {
        name        = "UserRole"
        description = "Valid user roles"
        values = [
          {
            value       = "admin"
            description = "Administrator"
          },
          {
            value       = "user"
            description = "Regular user"
          }
        ]
      }
    ]

    variables = [
      {
        name        = "user_role"
        type        = "UserRole"
        description = "The user's role"
      }
    ]

    rules = [
      {
        id         = "A1DMINACCESS"
        expression = "true"
      }
    ]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}