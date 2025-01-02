data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create the Inspector Assessment Target
resource "awscc_inspector_assessment_target" "example" {
  assessment_target_name = "example-assessment-target"
}