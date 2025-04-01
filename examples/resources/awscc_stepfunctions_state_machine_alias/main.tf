data "aws_region" "current" {}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["states.${data.aws_region.current.name}.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

# Create IAM role for Step Functions
resource "awscc_iam_role" "step_function_role" {
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  description                 = "Role for Step Functions State Machine"
  path                        = "/"
  role_name                   = "MyStepFunctionsRole"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Step Functions State Machine
resource "awscc_stepfunctions_state_machine" "example" {
  role_arn = awscc_iam_role.step_function_role.arn
  definition = jsonencode({
    "Comment" : "A Hello World example",
    "StartAt" : "HelloWorld",
    "States" : {
      "HelloWorld" : {
        "Type" : "Pass",
        "End" : true
      }
    }
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create State Machine Version
resource "awscc_stepfunctions_state_machine_version" "example" {
  state_machine_arn = awscc_stepfunctions_state_machine.example.arn
  description       = "Initial version"
}

# Create State Machine Alias
resource "awscc_stepfunctions_state_machine_alias" "example" {
  name        = "prod"
  description = "Production alias for the state machine"

  routing_configuration = [{
    state_machine_version_arn = awscc_stepfunctions_state_machine_version.example.arn
    weight                    = 100
  }]
}