resource "awscc_wisdom_assistant_association" "example" {
  assistant_id = awscc_wisdom_assistant.example.id # assistant_id works as well.
  association = {
    knowledge_base_id = awscc_wisdom_knowledge_base.example.id # knowledge_base_id works as well.
  }
  association_type = "KNOWLEDGE_BASE"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}