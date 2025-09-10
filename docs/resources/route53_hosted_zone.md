---
page_title: "awscc_route53_hosted_zone Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Creates a new public or private hosted zone. You create records in a public hosted zone to define how you want to route traffic on the internet for a domain, such as example.com, and its subdomains (apex.example.com, acme.example.com). You create records in a private hosted zone to define how you want to route traffic for a domain and its subdomains within one or more Amazon Virtual Private Clouds (Amazon VPCs).
  You can't convert a public hosted zone to a private hosted zone or vice versa. Instead, you must create a new hosted zone with the same name and create new resource record sets.
  For more information about charges for hosted zones, see Amazon Route 53 Pricing https://docs.aws.amazon.com/route53/pricing/.
  Note the following:
  You can't create a hosted zone for a top-level domain (TLD) such as .com.If your domain is registered with a registrar other than Route 53, you must update the name servers with your registrar to make Route 53 the DNS service for the domain. For more information, see Migrating DNS Service for an Existing Domain to Amazon Route 53 https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html in the Amazon Route 53 Developer Guide.
  When you submit a CreateHostedZone request, the initial status of the hosted zone is PENDING. For public hosted zones, this means that the NS and SOA records are not yet available on all Route 53 DNS servers. When the NS and SOA records are available, the status of the zone changes to INSYNC.
  The CreateHostedZone request requires the caller to have an ec2:DescribeVpcs permission.
  When creating private hosted zones, the Amazon VPC must belong to the same partition where the hosted zone is created. A partition is a group of AWS-Regions. Each AWS-account is scoped to one partition.
  The following are the supported partitions:
  aws - AWS-Regionsaws-cn - China Regionsaws-us-gov - govcloud-us-region
  For more information, see Access Management https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html in the General Reference.
---

# awscc_route53_hosted_zone (Resource)

