resource "awscc_datazone_environment_blueprint_configuration" "example" {
  domain_identifier                = awscc_datazone_domain.example.domain_id
  enabled_regions                  = ["us-east-1"]
  environment_blueprint_identifier = "DefaultDataLake"
  manage_access_role_arn           = awscc_iam_role.access.arn
  provisioning_role_arn            = awscc_iam_role.provisioning.arn
  regional_parameters = [{
    parameters = {
      "S3Location" : "s3://replace_with_bucket_name"
    }
    region = "us-east-1" # region is required for the datalake environment
  }]
}