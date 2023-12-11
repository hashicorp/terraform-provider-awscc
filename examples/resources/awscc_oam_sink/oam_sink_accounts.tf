resource "awscc_oam_sink" "example" {
  name = "SampleSink"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Resource = "*"
      Action = [
        "oam:CreateLink",
        "oam:UpdateLink"
      ]
      Principal = {
        AWS = ["1111111111111"]
      }
      Condition = {
        "ForAllValues:StringEquals" : {
          "oam:ResourceTypes" : [
            "AWS::CloudWatch::Metric",
            "AWS::Logs::LogGroup",
            "AWS::XRay::Trace"
          ]
        }
      }
    }]
  })
}