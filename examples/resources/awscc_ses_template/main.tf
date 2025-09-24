# Create a basic SES template
resource "awscc_ses_template" "example" {
  template = {
    template_name = "example-template"
    subject_part  = "Welcome to Our Service!"
    html_part     = <<EOF
<!DOCTYPE html>
<html>
<body>
    <h1>Welcome to Our Service!</h1>
    <p>Hello {{name}},</p>
    <p>Thank you for signing up. We're excited to have you on board!</p>
    <p>Best regards,<br>The Team</p>
</body>
</html>
EOF
    text_part     = <<EOF
Welcome to Our Service!

Hello {{name}},

Thank you for signing up. We're excited to have you on board!

Best regards,
The Team
EOF
  }
}
