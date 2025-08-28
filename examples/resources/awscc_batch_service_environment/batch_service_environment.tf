# AWS Batch Service Environment resource
resource "awscc_batch_service_environment" "example" {
  service_environment_name = "example-service-env"
  service_environment_type = "SAGEMAKER_TRAINING"
  state                    = "ENABLED"

  capacity_limits = [
    {
      capacity_unit = "NUM_INSTANCES"
      max_capacity  = 50
    }
  ]

  tags = {
    Environment = "example"
    Name        = "example-service-env"
  }
}
