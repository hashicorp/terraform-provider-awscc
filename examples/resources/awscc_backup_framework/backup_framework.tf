resource "awscc_backup_framework" "example" {
  framework_controls = [{
    control_name = "BACKUP_PLAN_MIN_FREQUENCY_AND_MIN_RETENTION_CHECK"
    control_input_parameters = [{
      parameter_name  = "requiredRetentionDays"
      parameter_value = "35"
      },
      {
        parameter_name  = "requiredFrequencyValue"
        parameter_value = "1"
      },
      {
        parameter_name  = "requiredFrequencyUnit"
        parameter_value = "days"
    }]
    },
    {
      control_name = "BACKUP_LAST_RECOVERY_POINT_CREATED"
      control_input_parameters = [
        {
          parameter_name  = "recoveryPointAgeValue"
          parameter_value = "1"
        },
        {
          parameter_name  = "recoveryPointAgeUnit"
          parameter_value = "days"
      }],
      control_scope = {
        compliance_resource_types = [
          "RDS",
          "S3",
          "Aurora",
          "EFS",
          "EC2",
          "Storage Gateway",
          "EBS",
          "DynamoDB",
          "FSx",
          "VirtualMachine"
        ]
      }
    }
  ]
  framework_description = "example framework"
  framework_name        = "example"
  framework_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
