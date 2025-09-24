resource "awscc_gamelift_alias" "example" {
  name        = "simple"
  description = "alias uses SIMPLE routing strategy"

  routing_strategy = {
    type     = "SIMPLE"
    fleet_id = awscc_gamelift_fleet.example.fleet_id
  }
}