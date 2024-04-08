


# The following resource enables the control 'AWS-GR_AUDIT_BUCKET_RETENTION_POLICY'

resource "awscc_controltower_enabled_control" "example" {
  control_identifier = "arn:aws:controltower:us-east-1::control/AWS-GR_AUDIT_BUCKET_RETENTION_POLICY"
  target_identifier  = "arn:aws:organizations::<<your-account-id>:ou/<<your-org-id>>/<<your-ou-id>>"
}

# Please change the Organization ID to your Organization ID, and the Organizational Unit ID to your desired OU where the controls need to be implemented.
