resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "amazon_linux_2023"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_bucket" "example" {
  bucket_name = "example-bucket"
  bundle_id   = "small_1_0"
  access_rules = {
    allow_public_overrides = false
    get_object             = "private"
  }
  object_versioning          = true
  read_only_access_accounts  = ["222222222222"]
  resources_receiving_access = [awscc_lightsail_instance.example.instance_name]
}
