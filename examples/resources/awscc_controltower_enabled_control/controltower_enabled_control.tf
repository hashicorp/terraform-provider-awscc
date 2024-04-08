


# The following resource enables the control 'AWS-GR_AUDIT_BUCKET_RETENTION_POLICY'

resource "awscc_controltower_enabled_control" "example" {
  control_identifier = "arn:aws:controltower:us-east-1::control/AWS-GR_AUDIT_BUCKET_RETENTION_POLICY"
  target_identifier  = "arn:aws:organizations::000123456789:ou/o-ydkk7uvfj1/ou-e3wm-2x760yqd"
}

# Please change the Organization ID from 'o-ydkk7uvfj1' to your Organization ID, and the Organizational Unit ID from 'u-e3wm-2x760yqd' to your desired OU where the controls need to be implemented.
