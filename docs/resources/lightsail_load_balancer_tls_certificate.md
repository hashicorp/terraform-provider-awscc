---
page_title: "awscc_lightsail_load_balancer_tls_certificate Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate
---

# awscc_lightsail_load_balancer_tls_certificate (Resource)

Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate

## Example Usage

```terraform
resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 443
  load_balancer_name = "example-lb"
}

# Note: Lightsail will verify the certificate during creation
resource "awscc_lightsail_load_balancer_tls_certificate" "example" {
  certificate_domain_name       = "example.com"
  certificate_name              = "example-lb-cert"
  load_balancer_name            = awscc_lightsail_load_balancer.example.load_balancer_name
  certificate_alternative_names = ["www.example.com"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `certificate_domain_name` (String) The domain name (e.g., example.com ) for your SSL/TLS certificate.
- `certificate_name` (String) The SSL/TLS certificate name.
- `load_balancer_name` (String) The name of your load balancer.

### Optional

- `certificate_alternative_names` (Set of String) An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.
- `https_redirection_enabled` (Boolean) A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.
- `is_attached` (Boolean) When true, the SSL/TLS certificate is attached to the Lightsail load balancer.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `load_balancer_tls_certificate_arn` (String)
- `status` (String) The validation status of the SSL/TLS certificate.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_lightsail_load_balancer_tls_certificate.example
  id = "certificate_name|load_balancer_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_lightsail_load_balancer_tls_certificate.example "certificate_name|load_balancer_name"
```