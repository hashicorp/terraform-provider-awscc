
# Data sources to get region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# IAM role for IoT Fleet Provisioning
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iot.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "provisioning" {
  statement {
    effect = "Allow"
    actions = [
      "iot:CreateThing",
      "iot:CreateThingGroup",
      "iot:AddThingToThingGroup",
      "iot:AttachThingPrincipal",
      "iot:CreateKeysAndCertificate",
      "iot:UpdateCertificate"
    ]
    resources = [
      "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:thing/*",
      "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:thinggroup/*",
      "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cert/*"
    ]
  }
}

resource "aws_iam_role" "provisioning" {
  name               = "IoTFleetProvisioningRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json

  inline_policy {
    name   = "IoTFleetProvisioning"
    policy = data.aws_iam_policy_document.provisioning.json
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# IoT Fleet Provisioning Template
resource "awscc_iot_provisioning_template" "example" {
  provisioning_role_arn = aws_iam_role.provisioning.arn
  template_name         = "example-fleet-template"
  description           = "Example IoT Fleet Provisioning Template"
  enabled               = true
  template_body = jsonencode({
    Parameters = {
      SerialNumber   = { Type = "String" }
      DeviceLocation = { Type = "String" }
    }
    Resources = {
      certificate = {
        Type = "AWS::IoT::Certificate"
        Properties = {
          CertificateId = { Ref = "AWS::IoT::Certificate::Id" }
          Status        = "ACTIVE"
        }
      }
      thing = {
        Type = "AWS::IoT::Thing"
        Properties = {
          ThingName = { Ref = "SerialNumber" }
          AttributePayload = {
            location = { Ref = "DeviceLocation" }
          }
        }
      }
    }
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}