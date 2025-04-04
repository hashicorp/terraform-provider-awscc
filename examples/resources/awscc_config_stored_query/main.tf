# Example stored query for Config service
resource "awscc_config_stored_query" "example" {
  query_name        = "find-running-instances"
  query_description = "Finds all running EC2 instances in the current region"
  query_expression  = "SELECT configuration.instanceType, configuration.imageId, tags, configuration.instanceId WHERE resourceType = 'AWS::EC2::Instance' AND configuration.state.name = 'running'"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}