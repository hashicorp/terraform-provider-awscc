resource "awscc_efs_access_point" "this" {
  file_system_id = "fs-0dbbfe225fc860016" #awscc_efs_file_system.this.id
  posix_user = {
    gid = 0
    uid = 0
  }

  access_point_tags = [
    {
      key   = "Name"
      value = "this"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}