resource "awscc_ses_custom_verification_email_template" "example" {
  template_name             = "example-verification-template"
  template_subject          = "Please verify your email address"
  from_email_address        = "wellsiau@amazon.com"
  success_redirection_url   = "https://example.com/success"
  failure_redirection_url   = "https://example.com/failure"
  template_content         = <<-EOT
    Please verify your email address by clicking the following link:
    
    {{VerificationURL}}
    
    If you did not request this verification, please ignore this email.
    
    Thank you!
  EOT
}

output "custom_verification_email_template_name" {
  description = "The name of the custom verification email template"
  value       = awscc_ses_custom_verification_email_template.example.template_name
}