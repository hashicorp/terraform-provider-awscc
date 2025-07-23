resource "awscc_gamelift_game_session_queue" "example" {
  name               = "ExampleQueue"
  timeout_in_seconds = 600

  destinations = [
    {
      destination_arn = "your-fleet-or-fleet-alias-arn" // ARN of your Fleet or Fleet Alias
    },
  ]

}
