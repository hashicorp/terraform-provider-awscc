resource "awscc_smsvoice_opt_out_list" "example" {
  opt_out_list_name = "example-opt-out-list"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-opt-out-list"
    }
  ]
}