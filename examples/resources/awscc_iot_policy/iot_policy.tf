resource "awscc_iot_policy" "example" {
  policy_name = "example"
  policy_document = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "iot:Connect"
          ],
          "Resource" : [
            "arn:aws:iot:us-east-1:123456789012:client/client1"
          ]
        }
      ]
    }
  )

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
