# Create DataSync location for FSx OpenZFS
resource "awscc_datasync_location_fsx_open_zfs" "example" {
  fsx_filesystem_arn  = "arn:aws:fsx:us-west-2:123456789012:file-system/fs-0123456789abcdef0"
  security_group_arns = ["arn:aws:ec2:us-west-2:123456789012:security-group/sg-0123456789abcdef0"]
  subdirectory        = "/fsx/example"
  protocol = {
    nfs = {
      mount_options = {
        version = "NFS3"
      }
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}