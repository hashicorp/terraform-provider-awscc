resource "awscc_inspectorv2_filter" "example" {
  name          = "example"
  description   = "Suppression rule filtered on ECR Image tag"
  filter_action = "NONE"
  filter_criteria = {
    ecr_image_tags = [{
      comparison = "EQUALS"
      value      = "v1.0"
    }]
  }

}