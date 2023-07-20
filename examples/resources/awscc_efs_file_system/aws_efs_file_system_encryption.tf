resource "awscc_efs_file_system" "this" {
  encrypted  = true
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/b1d4919e-3296-4104-a3a8-c9f3b1138fa8"

  file_system_tags = [{
    key   = "Name"
    value = "this"
  }]
}