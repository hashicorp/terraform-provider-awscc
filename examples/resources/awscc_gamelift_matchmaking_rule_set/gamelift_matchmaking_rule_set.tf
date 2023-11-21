resource "awscc_gamelift_matchmaking_rule_set" "example" {
  name = "ExampleRuleSet"
  rule_set_body = jsonencode({
    name                = "ExampleRuleSet",
    ruleLanguageVersion = "1.0",
    playerAttributes = [
      {
        name    = "skill",
        type    = "number",
        default = 10
      }
    ],
    teams = [
      {
        name       = "all",
        minPlayers = 3,
        maxPlayers = 5
      }
    ],
    rules = [
      {
        name        = "FairTeamSkill",
        description = "The average skill of players in each team is within 10 points from the average skill of players in the match",
        type        = "distance",
        // get players for each team, and average separately to produce list of 3
        measurements = ["avg(teams[*].players.attributes[skill])"],
        // get players for each team, flatten into a single list, and average to produce overall average
        referenceValue = "avg(flatten(teams[*].players.attributes[skill]))",
        maxDistance    = 10 // minDistance would achieve the opposite result
      }
    ],
    expansions = [{
      target = "rules[FairTeamSkill].maxDistance",
      steps = [
        {
          waitTimeSeconds = 20,
          value           = 100
        },
        {
          waitTimeSeconds = 30,
          value           = 150
        }
      ]
    }]
  })
}
