resource "awscc_iam_role" "main" {
  description = "AWS IAM role for a step function"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "states.amazonaws.com"
        }
      },
    ]
  })
}

resource "awscc_stepfunctions_state_machine" "sfn_stepmachine" {
  role_arn           = awscc_iam_role.main.arn
  state_machine_type = "STANDARD"
  definition_string  = <<EOT
    {
      "StartAt": "FirstState",
      "States": {
        "FirstState": {
          "Type": "Pass",
          "Result": "Hello World!",
          "End": true
        }
      }
    }
  EOT
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_stepfunctions_state_machine_version" "version" {
  description       = "State machine version description"
  state_machine_arn = awscc_stepfunctions_state_machine.sfn_stepmachine.arn
}