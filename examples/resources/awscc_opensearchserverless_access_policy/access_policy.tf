resource "awscc_opensearchserverless_access_policy" "os" {
  name        = "test-os-security-policy"
  type        = "data"
  description = "Access for test-user"
  policy = jsonencode([{
    "Description" = "Access for test-user",
    "Rules" = [
      {
        "ResourceType" = "index",
        "Resource" = [
          "index/*/*"
        ],
        "Permission" = [
          "aoss:*"
        ]
      },
      {
        "ResourceType" = "collection",
        "Resource" = [
          "collection/my-collection"
        ],
        "Permission" = [
          "aoss:*"
        ]
    }],
    "Principal" = [
      "arn:aws:iam::111122223333:user/test-user"
    ]
  }])
}