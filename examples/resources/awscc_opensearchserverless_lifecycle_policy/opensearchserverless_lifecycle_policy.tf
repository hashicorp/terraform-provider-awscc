resource "awscc_opensearchserverless_lifecycle_policy" "example" {
  name        = "example-lifecycle-policy"
  description = "OpenSearch Serverless lifecycle policy"
  type        = "retention"
  policy = jsonencode(
    { "Rules" : [
      {
        "Resource" : ["index/my-collection/*"],
        "ResourceType" : "index",
        "MinIndexRetention" : "2d"
      }
      ]
    }
  )
}
