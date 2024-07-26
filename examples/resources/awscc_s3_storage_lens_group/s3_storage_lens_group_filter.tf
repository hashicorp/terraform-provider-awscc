resource "awscc_s3_storage_lens_group" "example" {
  name = "example"
  filter = {
    match_any_tag = [
      {
        key   = "key1",
        value = "value1"
      },
      {
        key   = "key2",
        value = "value2"
      }
    ]
  }


  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
