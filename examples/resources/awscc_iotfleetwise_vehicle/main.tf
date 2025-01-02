data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "aws_iam_role" "fleetwise" {
  name = "example-fleetwise-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iotfleetwise.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "fleetwise" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSIoTFleetWiseServiceRole"
  role       = aws_iam_role.fleetwise.name
}
resource "awscc_iotfleetwise_vehicle" "example" {
  name                 = "example-vehicle"
  decoder_manifest_arn = "arn:aws:iotfleetwise:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:decoder-manifest/example"
  model_manifest_arn   = "arn:aws:iotfleetwise:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:model-manifest/example"
  attributes = {
    model = "example-model"
    year  = "2024"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}