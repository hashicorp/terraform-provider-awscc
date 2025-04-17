# Create an Amazon Connect instance
resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias          = "task-template-example"
  
  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }

  tags = [{
    key   = "Type"
    value = "Example"
  }]
}

# Create a basic contact flow
resource "awscc_connect_contact_flow" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "Task Template Example Flow"
  type         = "CONTACT_FLOW"
  content = jsonencode({
    "Version" : "2019-10-30",
    "StartAction" : "12345678-1234-1234-1234-123456789012",
    "Actions" : [
      {
        "Identifier" : "12345678-1234-1234-1234-123456789012",
        "Type" : "DisconnectParticipant"
      }
    ]
  })

  tags = [{
    key   = "Type"
    value = "Example"
  }]
}

# Create the task template
resource "awscc_connect_task_template" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "Example Task Template"
  description  = "An example task template"
  status       = "ACTIVE"

  contact_flow_arn = awscc_connect_contact_flow.example.contact_flow_arn

  fields = [
    {
      description = "Customer Name"
      id = {
        name = "CustomerName"
      }
      type = "NAME"
    },
    {
      description = "Issue Category"
      id = {
        name = "IssueCategory"
      }
      type = "SINGLE_SELECT"
      single_select_options = [
        "Billing",
        "Technical",
        "Account",
        "Other"
      ]
    }
  ]

  constraints = {
    required_fields = [
      {
        id = {
          name = "CustomerName"
        }
      }
    ]
  }

  defaults = [
    {
      id = {
        name = "IssueCategory"
      }
      default_value = "Other"
    }
  ]

  tags = [{
    key   = "Type"
    value = "Example"
  }]
}