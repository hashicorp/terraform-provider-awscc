# Create Cognito Identity Pool
resource "awscc_cognito_identity_pool" "example" {
  identity_pool_name               = "example_identity_pool"
  allow_unauthenticated_identities = false
}

# Create IAM role for authenticated users
resource "awscc_iam_role" "authenticated" {
  role_name = "cognito_authenticated_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Federated = "cognito-identity.amazonaws.com"
        }
        Action = "sts:AssumeRoleWithWebIdentity"
        Condition = {
          StringEquals = {
            "cognito-identity.amazonaws.com:aud" = awscc_cognito_identity_pool.example.id
          }
          "ForAnyValue:StringLike" = {
            "cognito-identity.amazonaws.com:amr" : "authenticated"
          }
        }
      }
    ]
  })

  policies = [
    {
      policy_name = "authenticated_policy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "mobileanalytics:PutEvents",
              "cognito-sync:*"
            ]
            Resource = ["*"]
          }
        ]
      })
    }
  ]
}

# Create IAM role for unauthenticated users
resource "awscc_iam_role" "unauthenticated" {
  role_name = "cognito_unauthenticated_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Federated = "cognito-identity.amazonaws.com"
        }
        Action = "sts:AssumeRoleWithWebIdentity"
        Condition = {
          StringEquals = {
            "cognito-identity.amazonaws.com:aud" = awscc_cognito_identity_pool.example.id
          }
          "ForAnyValue:StringLike" = {
            "cognito-identity.amazonaws.com:amr" : "unauthenticated"
          }
        }
      }
    ]
  })

  policies = [
    {
      policy_name = "unauthenticated_policy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "mobileanalytics:PutEvents"
            ]
            Resource = ["*"]
          }
        ]
      })
    }
  ]
}

# Create the role attachment
resource "awscc_cognito_identity_pool_role_attachment" "example" {
  identity_pool_id = awscc_cognito_identity_pool.example.id
  roles = {
    "authenticated"   = awscc_iam_role.authenticated.arn
    "unauthenticated" = awscc_iam_role.unauthenticated.arn
  }
}