resource "awscc_evidently_segment" "example" {
  name        = "example-segment"
  description = "Example segment created using AWSCC provider"
  pattern     = "browser.* = \"Chrome\""

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

output "segment_arn" {
  description = "ARN of the created Evidently segment"
  value       = awscc_evidently_segment.example.arn
}