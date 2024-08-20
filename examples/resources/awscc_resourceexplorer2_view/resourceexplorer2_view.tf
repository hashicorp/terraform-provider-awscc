resource "awscc_resourceexplorer2_view" "example" {
  view_name = "example"
  included_properties = [
    {
      name = "tags"
    }
  ]

  depends_on = [awscc_resourceexplorer2_index.example]
}

resource "awscc_resourceexplorer2_index" "example" {
  type = "LOCAL"
}