resource "awscc_sagemaker_cluster" "this" {
  cluster_name = "example"
  instance_groups = [
    {
      execution_role      = awscc_iam_role.example.arn
      instance_count      = 1
      instance_type       = "ml.c5.2xlarge"
      instance_group_name = "example"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.this.id}/config/"
        on_create     = "on_create_noop.sh"
      }
    }
  ]
  orchestrator = {
    eks = {
      cluster_arn = "arn:${data.aws_partition.current.partition}:eks:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster/hyperpod-eks-example"
    }
  }
  vpc_config = {
    security_group_ids = [var.sg_id]
    subnets            = [var.subnet_id]
  }

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

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_partition" "current" {}
