# Encryption SecurityPolicy
resource "awscc_opensearchserverless_security_policy" "security_policy" {
  name        = "awscc-security-policy"
  description = "created via awscc"
  type        = "encryption"
  policy = jsonencode({
    "Rules" = [
      {
        "ResourceType" = "collection",
        "Resource" = [
          "collection/awscc-collection"
        ]
      }
    ],
    "AWSOwnedKey" = true
  })
}