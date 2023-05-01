# Create a Collection
resource "awscc_opensearchserverless_collection" "simple_collection" {
    name = "awscc-collection"
    depends_on = [
      awscc_opensearchserverless_security_policy.security_policy
    ]
}

# Encryption SecurityPolicy
resource "awscc_opensearchserverless_security_policy" "security_policy" {
    name = "awscc-security-policy"
    description = "created via awscc"
    type = "encryption"
    policy = "{\"Rules\":[{\"ResourceType\":\"collection\",\"Resource\":[\"collection/awscc-collection\"]}],\"AWSOwnedKey\":true}"
}