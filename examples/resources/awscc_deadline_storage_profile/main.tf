resource "awscc_deadline_farm" "example" {
  display_name = "ExampleRenderFarm"
  description  = "Example Deadline render farm for Linux storage profile"

  tags = [
    {
      key   = "ManagedBy"
      value = "AWSCC"
    }
  ]
}

resource "awscc_deadline_storage_profile" "example" {
  display_name = "Linux Storage Profile"
  farm_id      = awscc_deadline_farm.example.farm_id
  os_family    = "LINUX"

  file_system_locations = [
    {
      name = "SharedAssets"
      path = "/mnt/shared/assets"
      type = "SHARED"
    }
  ]
}