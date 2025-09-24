resource "awscc_mediapackagev2_channel_policy" "example" {
  channel_group_name = awscc_mediapackagev2_channel_group.example.channel_group_name
  channel_name       = awscc_mediapackagev2_channel.example.channel_name
  policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Id" : "AllowMediaLiveChannelToIngestToEmpChannel",
      "Statement" : [
        {
          "Sid" : "AllowMediaLiveRoleToAccessEmpChannel",
          "Effect" : "Allow",
          "Principal" : {
            "AWS" : awscc_iam_role.example.arn
          },
          "Action" : "mediapackagev2:PutObject",
          "Resource" : awscc_mediapackagev2_channel.example.arn,
          "Condition" : {
            "IpAddress" : {
              "aws:SourceIp" : "0.0.0.0/24"
            }
          }
        }
      ]
    }
  )
}