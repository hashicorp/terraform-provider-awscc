# Redshift Serverless namespace
resource "awscc_redshiftserverless_namespace" "example" {
  namespace_name = "example-namespace"

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Name"
      value = "example-namespace"
    }
  ]
}

# Redshift Serverless workgroup
resource "awscc_redshiftserverless_workgroup" "example" {
  workgroup_name = "example-workgroup"
  namespace_name = awscc_redshiftserverless_namespace.example.namespace_name

  base_capacity = 8

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Name"
      value = "example-workgroup"
    }
  ]
}

# Redshift Serverless snapshot 
resource "awscc_redshiftserverless_snapshot" "example" {
  snapshot_name  = "example-snapshot"
  namespace_name = awscc_redshiftserverless_namespace.example.namespace_name

  retention_period = 7

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Name"
      value = "example-snapshot"
    }
  ]

  depends_on = [awscc_redshiftserverless_workgroup.example]
}
