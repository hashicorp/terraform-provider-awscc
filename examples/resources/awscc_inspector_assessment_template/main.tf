# Variable for rules package ARN
variable "rules_package_arn" {
  description = "The ARN of the Inspector rules package to use for the assessment"
  type        = string
  default     = "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p" # Common Vulnerabilities and Exposures for us-west-2
}

# Create an Inspector Assessment Target
resource "awscc_inspector_assessment_target" "example" {
  assessment_target_name = "example-assessment-target"
  resource_group_arn     = null # Set to null to target all EC2 instances
}

# Create an Inspector Assessment Template
resource "awscc_inspector_assessment_template" "example" {
  assessment_target_arn    = awscc_inspector_assessment_target.example.arn
  assessment_template_name = "example-assessment-template"
  duration_in_seconds      = 3600 # 1 hour
  rules_package_arns       = [var.rules_package_arn]

  user_attributes_for_findings = [{
    key   = "Environment"
    value = "Production"
  }]
}