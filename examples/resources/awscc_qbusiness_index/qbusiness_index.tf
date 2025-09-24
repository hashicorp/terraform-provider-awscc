resource "awscc_qbusiness_index" "example" {
  application_id = awscc_qbusiness_application.example.application_id
  display_name   = "example_q_index"
  description    = "Example QBusiness Index"
  type           = "ENTERPRISE"
  capacity_configuration = {
    units = 1
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
