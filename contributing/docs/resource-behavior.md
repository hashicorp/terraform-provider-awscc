# Terraform AWS Cloud Control Provider Resource Behavior

This document describes the behavior of resources implemented by the provider. In particular it describes in detail how [AWS CloudFormation resource types](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html) are defined as [Terraform resources](https://developer.hashicorp.com/terraform/language/resources).

We assume some familiarity with user-facing Terraform concepts like configuration, state, CLI workflow, etc. The Terraform website has documentation on these ideas.

## Resources

Resources describe infrastructure objects such as virtual networks or compute instances.