Creates a new public or private hosted zone. You create records in a public hosted zone to define how you want to route traffic on the internet for a domain, such as example.com, and its subdomains (apex.example.com, acme.example.com). You create records in a private hosted zone to define how you want to route traffic for a domain and its subdomains within one or more Amazon Virtual Private Clouds (Amazon VPCs). 
  You can't convert a public hosted zone to a private hosted zone or vice versa. Instead, you must create a new hosted zone with the same name and create new resource record sets.
  For more information about charges for hosted zones, see [Amazon Route 53 Pricing](https://docs.aws.amazon.com/route53/pricing/).
 Note the following:
  +  You can't create a hosted zone for a top-level domain (TLD) such as .com.
  +  If your domain is registered with a registrar other than Route 53, you must update the name servers with your registrar to make Route 53 the DNS service for the domain. For more information, see [Migrating DNS Service for an Existing Domain to Amazon Route 53](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html) in the *Amazon Route 53 Developer Guide*. 
  
 When you submit a ``CreateHostedZone`` request, the initial status of the hosted zone is ``PENDING``. For public hosted zones, this means that the NS and SOA records are not yet available on all Route 53 DNS servers. When the NS and SOA records are available, the status of the zone changes to ``INSYNC``.
 The ``CreateHostedZone`` request requires the caller to have an ``ec2:DescribeVpcs`` permission.
  When creating private hosted zones, the Amazon VPC must belong to the same partition where the hosted zone is created. A partition is a group of AWS-Regions. Each AWS-account is scoped to one partition.
 The following are the supported partitions:
  +  ``aws`` - AWS-Regions
  +  ``aws-cn`` - China Regions
  +  ``aws-us-gov`` - govcloud-us-region
  
 For more information, see [Access Management](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.

## Example Usage

### Route53 Hosted Zone
To create public `awscc_route53_hosted_zone` use below:
```terraform
resource "awscc_route53_hosted_zone" "this" {
  name = "this.com"
  hosted_zone_tags = [
    {
      key   = "Name"
      value = "this"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```

To create a private `awscc_route53_hosted_zone` use below:
```terraform
resource "awscc_route53_hosted_zone" "this" {
  name = "this.com"
  hosted_zone_tags = [
    {
      key   = "Name"
      value = "this"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
  vp_cs = [
    {
      vpc_id     = awscc_ec2_vpc.main.id
      vpc_region = "us-east-1"
    }
  ]
}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `hosted_zone_config` (Attributes) A complex type that contains an optional comment.
 If you don't want to specify a comment, omit the ``HostedZoneConfig`` and ``Comment`` elements. (see [below for nested schema](#nestedatt--hosted_zone_config))
- `hosted_zone_tags` (Attributes Set) Adds, edits, or deletes tags for a health check or a hosted zone.
 For information about using tags for cost allocation, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the *User Guide*. (see [below for nested schema](#nestedatt--hosted_zone_tags))
- `name` (String) The name of the domain. Specify a fully qualified domain name, for example, *www.example.com*. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats *www.example.com* (without a trailing dot) and *www.example.com.* (with a trailing dot) as identical.
 If you're creating a public hosted zone, this is the name you have registered with your DNS registrar. If your domain name is registered with a registrar other than Route 53, change the name servers for your domain to the set of ``NameServers`` that are returned by the ``Fn::GetAtt`` intrinsic function.
- `query_logging_config` (Attributes) Creates a configuration for DNS query logging. After you create a query logging configuration, Amazon Route 53 begins to publish log data to an Amazon CloudWatch Logs log group.
 DNS query logs contain information about the queries that Route 53 receives for a specified public hosted zone, such as the following:
  +  Route 53 edge location that responded to the DNS query
  +  Domain or subdomain that was requested
  +  DNS record type, such as A or AAAA
  +  DNS response code, such as ``NoError`` or ``ServFail``
  
  + Log Group and Resource Policy Before you create a query logging configuration, perform the following operations. If you create a query logging configuration using the Route 53 console, Route 53 performs these operations automatically. Create a CloudWatch Logs log group, and make note of the ARN, which you specify when you create a query logging configuration. Note the following: You must create the log group in the us-east-1 region. You must use the same to create the log group and the hosted zone that you want to configure query logging for. When you create log groups for query logging, we recommend that you use a consistent prefix, for example: /aws/route53/hosted zone name In the next step, you'll create a resource policy, which controls access to one or more log groups and the associated resources, such as Route 53 hosted zones. There's a limit on the number of resource policies that you can create, so we recommend that you use a consistent prefix so you can use the same resource policy for all the log groups that you create for query logging. Create a CloudWatch Logs resource policy, and give it the permissions that Route 53 needs to create log streams and to send query logs to log streams. You must create the CloudWatch Logs resource policy in the us-east-1 region. For the value of Resource, specify the ARN for the log group that you created in the previous step. To use the same resource policy for all the CloudWatch Logs log groups that you created for query logging configurations, replace the hosted zone name with *, for example: arn:aws:logs:us-east-1:123412341234:log-group:/aws/route53/* To avoid the confused deputy problem, a security issue where an entity without a permission for an action can coerce a more-privileged entity to perform it, you can optionally limit the permissions that a service has to a resource in a resource-based policy by supplying the following values: For aws:SourceArn, supply the hosted zone ARN used in creating the query logging configuration. For example, aws:SourceArn: arn:aws:route53:::hostedzone/hosted zone ID. For aws:SourceAccount, supply the account ID for the account that creates the query logging configuration. For example, aws:SourceAccount:111111111111. For more information, see The confused deputy problem in the IAM User Guide. You can't use the CloudWatch console to create or edit a resource policy. You must use the CloudWatch API, one of the SDKs, or the . + Log Streams and Edge Locations When Route 53 finishes creating the configuration for DNS query logging, it does the following: Creates a log stream for an edge location the first time that the edge location responds to DNS queries for the specified hosted zone. That log stream is used to log all queries that Route 53 responds to for that edge location. Begins to send query logs to the applicable log stream. The name of each log stream is in the following format: hosted zone ID/edge location code The edge location code is a three-letter code and an arbitrarily assigned number, for example, DFW3. The three-letter code typically corresponds with the International Air Transport Association airport code for an airport near the edge location. (These abbreviations might change in the future.) For a list of edge locations, see "The Route 53 Global Network" on the Route 53 Product Details page. + Queries That Are Logged Query logs contain only the queries that DNS resolvers forward to Route 53. If a DNS resolver has already cached the response to a query (such as the IP address for a load balancer for example.com), the resolver will continue to return the cached response. It doesn't forward another query to Route 53 until the TTL for the corresponding resource record set expires. Depending on how many DNS queries are submitted for a resource record set, and depending on the TTL for that resource record set, query logs might contain information about only one query out of every several thousand queries that are submitted to DNS. For more information about how DNS works, see Routing Internet Traffic to Your Website or Web Application in the Amazon Route 53 Developer Guide. + Log File Format For a list of the values in each query log and the format of each value, see Logging DNS Queries in the Amazon Route 53 Developer Guide. + Pricing For information about charges for query logs, see Amazon CloudWatch Pricing. + How to Stop Logging If you want Route 53 to stop sending query logs to CloudWatch Logs, delete the query logging configuration. For more information, see DeleteQueryLoggingConfig. (see [below for nested schema](#nestedatt--query_logging_config))
- `vp_cs` (Attributes Set) *Private hosted zones:* A complex type that contains information about the VPCs that are associated with the specified hosted zone.
  For public hosted zones, omit ``VPCs``, ``VPCId``, and ``VPCRegion``. (see [below for nested schema](#nestedatt--vp_cs))

### Read-Only

- `hosted_zone_id` (String)
- `id` (String) Uniquely identifies the resource.
- `name_servers` (List of String)

<a id="nestedatt--hosted_zone_config"></a>
### Nested Schema for `hosted_zone_config`

Optional:

- `comment` (String) Any comments that you want to include about the hosted zone.


<a id="nestedatt--hosted_zone_tags"></a>
### Nested Schema for `hosted_zone_tags`

Optional:

- `key` (String) The value of ``Key`` depends on the operation that you want to perform:
  +  *Add a tag to a health check or hosted zone*: ``Key`` is the name that you want to give the new tag.
  +  *Edit a tag*: ``Key`` is the name of the tag that you want to change the ``Value`` for.
  +  *Delete a key*: ``Key`` is the name of the tag you want to remove.
  +  *Give a name to a health check*: Edit the default ``Name`` tag. In the Amazon Route 53 console, the list of your health checks includes a *Name* column that lets you see the name that you've given to each health check.
- `value` (String) The value of ``Value`` depends on the operation that you want to perform:
  +  *Add a tag to a health check or hosted zone*: ``Value`` is the value that you want to give the new tag.
  +  *Edit a tag*: ``Value`` is the new value that you want to assign the tag.


<a id="nestedatt--query_logging_config"></a>
### Nested Schema for `query_logging_config`

Optional:

- `cloudwatch_logs_log_group_arn` (String) The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.


<a id="nestedatt--vp_cs"></a>
### Nested Schema for `vp_cs`

Optional:

- `vpc_id` (String) *Private hosted zones only:* The ID of an Amazon VPC.
  For public hosted zones, omit ``VPCs``, ``VPCId``, and ``VPCRegion``.
- `vpc_region` (String) *Private hosted zones only:* The region that an Amazon VPC was created in.
  For public hosted zones, omit ``VPCs``, ``VPCId``, and ``VPCRegion``.

## Import

Import is supported using the following syntax:

In Terraform v1.12.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `identity` attribute, for example:

```terraform
import {
  to = awscc_route53_hosted_zone.example
  identity = {
    id = "id"
  }
}
```

<!-- schema generated by tfplugindocs -->
### Identity Schema

#### Required

- `id` (String)

#### Optional

- `account_id` (String) AWS Account where this resource is managed

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_route53_hosted_zone.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_route53_hosted_zone.example "id"
```
