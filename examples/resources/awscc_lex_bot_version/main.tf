# This example demonstrates creating a Lex Bot Version using AWSCC provider.
# Note: This requires a pre-existing Lex V2 bot.

# Create a bot version using AWSCC provider
resource "awscc_lex_bot_version" "example" {
  # You would need to replace this with a real Lex V2 bot ID
  bot_id      = "ABC123DEF4"
  description = "Example bot version configuration"

  bot_version_locale_specification = [
    {
      locale_id = "en_US"
      bot_version_locale_details = {
        source_bot_version = "DRAFT"
      }
    }
  ]
}