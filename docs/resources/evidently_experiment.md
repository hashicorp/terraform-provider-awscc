
---
page_title: "awscc_evidently_experiment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Evidently::Experiment.
---

# awscc_evidently_experiment (Resource)

Resource Type definition for AWS::Evidently::Experiment.

## Example Usage

### Configure A/B Testing Experiment with Amazon CloudWatch Evidently

Creates an Evidently experiment for A/B testing with a 60/40 split between control and treatment groups, measuring page load time metrics with custom sampling rate and automatic experiment start.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create an Evidently Project first since it's required for the experiment
resource "awscc_evidently_project" "example" {
  name        = "example-project"
  description = "Example project for Evidently experiment"
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create the Evidently Experiment
resource "awscc_evidently_experiment" "example" {
  name        = "example-experiment"
  description = "Example experiment using AWSCC provider"
  project     = awscc_evidently_project.example.name

  metric_goals = [
    {
      desired_change = "INCREASE"
      entity_id_key  = "userDetails.userId"
      metric_name    = "page_load_time"
      value_key      = "details.timeInMs"
      event_pattern = jsonencode({
        detail-type : ["page-load"]
        source : ["com.example.web"]
      })
      unit_label = "milliseconds"
    }
  ]

  online_ab_config = {
    control_treatment_name = "control"
    treatment_weights = [
      {
        treatment    = "treatment-A"
        split_weight = 40
      },
      {
        treatment    = "control"
        split_weight = 60
      }
    ]
  }

  treatments = [
    {
      feature        = "page-rendering"
      treatment_name = "control"
      description    = "Current page rendering logic"
      variation      = "original"
    },
    {
      feature        = "page-rendering"
      treatment_name = "treatment-A"
      description    = "New optimized page rendering"
      variation      = "optimized"
    }
  ]

  sampling_rate = 10

  running_status = {
    status = "START"
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metric_goals` (Attributes List) (see [below for nested schema](#nestedatt--metric_goals))
- `name` (String)
- `online_ab_config` (Attributes) (see [below for nested schema](#nestedatt--online_ab_config))
- `project` (String)
- `treatments` (Attributes List) (see [below for nested schema](#nestedatt--treatments))

### Optional

- `description` (String)
- `randomization_salt` (String)
- `remove_segment` (Boolean)
- `running_status` (Attributes) Start Experiment. Default is False (see [below for nested schema](#nestedatt--running_status))
- `sampling_rate` (Number)
- `segment` (String)
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--metric_goals"></a>
### Nested Schema for `metric_goals`

Required:

- `desired_change` (String)
- `entity_id_key` (String) The JSON path to reference the entity id in the event.
- `metric_name` (String)
- `value_key` (String) The JSON path to reference the numerical metric value in the event.

Optional:

- `event_pattern` (String) Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.
- `unit_label` (String)


<a id="nestedatt--online_ab_config"></a>
### Nested Schema for `online_ab_config`

Optional:

- `control_treatment_name` (String)
- `treatment_weights` (Attributes Set) (see [below for nested schema](#nestedatt--online_ab_config--treatment_weights))

<a id="nestedatt--online_ab_config--treatment_weights"></a>
### Nested Schema for `online_ab_config.treatment_weights`

Optional:

- `split_weight` (Number)
- `treatment` (String)



<a id="nestedatt--treatments"></a>
### Nested Schema for `treatments`

Required:

- `feature` (String)
- `treatment_name` (String)
- `variation` (String)

Optional:

- `description` (String)


<a id="nestedatt--running_status"></a>
### Nested Schema for `running_status`

Optional:

- `analysis_complete_time` (String) Provide the analysis Completion time for an experiment
- `desired_state` (String) Provide CANCELLED or COMPLETED desired state when stopping an experiment
- `reason` (String) Reason is a required input for stopping the experiment
- `status` (String) Provide START or STOP action to apply on an experiment


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_evidently_experiment.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_evidently_experiment.example "arn"
```
