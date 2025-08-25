# Create IAM role for Lex bot
resource "aws_iam_role" "lex_bot_role" {
  name = "lex-bot-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "lexv2.amazonaws.com"
        },
        Effect = "Allow",
        Sid    = ""
      },
    ]
  })
}

# Attach the required policy to the Lex bot role
resource "aws_iam_role_policy_attachment" "lex_bot_policy" {
  role       = aws_iam_role.lex_bot_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonLexFullAccess"
}

# Create example bot user role
resource "aws_iam_role" "bot_user" {
  name = "example-bot-user"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Effect = "Allow",
        Sid    = ""
      },
    ]
  })
}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create Lex Bot using AWS provider
resource "aws_lexv2models_bot" "example" {
  name     = "ExampleBot"
  role_arn = aws_iam_role.lex_bot_role.arn

  data_privacy {
    child_directed = false
  }

  idle_session_ttl_in_seconds = 300
}

# Create a bot locale
resource "aws_lexv2models_bot_locale" "example" {
  bot_id      = aws_lexv2models_bot.example.id
  locale_id   = "en_US"
  description = "English US locale for Example Bot"

  n_lu_intent_confidence_threshold = 0.4

  voice_settings {
    voice_id = "Salli"
  }

  depends_on = [aws_lexv2models_bot.example]
}

# Create a welcome intent
resource "aws_lexv2models_intent" "welcome" {
  bot_id    = aws_lexv2models_bot.example.id
  locale_id = aws_lexv2models_bot_locale.example.locale_id

  name        = "WelcomeIntent"
  description = "Intent to welcome users"

  sample_utterance {
    utterance = "hello"
  }

  sample_utterance {
    utterance = "hi"
  }

  depends_on = [aws_lexv2models_bot_locale.example]
}

# Create a bot version
resource "aws_lexv2models_bot_version" "example" {
  bot_id = aws_lexv2models_bot.example.id

  bot_version_locale_specification = {
    "en_US" = {
      source_bot_version = "DRAFT"
    }
  }

  depends_on = [aws_lexv2models_intent.welcome]
}

# Create a bot alias using awscc provider
resource "awscc_lex_bot_alias" "example" {
  bot_alias_name = "ExampleAlias"
  bot_id         = aws_lexv2models_bot.example.id
  bot_version    = aws_lexv2models_bot_version.example.bot_version
}

# Create a Lex resource policy using awscc provider
resource "awscc_lex_resource_policy" "example" {
  resource_arn = "arn:aws:lex:us-west-2:${data.aws_caller_identity.current.account_id}:bot-alias/${awscc_lex_bot_alias.example.bot_id}/${awscc_lex_bot_alias.example.bot_alias_id}"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "ExampleBotAliasPermission",
        Effect = "Allow",
        Principal = {
          AWS = aws_iam_role.bot_user.arn
        },
        Action = [
          "lex:RecognizeText",
          "lex:RecognizeUtterance",
          "lex:StartConversation"
        ],
        Resource = [
          "arn:aws:lex:us-west-2:${data.aws_caller_identity.current.account_id}:bot-alias/${awscc_lex_bot_alias.example.bot_id}/${awscc_lex_bot_alias.example.bot_alias_id}"
        ]
      }
    ]
  })

  depends_on = [awscc_lex_bot_alias.example]
}
