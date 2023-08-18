resource "awscc_efs_access_point" "this" {
  file_system_id = awscc_efs_file_system.this.id
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

resource "awscc_efs_file_system" "this" {

  file_system_tags = [
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