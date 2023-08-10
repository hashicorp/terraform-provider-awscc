resource "awscc_connect_approved_origin" "this" {
  instance_id = "arn:aws:connect:us-east-1:111122223333:instance/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
  origin      = "https://example.com"
}