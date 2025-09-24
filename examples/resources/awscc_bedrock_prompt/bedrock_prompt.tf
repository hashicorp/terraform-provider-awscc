resource "awscc_bedrock_prompt" "example" {
  name                        = "example"
  description                 = "example"
  customer_encryption_key_arn = awscc_kms_key.example.arn
  default_variant             = "variant-example"

  variants = [
    {
      name          = "variant-example"
      template_type = "TEXT"
      model_id      = "amazon.titan-text-express-v1"
      inference_configuration = {
        text = {
          temperature    = 1
          top_p          = 0.9900000095367432
          max_tokens     = 300
          stop_sequences = ["\\n\\nHuman:"]
          top_k          = 250
        }
      }
      template_configuration = {
        text = {
          input_variables = [
            {
              name = "topic"
            }
          ]
          text = "Make me a {{genre}} playlist consisting of the following number of songs: {{number}}."
        }
      }
    }

  ]

}
