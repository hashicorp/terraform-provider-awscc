# Redshift cluster parameter group
resource "awscc_redshift_cluster_parameter_group" "example" {
  parameter_group_name   = "example-paramgroup"
  parameter_group_family = "redshift-1.0"
  description            = "Example Redshift cluster parameter group"

  parameters = [
    {
      parameter_name  = "enable_user_activity_logging"
      parameter_value = "true"
    },
    {
      parameter_name  = "require_ssl"
      parameter_value = "true"
    }
  ]

  tags = [
    {
      key   = "Environment"
      value = "Test"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}