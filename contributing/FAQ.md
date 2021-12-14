# Frequently Asked Questions

<!-- markdownlint-disable no-trailing-punctuation -->
<!-- markdownlint-disable-next-line heading-increment -->
### Who are the maintainers?

The HashiCorp Terraform AWS provider team is:

* Mary Cutrali, Product Manager - GitHub [@maryelizbeth](https://github.com/maryelizbeth), Twitter [@marycutrali](https://twitter.com/marycutrali)
* Kit Ewbank, Engineer - GitHub [@ewbankkit](https://github.com/ewbankkit)
* Graham Davison, Engineer - GitHub [@gdavison](https://github.com/gdavison)
* Angie Pinilla, Engineer - GitHub [@angie44](https://github.com/angie44)
* Dirk Avery (Federal), Engineer - GitHub [@YakDriver](https://github.com/yakdriver)
* Zoe Helding, Engineer - GitHub [@zhelding](https://github.com/zhelding)
* Simon Davis, Engineering Manager - GitHub [@breathingdust](https://github.com/breathingdust)
* Kerim Satirli, Developer Advocate - GitHub [@ksatirli](https://github.com/ksatirli)
* Justin Retzolk, Technical Community Manager - GitHub [@justinretzolk](https://github.com/justinretzolk)

### How is this provider different from the existing AWS provider?

The Cloud Control provider is wholly generated from AWS CloudFormation schema. This means no manual work is required to add new features or services to the provider, and practitioners can expect to use AWS services within days of the AWS's launch rather than potentially wait for community support to prioritize that feature for inclusion. The standard [AWS Provider](https://github.com/hashicorp/terraform-provider-aws) is manually written and is the product of thousands of contributors and years of work.

### CloudFormation seems to support the resource I want to use, but I don’t see it in the provider?

There are a few reasons why this might happen:

#### Launch Schedule

At present, we plan to release the provider weekly, rolling up any additions from AWS collecting over the preceding week. This means you could wait up to a week to see the service in the provider. We do plan to narrow this gap as the provider reaches GA status.

#### CloudFormation Schema Version

The Cloud Control API is only compatible with the latest version of the CloudFormation schema. While many CloudFormation resource schemas are now updated to use this latest version (and all new services will use it at launch), there will be a period where some AWS services still use the prior version. AWS are updating all services using the older schema to use the newer one. To see which services are supported you can refer to the provider documentation or use the following AWS CLI command:

```console
$ aws cloudformation list-types --type RESOURCE --visibility PUBLIC --provisioning-type FULLY_MUTABLE --filters Category=AWS_TYPES
```

This command lists all AWS resources that are usable with Cloud Control.

#### Bug in the Provider

It's possible that as AWS release new services, the code which generates the provider may contain bugs which prevent a specific resource to be added. It's also possible, that a CloudFormation resource schema exhibits a configuration that we haven't encountered before and requires special handling. We will exclude those resources until we are able to resolve those issues.

### What is your release schedule?

We release weekly on Thursday rolling up the preceding week's CloudFormation additions/enhancements. For this reason there may be a delay of up to a week before a feature available in AWS can be usable in the provider. We do plan to narrow this gap as the provider reaches GA status.

### CloudFormation doesn't support the service or resource I want to use, how do I request coverage?

AWS are aiming for 100% coverage of their AWS service area in CloudFormation. New services will nearly always launch with CloudFormation support, and older services will be updated in time. Please refer to the [CloudFormation](https://github.com/aws-cloudformation/cloudformation-coverage-roadmap/projects/1) Open Coverage roadmap for more details or to upvote the service you would like to see coverage for.

### The CloudFormation schema doesn't cover all API actions, and I can’t do what I used to do with the classic provider?

CloudFormation takes a different approach to the existing Terraform AWS Provider and typically does not model more operational concerns in its schema focusing exclusively on configuration. At this time we do not intend to customize the generated resources to add this functionality, but this is under active consideration. We expect this to be a rare occurrence, but do want to hear from you if you feel this is impeding your usage.

### Which provider should I use?

While this provider is under Technical Preview, we do not recommend using it for your production setup. While we expect the functionality to be stable, it's possible the interface will not be. As a result, if you choose to use it, please expect and plan for changes.

When the provider is announced as Generally Available the choice of provider will depend on which resources you need to configure. Until CloudFormation achieves 100% coverage there is likely to be gaps in both providers, so in the case where a single provider can’t manage your infrastructure we would recommend using both. Once this provider is GA, we recommend using it for any resource that is supported including ones that may be available in the “classic” provider.

### Is there a way to migrate my existing resources to the new provider?

Not at this time. There are no plans to deprecate the existing provider at this time, and we plan to fully support it for the foreseeable future. We are actively researching possibilities to ease migration.

### How can I help?

Contributing to this provider will be very different from the prior AWS provider but we will accept PR’s to help with any aspects of code generation, runtime and documentation. We will not accept any PR’s which alter or modify the behavior of generated resources at this time, but would love to hear where you think this is necessary.
