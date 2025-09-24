resource "awscc_connect_queue" "example" {
  instance_arn           = awscc_connect_instance.example.arn
  name                   = "example"
  description            = "Example Description"
  hours_of_operation_arn = awscc_connect_hours_of_operation.example.hours_of_operation_arn
  status                 = "ENABLED"
  max_contacts           = 10
  outbound_caller_config = {
    outbound_caller_id_name       = "outbound"
    outbound_caller_id_number_arn = awscc_connect_phone_number.example.phone_number_arn
    outbound_flow_arn             = awscc_connect_contact_flow.example.contact_flow_arn
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
