<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-terraform-main.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform AWS Cloud Control Provider

*This provider is currently in technical preview. This means some aspects of its design and implementation are not yet considered stable. We are actively looking for community feedback in order to solidify its form.*

- Announcement: [HashiCorp Blog](https://www.hashicorp.com/blog/announcing-terraform-aws-cloud-control-provider-tech-preview)
- Terraform Website: [terraform.io](https://terraform.io)
- Provider Documentation: [Terraform Registry](https://registry.terraform.io/providers/hashicorp/awscc/latest)
- Forum: [discuss.hashicorp.com](https://discuss.hashicorp.com/c/terraform-providers/tf-awscc/)
- Tutorial: [learn.hashicorp.com](https://learn.hashicorp.com/tutorials/terraform/aws-cloud-control)

The Terraform AWS Cloud Control Provider is a plugin for Terraform that allows for the full lifecycle management of AWS resources using the AWS CloudFormation Cloud Control API.
This provider is maintained internally by the HashiCorp AWS Provider team.

### AWS Cloud Control API

The [AWS Cloud Control API](https://docs.aws.amazon.com/cloudcontrolapi) is a lightweight proxy API to discover, provision and manage cloud resources through a simple, uniform and predictable control plane.
The AWS Cloud Control API supports **C**reate, **R**ead, **U**pdate, **D**elete and **L**ist (CRUDL) operations on any AWS resource that is registered in the [AWS CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html).

#### Coverage

At launch a subset of AWS resources which can be managed by CloudFormation are supported, some services use an older CloudFormation schema and cannot be used with Cloud Control. AWS are updating all of the older CloudFormation schemas to conform to the new standard, and are actively pursuing full coverage for CloudFormation. For the latest coverage information please refer to the AWS CloudFormation public [roadmap](https://github.com/aws-cloudformation/cloudformation-coverage-roadmap/projects/1).

To see the list of supported resources within this provider please refer to the registry.

### Release Schedule

This provider is generated from the latest CloudFormation schemas, and will release weekly containing all new services and enhancements added to Cloud Control.

### Credentials

When performing CRUDL operations the Cloud Control API make calls to downstream AWS services on your behalf. By default, the Cloud Control API will create a temporary session using the AWS credentials of the user making the Cloud Control API call. This session lasts up to a maximum of 24 hours.

All CRUDL operations also accept a `RoleArn` parameter which represents the [AWS CloudFormation service role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-servicerole.html). In addition to federating access, using a role allows you to extend the allowed time of an operation to 36 hours, as the Cloud Control API can refresh the role credentials by re-assuming the role. The Terraform AWS Cloud Control API Provider has a `role_arn` argument which enables support for this functionality.

**Please note:** We take Terraform's security and our users' trust very seriously. If you believe you have found a security issue in the Terraform AWS Cloud Control Provider, please responsibly disclose by contacting us at security@hashicorp.com.

## Quick Starts

- [Using the Provider](https://learn.hashicorp.com/tutorials/terraform/aws-cloud-control)
- [Provider development](contributing/DEVELOPMENT.md)

## Documentation

Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/hashicorp/awscc/latest)

## Frequently Asked Questions

Responses to our most frequently asked questions can be found in our [FAQ](contributing/FAQ.md )

## Contributing

The Terraform Provider for AWS CloudFormation Cloud Control API is the work of a handful of contributors. We appreciate your help!

To contribute, please read the contribution guidelines: [Contributing to Terraform - AWS Cloud Control Provider](contributing/CONTRIBUTING.md)
