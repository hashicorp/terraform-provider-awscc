resource "awscc_sagemaker_cluster" "example" {
  cluster_name = "example"
  instance_groups = [
    {
      execution_role      = awscc_iam_role.example.arn
      instance_count      = 1
      instance_type       = "ml.c5.2xlarge"
      instance_group_name = "example"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.example.id}/config/"
        on_create     = "on_create_noop.sh"
      }
      instance_storage_configs = [{
        ebs_volume_config = {
          volume_size_in_gb = 30
        }
      }]
    }
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]

}

resource "aws_s3_bucket" "example" {
  bucket = "example"
}

resource "aws_s3_object" "script" {
  bucket = aws_s3_bucket.example.id
  key    = "config/on_create_noop.sh"
  source = "on_create_noop.sh"
}

resource "aws_s3_object" "params" {
  bucket = aws_s3_bucket.example.id
  key    = "config/provisioning_parameters.json"
  source = "provisioning_parameters.json"
}