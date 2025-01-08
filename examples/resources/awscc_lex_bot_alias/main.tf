data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example configuration for an Amazon Lex bot alias 
# Please replace XXXXXXXXXXX with your bot ID
resource "awscc_lex_bot_alias" "example" {
  bot_alias_name = "TestBotAlias"
  bot_id         = "XXXXXXXXXX" # Replace with a valid bot ID
  bot_version    = "DRAFT"
  description    = "Test bot alias for example"

  bot_alias_locale_settings = [{
    locale_id = "en_US"
    bot_alias_locale_setting = {
      enabled = true
      code_hook_specification = {
        lambda_code_hook = {
          lambda_arn                  = "arn:aws:lambda:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:function:example-function"
          code_hook_interface_version = "1.0"
        }
      }
    }
  }]

  sentiment_analysis_settings = {
    detect_sentiment = true
  }

  conversation_log_settings = {
    text_log_settings = [{
      enabled = true
      destination = {
        cloudwatch = {
          cloudwatch_log_group_arn = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lex/TestBot:*",
          log_prefix               = "TestBot/"
        }
      }
    }]
  }

  bot_alias_tags = [{
    key   = "Environment"
    value = "Test"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
  }]
}