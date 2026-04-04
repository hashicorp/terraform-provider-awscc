resource "awscc_elementalinference_feed" "example" {
  name = "example-inference-feed"

  outputs = [
    {
      name   = "example-output"
      status = "ENABLED"
      output_config = {
        client_configuration = {
          hls_settings = {
            rendition_groups = [
              {
                name                = "example-group"
                rendition_selection = "ALL"
              }
            ]
          }
        }
      }
    }
  ]

  sources = [
    {
      name = "example-source"
      url  = "https://example.com/feed"
    }
  ]

  tags = {
    Environment = "test"
    Purpose     = "example"
  }
}