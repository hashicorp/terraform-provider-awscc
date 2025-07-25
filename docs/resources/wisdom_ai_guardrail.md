
---
page_title: "awscc_wisdom_ai_guardrail Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Wisdom::AIGuardrail Resource Type
---

# awscc_wisdom_ai_guardrail (Resource)

Definition of AWS::Wisdom::AIGuardrail Resource Type

## Example Usage

### Comprehensive Wisdom AI Guardrail Configuration

Creates a Wisdom AI Guardrail with extensive content filtering policies including hate speech detection, word restrictions, topic controls, PII protection, and contextual grounding configurations, all integrated with a Wisdom Assistant.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a Wisdom Assistant first (required for AI Guardrail)
resource "awscc_wisdom_assistant" "example" {
  name        = "example-assistant-${formatdate("YYYYMMDD-hhmmss", timestamp())}"
  type        = "AGENT"
  description = "Example Wisdom Assistant for AI Guardrail"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the AI Guardrail
resource "awscc_wisdom_ai_guardrail" "example" {
  name         = "example-guardrail-${formatdate("YYYYMMDD-hhmmss", timestamp())}"
  description  = "Example AI Guardrail configuration"
  assistant_id = awscc_wisdom_assistant.example.id

  # Required messaging configurations
  blocked_input_messaging   = "Your input contains prohibited content. Please revise and try again."
  blocked_outputs_messaging = "The response was blocked due to policy violations."

  # Content policy configuration
  content_policy_config = {
    filters_config = [
      {
        type            = "HATE"
        input_strength  = "HIGH"
        output_strength = "HIGH"
      }
    ]
  }

  # Word policy configuration
  word_policy_config = {
    words_config = [
      {
        text = "restricted_word"
      }
    ]
    managed_word_lists_config = [
      {
        type = "PROFANITY"
      }
    ]
  }

  # Topic policy configuration
  topic_policy_config = {
    topics_config = [
      {
        name       = "Restricted Topic"
        type       = "DENY"
        definition = "Any discussion related to restricted topics"
        examples   = ["This is an example of restricted content"]
      }
    ]
  }

  # Sensitive information policy configuration
  sensitive_information_policy_config = {
    pii_entities_config = [
      {
        type   = "US_BANK_ACCOUNT_NUMBER"
        action = "BLOCK"
      }
    ]
    regexes_config = [
      {
        name        = "CustomPattern"
        description = "Custom pattern to detect sensitive information"
        pattern     = "[A-Z]{2}\\d{6}"
        action      = "BLOCK"
      }
    ]
  }

  # Contextual grounding policy configuration
  contextual_grounding_policy_config = {
    filters_config = [
      {
        type      = "GROUNDING"
        threshold = 0.7
      }
    ]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `assistant_id` (String)
- `blocked_input_messaging` (String) Messaging for when violations are detected in text
- `blocked_outputs_messaging` (String) Messaging for when violations are detected in text

### Optional

- `content_policy_config` (Attributes) Content policy config for a guardrail. (see [below for nested schema](#nestedatt--content_policy_config))
- `contextual_grounding_policy_config` (Attributes) Contextual grounding policy config for a guardrail. (see [below for nested schema](#nestedatt--contextual_grounding_policy_config))
- `description` (String) Description of the guardrail or its version
- `name` (String)
- `sensitive_information_policy_config` (Attributes) Sensitive information policy config for a guardrail. (see [below for nested schema](#nestedatt--sensitive_information_policy_config))
- `tags` (Map of String)
- `topic_policy_config` (Attributes) Topic policy config for a guardrail. (see [below for nested schema](#nestedatt--topic_policy_config))
- `word_policy_config` (Attributes) Word policy config for a guardrail. (see [below for nested schema](#nestedatt--word_policy_config))

### Read-Only

- `ai_guardrail_arn` (String)
- `ai_guardrail_id` (String)
- `assistant_arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--content_policy_config"></a>
### Nested Schema for `content_policy_config`

Optional:

- `filters_config` (Attributes List) List of content filter configs in content policy. (see [below for nested schema](#nestedatt--content_policy_config--filters_config))

<a id="nestedatt--content_policy_config--filters_config"></a>
### Nested Schema for `content_policy_config.filters_config`

Optional:

- `input_strength` (String) Strength for filters
- `output_strength` (String) Strength for filters
- `type` (String) Type of text to text filter in content policy



<a id="nestedatt--contextual_grounding_policy_config"></a>
### Nested Schema for `contextual_grounding_policy_config`

Optional:

- `filters_config` (Attributes List) List of contextual grounding filter configs. (see [below for nested schema](#nestedatt--contextual_grounding_policy_config--filters_config))

<a id="nestedatt--contextual_grounding_policy_config--filters_config"></a>
### Nested Schema for `contextual_grounding_policy_config.filters_config`

Optional:

- `threshold` (Number) The threshold for this filter.
- `type` (String) Type of contextual grounding filter



<a id="nestedatt--sensitive_information_policy_config"></a>
### Nested Schema for `sensitive_information_policy_config`

Optional:

- `pii_entities_config` (Attributes List) List of entities. (see [below for nested schema](#nestedatt--sensitive_information_policy_config--pii_entities_config))
- `regexes_config` (Attributes List) List of regex. (see [below for nested schema](#nestedatt--sensitive_information_policy_config--regexes_config))

<a id="nestedatt--sensitive_information_policy_config--pii_entities_config"></a>
### Nested Schema for `sensitive_information_policy_config.pii_entities_config`

Optional:

- `action` (String) Options for sensitive information action.
- `type` (String) The currently supported PII entities


<a id="nestedatt--sensitive_information_policy_config--regexes_config"></a>
### Nested Schema for `sensitive_information_policy_config.regexes_config`

Optional:

- `action` (String) Options for sensitive information action.
- `description` (String) The regex description.
- `name` (String) The regex name.
- `pattern` (String) The regex pattern.



<a id="nestedatt--topic_policy_config"></a>
### Nested Schema for `topic_policy_config`

Optional:

- `topics_config` (Attributes List) List of topic configs in topic policy. (see [below for nested schema](#nestedatt--topic_policy_config--topics_config))

<a id="nestedatt--topic_policy_config--topics_config"></a>
### Nested Schema for `topic_policy_config.topics_config`

Optional:

- `definition` (String) Definition of topic in topic policy
- `examples` (List of String) List of text examples
- `name` (String) Name of topic in topic policy
- `type` (String) Type of topic in a policy



<a id="nestedatt--word_policy_config"></a>
### Nested Schema for `word_policy_config`

Optional:

- `managed_word_lists_config` (Attributes List) A config for the list of managed words. (see [below for nested schema](#nestedatt--word_policy_config--managed_word_lists_config))
- `words_config` (Attributes List) List of custom word configs. (see [below for nested schema](#nestedatt--word_policy_config--words_config))

<a id="nestedatt--word_policy_config--managed_word_lists_config"></a>
### Nested Schema for `word_policy_config.managed_word_lists_config`

Optional:

- `type` (String) Options for managed words.


<a id="nestedatt--word_policy_config--words_config"></a>
### Nested Schema for `word_policy_config.words_config`

Optional:

- `text` (String) The custom word text.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_wisdom_ai_guardrail.example
  id = "ai_guardrail_id|assistant_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_wisdom_ai_guardrail.example "ai_guardrail_id|assistant_id"
```
