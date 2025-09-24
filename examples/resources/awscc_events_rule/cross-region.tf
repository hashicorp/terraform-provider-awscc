data "aws_caller_identity" "current" {}

resource "awscc_events_event_bus" "example" {
  name = "CrossRegionDestinationBus"
}

resource "awscc_events_rule" "example" {
  description    = "Routes to us-east-1 event bus"
  event_bus_name = "MyBusName"
  state          = "ENABLED"

  event_pattern = jsonencode({
    "source" : ["MyTestApp"],
    "detail" : ["MyTestAppDetail"]
  })

  targets = [{
    arn      = awscc_events_event_bus.example.arn
    id       = "CrossRegionDestinationBus"
    role_arn = awscc_iam_role.example.arn
  }]
}

resource "awscc_iam_role" "example" {
  assume_role_policy_document = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "events.amazonaws.com"
        },
        "Action" : "sts:AssumeRole"
      }
    ]
  })

  path = "/"
  policies = [
    {
      policy_document = data.aws_iam_policy_document.example.json
      policy_name     = "PutEventsDestinationBus"
    }
  ]
}

data "aws_iam_policy_document" "example" {
  statement {
    sid       = "PutEventsDestinationBus"
    effect    = "Allow"
    actions   = ["events:PutEvents"]
    resources = [awscc_events_event_bus.example.arn]
  }
}