# Create a patch baseline for Windows Server
resource "awscc_ssm_patch_baseline" "windows_baseline" {
  name        = "WindowsServerPatchBaseline"
  description = "Patch baseline for Windows Server"

  operating_system                  = "WINDOWS"
  approved_patches_compliance_level = "HIGH"

  # Global filters for the patch baseline
  global_filters = {
    patch_filters = [
      {
        key    = "PRODUCT"
        values = ["WindowsServer2019"]
      },
      {
        key    = "CLASSIFICATION"
        values = ["CriticalUpdates", "SecurityUpdates"]
      }
    ]
  }

  # Approval rules for patches
  approval_rules = {
    patch_rules = [
      {
        approve_after_days = 7
        compliance_level   = "HIGH"

        patch_filter_group = {
          patch_filters = [
            {
              key    = "MSRC_SEVERITY"
              values = ["Critical", "Important"]
            },
            {
              key    = "PATCH_SET"
              values = ["OS"]
            }
          ]
        }
      }
    ]
  }

  # Tag the resource
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}