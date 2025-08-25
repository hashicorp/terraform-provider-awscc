# Neptune DB Subnet Group
resource "awscc_neptune_db_subnet_group" "example" {
  db_subnet_group_description = "Example Neptune DB subnet group"
  db_subnet_group_name        = "example-neptune-subnet-group"
  subnet_ids                  = ["subnet-09110354a9ca2df61", "subnet-080f5e82b355239f5", "subnet-0c8fe0ba56847ef42"]
  tags = [
    {
      key   = "Name"
      value = "example-neptune-subnet-group"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Neptune DB Cluster
resource "awscc_neptune_db_cluster" "example" {
  db_subnet_group_name  = awscc_neptune_db_subnet_group.example.db_subnet_group_name
  db_cluster_identifier = "example-neptune-cluster"
  storage_encrypted     = true
  tags = [
    {
      key   = "Name"
      value = "example-neptune-cluster"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Neptune DB Instance
resource "awscc_neptune_db_instance" "example" {
  db_instance_class      = "db.r5.large"
  db_instance_identifier = "example-neptune-instance"
  db_cluster_identifier  = awscc_neptune_db_cluster.example.db_cluster_identifier
  tags = [
    {
      key   = "Name"
      value = "example-neptune-instance"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
