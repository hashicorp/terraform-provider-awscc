resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-connect-instance"

  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

resource "awscc_connect_data_table_record" "example" {
  instance_id = awscc_connect_instance.example.id
  
  contact_id = "12345678-1234-1234-1234-123456789012"
  
  data = {
    customer_name = "John Doe"
    phone_number = "+1234567890"
    email = "john.doe@example.com"
    product_interest = "Premium Service"
    call_reason = "Product Inquiry"
    notes = "Customer interested in upgrading to premium service package"
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Outputs
output "connect_instance_id" {
  description = "ID of the Connect instance"
  value       = awscc_connect_instance.example.id
}

output "connect_instance_arn" {
  description = "ARN of the Connect instance"
  value       = awscc_connect_instance.example.arn
}

output "data_table_record_id" {
  description = "ID of the data table record"
  value       = awscc_connect_data_table_record.example.id
}

output "data_table_record_instance_id" {
  description = "Instance ID associated with the data table record"
  value       = awscc_connect_data_table_record.example.instance_id
}

output "data_table_record_contact_id" {
  description = "Contact ID associated with the data table record"
  value       = awscc_connect_data_table_record.example.contact_id
}

output "data_table_record_data" {
  description = "Data stored in the table record"
  value       = awscc_connect_data_table_record.example.data
}