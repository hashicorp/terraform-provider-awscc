# First create a view to associate with
resource "awscc_resourceexplorer2_view" "example" {
  view_name = "example-view"
  filters = {
    filter_string = "resourcetype:AWS::EC2::Instance"
  }
  included_properties = [{
    name = "tags"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the default view association
resource "awscc_resourceexplorer2_default_view_association" "example" {
  view_arn = awscc_resourceexplorer2_view.example.view_arn
}