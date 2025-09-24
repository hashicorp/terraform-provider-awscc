resource "awscc_gamelift_alias" "example" {
  name        = "terminal"
  description = "alias uses TERMINAL routing strategy"

  routing_strategy = {
    type    = "TERMINAL"
    message = "Terminal routing strategy message"
  }
}