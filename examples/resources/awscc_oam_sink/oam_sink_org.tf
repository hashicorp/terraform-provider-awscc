data "aws_organizations_organization" "example" {}

resource "awscc_oam_sink" "example" {
  name = "SampleSink"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect    = "Allow"
      Principal = "*"
      Resource  = "*"
      Action    = ["oam:CreateLink", "oam:UpdateLink"]
      Condition = {
        StringEquals = {
          "aws:PrincipalOrgID" = data.aws_organizations_organization.example.id
        }
        "ForAllValues:StringEquals" = {
          "oam:ResourceTypes" = [
            "AWS::CloudWatch::Metric",
            "AWS::Logs::LogGroup"
          ]
        }
      }
    }]
  })
}