# Main IoT Thing Group
resource "awscc_iot_thing_group" "example" {
  thing_group_name = "example-thing-group"

  thing_group_properties = {
    thing_group_description = "Example IoT Thing Group created by AWSCC provider"
    attribute_payload = {
      attributes = {
        "Environment" = "Test"
        "Project"     = "Demo"
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Optional nested group example
resource "awscc_iot_thing_group" "child_group" {
  thing_group_name  = "example-child-thing-group"
  parent_group_name = awscc_iot_thing_group.example.thing_group_name

  thing_group_properties = {
    thing_group_description = "Child IoT Thing Group example"
    attribute_payload = {
      attributes = {
        "Environment" = "Test"
        "Parent"      = awscc_iot_thing_group.example.thing_group_name
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}