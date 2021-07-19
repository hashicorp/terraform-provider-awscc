<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform AWS CloudAPI Provider

- Website: [terraform.io](https://terraform.io)
- Tutorials: [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
- Forum: [discuss.hashicorp.com](https://discuss.hashicorp.com/c/terraform-providers/tf-aws/)
- Chat: [gitter](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing List: [Google Groups](http://groups.google.com/group/terraform-tool)

The Terraform AWS CloudAPI Provider is a plugin for Terraform that allows for the full lifecycle management of AWS resources using the AWS CloudFormation Cloud API.
This provider is maintained internally by the HashiCorp AWS Provider team.

### AWS CloudFormation Cloud API

The Cloud API is a lighweight proxy API to discover, provision and manage cloud resources through a simple, uniform and predictable control plane.
The Cloud API supports **C**reate, **R**ead, **U**pdate, **D**elete and **L**ist (CRUDL) operations on any resource that is registered in the [AWS CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html).

### Credentials

When performing CRUDL operations the Cloud API make calls to downstream AWS services on your behalf. By default, the Cloud API will create a temporary session using the AWS credentials of the user making the Cloud API call. This session lasts up to a maximum of 24 hours.

All CRUDL operations also accept a `RoleArn` parameter which represents the [AWS CloudFormation service role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-servicerole.html). In addition to federating access, using a role allows you to extend the allowed time of an operation to 36 hours, as the Cloud API can refresh the role credentials by re-assuming the role. The Terraform AWS Cloud API Provider has a `role_arn` argument which enables support for this functionality.

**Please note:** We take Terraform's security and our users' trust very seriously. If you believe you have found a security issue in the Terraform AWS CloudAPI Provider, please responsibly disclose by contacting us at security@hashicorp.com.

## Quick Starts

- Using the provider
- [Provider development](docs/DEVELOPMENT.md)

## Documentation

Full, comprehensive documentation is available on the Terraform website:

## Roadmap

Our roadmap for expanding support in Terraform for AWS resources can be found in our [Roadmap](ROADMAP.md) which is published quarterly.

## Frequently Asked Questions

Responses to our most frequently asked questions can be found in our [FAQ](docs/FAQ.md )

## Contributing

The Terraform Provider for AWS CloudFormation Cloud API is the work of a handful of contributors. We appreciate your help!

To contribute, please read the contribution guidelines: [Contributing to Terraform - AWS CloudAPI Provider](docs/CONTRIBUTING.md)
