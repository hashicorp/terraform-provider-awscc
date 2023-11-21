resource "awscc_gamelift_matchmaking_rule_set" "example" {
  name             = "ExampleRuleSet"
  rule_set_body = <<-EOF
  {
    "name": "ExampleRuleSet",
    "ruleLanguageVersion": "1.0",
    "playerAttributes": [{
        "name": "skill",
        "type": "number",
        "default": 10
    }],
    "teams": [{
        "name": "all",
        "minPlayers": 3,
        "maxPlayers": 5
    }],
    "rules": [{
        "name": "FairTeamSkill",
        "description": "The average skill of players in each team is within 10 points from the average skill of players in the match",
        "type": "distance",
        "measurements": [ "avg(teams[*].players.attributes[skill])" ],
        "referenceValue": "avg(flatten(teams[*].players.attributes[skill]))",
        "maxDistance": 10
    }],
    "expansions": [{
        "target": "rules[FairTeamSkill].maxDistance",
        "steps": [{
            "waitTimeSeconds": 20,
            "value": 100
        }, {
            "waitTimeSeconds": 30,
            "value": 150
        }]
    }]
}
  EOF
}

resource "awscc_gamelift_game_session_queue" "example" {
  name             = "ExampleQueue"
  timeout_in_seconds = 600

  destinations = [
    {
      destination_arn = "your-fleet-or-fleet-alias-arn" // ARN of your Fleet or Fleet Alias
    },
  ]

}

resource "awscc_gamelift_matchmaking_configuration" "example" {
  name             = "ExampleRuleSet"
  acceptance_required = false
  request_timeout_seconds = 60
  rule_set_name = awscc_gamelift_matchmaking_rule_set.example.id

  flex_match_mode = "WITH_QUEUE" // valid values are "WITH_QUEUE" or "STANDALONE"
  game_session_queue_arns = [
    awscc_gamelift_game_session_queue.example.arn,
  ]

}
