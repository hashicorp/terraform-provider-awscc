# Get current AWS region
data "aws_region" "current" {}

# Create SageMaker Model Package Group
resource "awscc_sagemaker_model_package_group" "example" {
  model_package_group_name        = "example-model-package-group"
  model_package_group_description = "Example SageMaker Model Package Group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create SageMaker Model Package
resource "awscc_sagemaker_model_package" "example" {
  model_package_group_name  = awscc_sagemaker_model_package_group.example.model_package_group_name
  model_package_description = "Example SageMaker Model Package"

  inference_specification = {
    containers = [
      {
        image             = "763104351884.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/tensorflow-inference:2.12.1-cpu"
        framework         = "TENSORFLOW"
        framework_version = "2.12.1"
      }
    ]
    supported_content_types                     = ["application/json"]
    supported_response_mime_types               = ["application/json"]
    supported_realtime_inference_instance_types = ["ml.t2.medium"]
  }

  model_approval_status = "PendingManualApproval"

  customer_metadata_properties = {
    "Domain"    = "Computer Vision"
    "Framework" = "TensorFlow"
  }
}