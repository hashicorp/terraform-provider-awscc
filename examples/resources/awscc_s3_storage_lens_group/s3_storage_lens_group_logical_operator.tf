resource "awscc_s3_storage_lens_group" "example" {
  name = "example"
  filter = {
    and = {
      match_any_prefix = ["group1"],
      match_object_age = {
        days_greater_than = 10,
        days_less_than    = 60
      },
      match_object_size = {
        bytes_greater_than = 10,
        bytes_less_than    = 60
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
