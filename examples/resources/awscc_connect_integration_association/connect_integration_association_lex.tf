resource "awscc_connect_integration_association" "this" {
  instance_id      = "arn:aws:connect:us-east-1:111122223333:instance/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
  integration_arn  = "arn:aws:lex:us-east-1:111122223333:bot-alias/AAAAAAAAAA/BBBBBBBBBB"
  integration_type = "LEX_BOT"
}