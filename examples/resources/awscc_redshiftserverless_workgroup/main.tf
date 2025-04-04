# Create namespace
resource "awscc_redshiftserverless_namespace" "example" {
  namespace_name = "example-namespace"
  db_name        = "exampledb"

  admin_username      = "admin"
  admin_user_password = "Admin123!"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create workgroup
resource "awscc_redshiftserverless_workgroup" "example" {
  workgroup_name      = "example-workgroup"
  namespace_name      = awscc_redshiftserverless_namespace.example.namespace_name
  base_capacity       = 32
  publicly_accessible = false

  config_parameters = [{
    parameter_key   = "enable_user_activity_logging"
    parameter_value = "true"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}