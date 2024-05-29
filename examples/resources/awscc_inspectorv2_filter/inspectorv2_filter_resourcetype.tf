resource "awscc_inspectorv2_filter" "example" {
  name          = "example"
  description   = "Suppression rule filtered on EC2 Instance and Lambda function "
  filter_action = "NONE"
  filter_criteria = {
    resource_type = [
      {
        comparison = "EQUALS"
        value      = "Amazon::Lambda:Function"
      },
      {
        comparison = "EQUALS"
        value      = "Amazon::EC2:Instance"
      }
    ]

  }

}