# Create S3 bucket for job attachments
resource "awscc_s3_bucket" "example" {
  bucket_name = "deadline-job-attachments-${random_id.bucket_suffix.hex}"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Generate random suffix for bucket name uniqueness
resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "awscc_deadline_farm" "example" {
  display_name = "ExampleRenderFarm"
  description  = "Example Deadline Farm for queue demonstration"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create storage profiles for different operating systems
resource "awscc_deadline_storage_profile" "linux_storage" {
  display_name = "Linux Shared Storage"
  farm_id      = awscc_deadline_farm.example.farm_id
  os_family    = "LINUX"

  file_system_locations = [{
    name = "shared storage"
    path = "/mnt/shared"
    type = "SHARED"
    }, {
    name = "render assets"
    path = "/mnt/assets"
    type = "SHARED"
  }]
}

resource "awscc_deadline_storage_profile" "windows_storage" {
  display_name = "Windows Shared Storage"
  farm_id      = awscc_deadline_farm.example.farm_id
  os_family    = "WINDOWS"

  file_system_locations = [{
    name = "shared storage"
    path = "Z:\\"
    type = "SHARED"
    }, {
    name = "render assets"
    path = "Y:\\"
    type = "SHARED"
  }]
}

# Create an advanced Deadline Queue with job attachment settings
resource "awscc_deadline_queue" "example" {
  display_name          = "AdvancedRenderQueue"
  description           = "Advanced render queue with S3 job attachments and custom settings"
  farm_id               = awscc_deadline_farm.example.farm_id
  default_budget_action = "STOP_SCHEDULING_AND_COMPLETE_TASKS"

  # Configure job attachment settings for S3
  job_attachment_settings = {
    s3_bucket_name = awscc_s3_bucket.example.bucket_name
    root_prefix    = "job-attachments/"
  }

  # Configure job run-as user settings for POSIX systems
  job_run_as_user = {
    run_as = "QUEUE_CONFIGURED_USER"
    posix = {
      user  = "deadline-worker"
      group = "deadline-group"
    }
  }

  # Specify allowed storage profile IDs (dynamically referenced)
  allowed_storage_profile_ids = [
    awscc_deadline_storage_profile.linux_storage.storage_profile_id,
    awscc_deadline_storage_profile.windows_storage.storage_profile_id
  ]

  # Specify required file system location names
  required_file_system_location_names = [
    "shared storage",
    "render assets"
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}