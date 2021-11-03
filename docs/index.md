---
layout: ""
page_title: "Provider: AWS Cloud Control"
description: |-
  Use the Amazon Web Services (AWS) Cloud Control provider to interact with the many resources supported by AWS via the Cloud Control API.
---

# AWS Cloud Control Provider

Use the Amazon Web Services (AWS) Cloud Control provider to interact with the many resources supported by AWS via the Cloud Control API. You must configure the provider with the proper credentials before you can use it.

Use the navigation to the left to read about the available resources.

To learn the basics of Terraform using this provider, follow the hands-on [get started tutorials](https://learn.hashicorp.com/tutorials/terraform/aws-cloud-control?in=terraform/aws) on HashiCorp's Learn platform.

## Example Usage

Terraform 1.0.7 and later:

```terraform
terraform {
  required_providers {
    awscc = {
      source  = "hashicorp/awscc"
      version = "~> 0.1"
    }
  }
}

# Configure the AWS Provider
provider "awscc" {
  region = "us-west-2"
}

# Create a Log Group
resource "awscc_logs_log_group" "example" {
  log_group_name = "example"
}
```

## Authentication

The AWS provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables
- Shared credentials/configuration file
- CodeBuild, ECS, and EKS Roles
- EC2 Instance Metadata Service (IMDS and IMDSv2)

### Static Credentials

!> **Warning:** Hard-coded credentials are not recommended in any Terraform
configuration and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `access_key` and `secret_key`
in-line in the AWS provider block:

Usage:

```terraform
provider "awscc" {
  region     = "us-west-2"
  access_key = "AKIAIOSFODNN7EXAMPLE"
  secret_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
}
```

### Environment Variables

You can provide your credentials via the `AWS_ACCESS_KEY_ID` and
`AWS_SECRET_ACCESS_KEY`, environment variables, representing your AWS
Access Key and AWS Secret Key, respectively.  Note that setting your
AWS credentials using either these (or legacy) environment variables
will override the use of `AWS_SHARED_CREDENTIALS_FILE` and `AWS_PROFILE`.
The `AWS_DEFAULT_REGION` and `AWS_SESSION_TOKEN` environment variables
are also used, if applicable:

```terraform
provider "awscc" {}
```

Usage:

```sh
$ export AWS_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE"
$ export AWS_SECRET_ACCESS_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
$ export AWS_DEFAULT_REGION="us-west-2"
$ terraform plan
```

### Shared Configuration and Credentials Files

You can use [AWS shared configuration or credentials files](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) to specify configuration and credentials.
The default locations are `$HOME/.aws/config` and `$HOME/.aws/credentials` on Linux and macOS,
or `"%USERPROFILE%\.aws\config"` and `"%USERPROFILE%\.aws\credentials"` on Windows.
You can optionally specify a different location in the Terraform configuration by providing the `shared_config_files` and `shared_credentials_files` arguments
or using the `AWS_SHARED_CONFIG_FILE` and `AWS_SHARED_CREDENTIALS_FILE` environment variables.
This method also supports the `profile` configuration and corresponding `AWS_PROFILE` environment variable:

Usage:

```terraform
provider "awscc" {
  region                   = "us-west-2"
  shared_config_files      = ["/Users/tf_user/.aws/conf"]
  shared_credentials_files = ["/Users/tf_user/.aws/creds"]
  profile                  = "customprofile"
}
```

Please note that the [AWS Go SDK](https://aws.amazon.com/sdk-for-go/), the underlying authentication handler used by the Terraform AWS Provider, does not support all AWS CLI features.

### CodeBuild, ECS, and EKS Roles

If you're running Terraform on CodeBuild or ECS and have configured an [IAM Task Role](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html), Terraform will use the container's Task Role. This support is based on the underlying `AWS_CONTAINER_CREDENTIALS_RELATIVE_URI` and `AWS_CONTAINER_CREDENTIALS_FULL_URI` environment variables being automatically set by those services or manually for advanced usage.

If you're running Terraform on EKS and have configured [IAM Roles for Service Accounts (IRSA)](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html), Terraform will use the pod's role. This support is based on the underlying `AWS_ROLE_ARN` and `AWS_WEB_IDENTITY_TOKEN_FILE` environment variables being automatically set by Kubernetes or manually for advanced usage.

### Custom User-Agent Information

By default, the underlying AWS client used by the Terraform AWS Provider creates requests with User-Agent headers including information about Terraform and AWS Go SDK versions.
To provide additional information in the User-Agent headers, set the User-Agent product or comment information using the `user_agent` argument.
For example,

```terraform
provider "awscc" {
  user_agent = [
    {
      product_name    = "example-module"
      product_version = "1.0"
    }
    {
      product_name    = "BuildID"
      product_version = "1234"
    }
  ]
}
```

will append `example-module/1.0 BuildID/1234` to the User-Agent.

In addition, the `TF_APPEND_USER_AGENT` environment variable can be set and its value will be directly added to HTTP requests. e.g.

```sh
$ export TF_APPEND_USER_AGENT="JenkinsAgent/i-12345678 BuildID/1234 (Optional Extra Information)"
```

### EC2 Instance Metadata Service

If you're running Terraform from an EC2 instance with IAM Instance Profile
using IAM Role, Terraform will just ask
[the metadata API](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html#instance-metadata-security-credentials)
endpoint for credentials.

This is a preferred approach over any other when running in EC2 as you can avoid
hard coding credentials. Instead these are leased on-the-fly by Terraform
which reduces the chance of leakage.

You can provide the custom metadata API endpoint via the `AWS_METADATA_URL` variable
which expects the endpoint URL, including the version, and defaults to `http://169.254.169.254:80/latest`.

### Assume Role

If provided with a role ARN, Terraform will attempt to assume this role
using the supplied credentials.

Usage:

```terraform
provider "awscc" {
  assume_role = {
    role_arn     = "arn:aws:iam::ACCOUNT_ID:role/ROLE_NAME"
    session_name = "SESSION_NAME"
    external_id  = "EXTERNAL_ID"
  }
}
```

> **Hands-on:** Try the [Use AssumeRole to Provision AWS Resources Across Accounts](https://learn.hashicorp.com/tutorials/terraform/aws-assumerole) tutorial on HashiCorp Learn.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **access_key** (String) This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment variable, or via a shared credentials file if `profile` is specified.
- **assume_role** (Attributes) An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration. (see [below for nested schema](#nestedatt--assume_role))
- **http_proxy** (String) The address of an HTTP proxy to use when accessing the AWS API. Can also be configured using the `HTTP_PROXY` or `HTTPS_PROXY` environment variables.
- **insecure** (Boolean) Explicitly allow the provider to perform "insecure" SSL requests. If not set, defaults to `false`.
- **max_retries** (Number) The maximum number of times an AWS API request is retried on failure. If not set, defaults to 25.
- **profile** (String) This is the AWS profile name as set in the shared credentials file.
- **region** (String) This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment variables, or via a shared config file.
- **role_arn** (String) Amazon Resource Name of the AWS CloudFormation service role that is used on your behalf to perform operations.
- **secret_key** (String) This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment variable, or via a shared credentials file if `profile` is specified.
- **shared_config_files** (List of String) List of paths to shared config files. If not set, defaults to `~/.aws/config`.
- **shared_credentials_files** (List of String) List of paths to shared credentials files. If not set, defaults to `~/.aws/credentials`.
- **skip_medatadata_api_check** (Boolean) Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint.  Setting to `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
- **token** (String) Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.  It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
- **user_agent** (Attributes List) Product details to append to User-Agent string in all AWS API calls. (see [below for nested schema](#nestedatt--user_agent))

<a id="nestedatt--assume_role"></a>
### Nested Schema for `assume_role`

Optional:

- **duration** (String) Duration of the assume role session. You can provide a value from 15 minutes up to the maximum session duration setting for the role. A sequence of numbers with a unit suffix, "h" for hour, "m" for minute, and "s" for second. Default value is 1h0m0s
- **external_id** (String) External identifier to use when assuming the role.
- **policy** (String) IAM policy in JSON format to use as a session policy. The effective permissions for the session will be the intersection between this polcy and the role's policies.
- **policy_arns** (List of String) Amazon Resource Names (ARNs) of IAM Policies to use as managed session policies. The effective permissions for the session will be the intersection between these polcy and the role's policies.
- **role_arn** (String) Amazon Resource Name (ARN) of the IAM Role to assume.
- **session_name** (String) Session name to use when assuming the role.
- **tags** (Map of String) Map of assume role session tags.
- **transitive_tag_keys** (Set of String) Set of assume role session tag keys to pass to any subsequent sessions.


<a id="nestedatt--user_agent"></a>
### Nested Schema for `user_agent`

Optional:

- **comment** (String) User-Agent comment. At least one of `comment` or `product_name` must be set.
- **product_name** (String) Product name. At least one of `product_name` or `comment` must be set.
- **product_version** (String) Product version. Optional, and should only be set when `product_name` is set.
