resource "awscc_qbusiness_retriever" "example" {
  application_id = awscc_qbusiness_application.example.application_id
  display_name   = "example_q_retriever"
  type           = "NATIVE_INDEX"

  configuration = {
    native_index_configuration = {
      index_id = awscc_qbusiness_index.example.index_id
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
