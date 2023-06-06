---
page_title: "awscc_cloudfront_distribution Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::CloudFront::Distribution
---

# awscc_cloudfront_distribution (Resource)

Resource Type definition for AWS::CloudFront::Distribution

## Example Usage

### Cloudfront Distribution with S3 Origin using Origin Access Control

```terraform
# S3 Bucket Origin with bucket policy to Origin Access Control
resource "aws_s3_bucket" "s3_origin" {
  bucket = "sampleawsccbucket345"
}

# Block public access to S3 bucket
resource "aws_s3_bucket_public_access_block" "s3_block_public_access" {
  bucket                  = aws_s3_bucket.s3_origin.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Attach bucket policy with object access to cloudfront origin
resource "aws_s3_bucket_policy" "allow_access_from_cloudfront" {
  bucket = aws_s3_bucket.s3_origin.id
  policy = data.aws_iam_policy_document.bucket_policy.json
}

# IAM policy document to allow S3 bucket read access to cloudfront origin access control
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
    effect = "Allow"
    actions = [
      "s3:GetObject",
    ]
    resources = [
      "arn:aws:s3:::${aws_s3_bucket.s3_origin.id}/*"
    ]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = ["arn:aws:cloudfront::111111111111:distribution/${awscc_cloudfront_distribution.cloudfront_s3_origin.id}"]
    }
  }
}

# Cloudfront origin access control using AWSCC provider
resource "awscc_cloudfront_origin_access_control" "cf_oac" {
  origin_access_control_config = {
    name                              = "sample-oac"
    description                       = "Sample Origin Access Control Setting using AWSCC"
    origin_access_control_origin_type = "s3"
    signing_behavior                  = "always"
    signing_protocol                  = "sigv4"
  }
}

# Cloudfront distribution with S3 origin using AWSCC provider
resource "awscc_cloudfront_distribution" "cloudfront_s3_origin" {
  distribution_config = {
    enabled             = true
    compress            = true
    default_root_object = "index.html"
    comment             = "Sample Cloudfront Distribution using AWSCC provider"
    default_cache_behavior = {
      target_origin_id       = aws_s3_bucket.s3_origin.id
      viewer_protocol_policy = "redirect-to-https"
      allowed_methods        = ["GET", "HEAD", "OPTIONS"]
      cached_methods         = ["GET", "HEAD", "OPTIONS"]
      min_ttl                = 0
      default_ttl            = 5 * 60
      max_ttl                = 60 * 60
    }
    restrictions = {
      geo_restriction = {
        restriction_type = "none"
      }
    }
    viewer_certificate = {
      cloudfront_default_certificate = true
      minimum_protocol_version       = "TLSv1.2_2018"
    }
    s3_origin = {
      dns_name = aws_s3_bucket.s3_origin.bucket_regional_domain_name
    }
    origins = [{
      domain_name              = aws_s3_bucket.s3_origin.bucket_regional_domain_name
      id                       = "SampleCloudfrontOrigin"
      origin_access_control_id = awscc_cloudfront_origin_access_control.cf_oac.id
    }]
  }
  tags = [{
    key   = "Name"
    value = "Cloudfront Distribution with S3 Origin"
  }]
}
```

### Cloudfront Distribution with S3 Origin using Origin Access Identity

```terraform
# S3 Bucket Origin with bucket policy to Origin Access Control
resource "aws_s3_bucket" "s3_origin" {
  bucket = "sampleawsccbucket345"
}

# Block public access to S3 bucket
resource "aws_s3_bucket_public_access_block" "s3_block_public_access" {
  bucket                  = aws_s3_bucket.s3_origin.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Attach bucket policy with object access to cloudfront origin
resource "aws_s3_bucket_policy" "allow_access_from_cloudfront" {
  bucket = aws_s3_bucket.s3_origin.id
  policy = data.aws_iam_policy_document.bucket_policy.json
}

# IAM policy document to allow S3 bucket read access to cloudfront origin access identity
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    principals {
      type        = "CanonicalUser"
      identifiers = [awscc_cloudfront_cloudfront_origin_access_identity.cf_oai.s3_canonical_user_id]
    }
    effect = "Allow"
    actions = [
      "s3:GetObject",
    ]
    resources = [
      "arn:aws:s3:::${aws_s3_bucket.s3_origin.id}/*"
    ]
  }
}

# Cloudfront origin access identity
resource "awscc_cloudfront_cloudfront_origin_access_identity" "cf_oai" {
  cloudfront_origin_access_identity_config = {
    comment = "SampleCloudFrontOAI"
  }
}

# Cloudfront distribution with S3 origin using AWSCC provider
resource "awscc_cloudfront_distribution" "cloudfront_s3_origin" {
  distribution_config = {
    enabled             = true
    compress            = true
    default_root_object = "index.html"
    comment             = "Sample Cloudfront Distribution using AWSCC provider"
    default_cache_behavior = {
      target_origin_id       = aws_s3_bucket.s3_origin.id
      viewer_protocol_policy = "redirect-to-https"
      allowed_methods        = ["GET", "HEAD", "OPTIONS"]
      cached_methods         = ["GET", "HEAD", "OPTIONS"]
      min_ttl                = 0
      default_ttl            = 5 * 60
      max_ttl                = 60 * 60
    }
    restrictions = {
      geo_restriction = {
        restriction_type = "none"
      }
    }
    viewer_certificate = {
      cloudfront_default_certificate = true
      minimum_protocol_version       = "TLSv1.2_2018"
    }
    s3_origin = {
      dns_name = aws_s3_bucket.s3_origin.bucket_regional_domain_name
    }
    origins = [{
      domain_name = aws_s3_bucket.s3_origin.bucket_regional_domain_name
      id          = "SampleCloudfrontOrigin"
      s3_origin_config = {
        origin_access_identity = awscc_cloudfront_cloudfront_origin_access_identity.cf_oai.id
      }
    }]
  }
  tags = [{
    key   = "Name"
    value = "Cloudfront Distribution with S3 Origin"
  }]
}
```

## NOTE: After successful resource creation, edit the Cloufront Origin, origin access from default public setting to use either the OAC or OAI which were created

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `distribution_config` (Attributes) (see [below for nested schema](#nestedatt--distribution_config))

### Optional

- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `domain_name` (String)
- `id` (String) The ID of this resource.

<a id="nestedatt--distribution_config"></a>
### Nested Schema for `distribution_config`

Required:

- `default_cache_behavior` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--default_cache_behavior))
- `enabled` (Boolean)

Optional:

- `aliases` (List of String)
- `cache_behaviors` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--cache_behaviors))
- `cnames` (List of String)
- `comment` (String)
- `continuous_deployment_policy_id` (String)
- `custom_error_responses` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--custom_error_responses))
- `custom_origin` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--custom_origin))
- `default_root_object` (String)
- `http_version` (String)
- `ipv6_enabled` (Boolean)
- `logging` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--logging))
- `origin_groups` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origin_groups))
- `origins` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--origins))
- `price_class` (String)
- `restrictions` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--restrictions))
- `s3_origin` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--s3_origin))
- `staging` (Boolean)
- `viewer_certificate` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--viewer_certificate))
- `web_acl_id` (String)

<a id="nestedatt--distribution_config--default_cache_behavior"></a>
### Nested Schema for `distribution_config.default_cache_behavior`

Required:

- `target_origin_id` (String)
- `viewer_protocol_policy` (String)

Optional:

- `allowed_methods` (List of String)
- `cache_policy_id` (String)
- `cached_methods` (List of String)
- `compress` (Boolean)
- `default_ttl` (Number)
- `field_level_encryption_id` (String)
- `forwarded_values` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--default_cache_behavior--forwarded_values))
- `function_associations` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--default_cache_behavior--function_associations))
- `lambda_function_associations` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--default_cache_behavior--lambda_function_associations))
- `max_ttl` (Number)
- `min_ttl` (Number)
- `origin_request_policy_id` (String)
- `realtime_log_config_arn` (String)
- `response_headers_policy_id` (String)
- `smooth_streaming` (Boolean)
- `trusted_key_groups` (List of String)
- `trusted_signers` (List of String)

<a id="nestedatt--distribution_config--default_cache_behavior--forwarded_values"></a>
### Nested Schema for `distribution_config.default_cache_behavior.forwarded_values`

Required:

- `query_string` (Boolean)

Optional:

- `cookies` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--default_cache_behavior--forwarded_values--cookies))
- `headers` (List of String)
- `query_string_cache_keys` (List of String)

<a id="nestedatt--distribution_config--default_cache_behavior--forwarded_values--cookies"></a>
### Nested Schema for `distribution_config.default_cache_behavior.forwarded_values.query_string_cache_keys`

Required:

- `forward` (String)

Optional:

- `whitelisted_names` (List of String)



<a id="nestedatt--distribution_config--default_cache_behavior--function_associations"></a>
### Nested Schema for `distribution_config.default_cache_behavior.function_associations`

Optional:

- `event_type` (String)
- `function_arn` (String)


<a id="nestedatt--distribution_config--default_cache_behavior--lambda_function_associations"></a>
### Nested Schema for `distribution_config.default_cache_behavior.lambda_function_associations`

Optional:

- `event_type` (String)
- `include_body` (Boolean)
- `lambda_function_arn` (String)



<a id="nestedatt--distribution_config--cache_behaviors"></a>
### Nested Schema for `distribution_config.cache_behaviors`

Required:

- `path_pattern` (String)
- `target_origin_id` (String)
- `viewer_protocol_policy` (String)

Optional:

- `allowed_methods` (List of String)
- `cache_policy_id` (String)
- `cached_methods` (List of String)
- `compress` (Boolean)
- `default_ttl` (Number)
- `field_level_encryption_id` (String)
- `forwarded_values` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--cache_behaviors--forwarded_values))
- `function_associations` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--cache_behaviors--function_associations))
- `lambda_function_associations` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--cache_behaviors--lambda_function_associations))
- `max_ttl` (Number)
- `min_ttl` (Number)
- `origin_request_policy_id` (String)
- `realtime_log_config_arn` (String)
- `response_headers_policy_id` (String)
- `smooth_streaming` (Boolean)
- `trusted_key_groups` (List of String)
- `trusted_signers` (List of String)

<a id="nestedatt--distribution_config--cache_behaviors--forwarded_values"></a>
### Nested Schema for `distribution_config.cache_behaviors.forwarded_values`

Required:

- `query_string` (Boolean)

Optional:

- `cookies` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--cache_behaviors--forwarded_values--cookies))
- `headers` (List of String)
- `query_string_cache_keys` (List of String)

<a id="nestedatt--distribution_config--cache_behaviors--forwarded_values--cookies"></a>
### Nested Schema for `distribution_config.cache_behaviors.forwarded_values.query_string_cache_keys`

Required:

- `forward` (String)

Optional:

- `whitelisted_names` (List of String)



<a id="nestedatt--distribution_config--cache_behaviors--function_associations"></a>
### Nested Schema for `distribution_config.cache_behaviors.function_associations`

Optional:

- `event_type` (String)
- `function_arn` (String)


<a id="nestedatt--distribution_config--cache_behaviors--lambda_function_associations"></a>
### Nested Schema for `distribution_config.cache_behaviors.lambda_function_associations`

Optional:

- `event_type` (String)
- `include_body` (Boolean)
- `lambda_function_arn` (String)



<a id="nestedatt--distribution_config--custom_error_responses"></a>
### Nested Schema for `distribution_config.custom_error_responses`

Required:

- `error_code` (Number)

Optional:

- `error_caching_min_ttl` (Number)
- `response_code` (Number)
- `response_page_path` (String)


<a id="nestedatt--distribution_config--custom_origin"></a>
### Nested Schema for `distribution_config.custom_origin`

Required:

- `dns_name` (String)
- `origin_protocol_policy` (String)
- `origin_ssl_protocols` (List of String)

Optional:

- `http_port` (Number)
- `https_port` (Number)


<a id="nestedatt--distribution_config--logging"></a>
### Nested Schema for `distribution_config.logging`

Required:

- `bucket` (String)

Optional:

- `include_cookies` (Boolean)
- `prefix` (String)


<a id="nestedatt--distribution_config--origin_groups"></a>
### Nested Schema for `distribution_config.origin_groups`

Required:

- `quantity` (Number)

Optional:

- `items` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--origin_groups--items))

<a id="nestedatt--distribution_config--origin_groups--items"></a>
### Nested Schema for `distribution_config.origin_groups.items`

Required:

- `failover_criteria` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origin_groups--items--failover_criteria))
- `id` (String)
- `members` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origin_groups--items--members))

<a id="nestedatt--distribution_config--origin_groups--items--failover_criteria"></a>
### Nested Schema for `distribution_config.origin_groups.items.members`

Required:

- `status_codes` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origin_groups--items--members--status_codes))

<a id="nestedatt--distribution_config--origin_groups--items--members--status_codes"></a>
### Nested Schema for `distribution_config.origin_groups.items.members.status_codes`

Required:

- `items` (List of Number)
- `quantity` (Number)



<a id="nestedatt--distribution_config--origin_groups--items--members"></a>
### Nested Schema for `distribution_config.origin_groups.items.members`

Required:

- `items` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--origin_groups--items--members--items))
- `quantity` (Number)

<a id="nestedatt--distribution_config--origin_groups--items--members--items"></a>
### Nested Schema for `distribution_config.origin_groups.items.members.items`

Required:

- `origin_id` (String)





<a id="nestedatt--distribution_config--origins"></a>
### Nested Schema for `distribution_config.origins`

Required:

- `domain_name` (String)
- `id` (String)

Optional:

- `connection_attempts` (Number)
- `connection_timeout` (Number)
- `custom_origin_config` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origins--custom_origin_config))
- `origin_access_control_id` (String)
- `origin_custom_headers` (Attributes List) (see [below for nested schema](#nestedatt--distribution_config--origins--origin_custom_headers))
- `origin_path` (String)
- `origin_shield` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origins--origin_shield))
- `s3_origin_config` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--origins--s3_origin_config))

<a id="nestedatt--distribution_config--origins--custom_origin_config"></a>
### Nested Schema for `distribution_config.origins.custom_origin_config`

Required:

- `origin_protocol_policy` (String)

Optional:

- `http_port` (Number)
- `https_port` (Number)
- `origin_keepalive_timeout` (Number)
- `origin_read_timeout` (Number)
- `origin_ssl_protocols` (List of String)


<a id="nestedatt--distribution_config--origins--origin_custom_headers"></a>
### Nested Schema for `distribution_config.origins.origin_custom_headers`

Required:

- `header_name` (String)
- `header_value` (String)


<a id="nestedatt--distribution_config--origins--origin_shield"></a>
### Nested Schema for `distribution_config.origins.origin_shield`

Optional:

- `enabled` (Boolean)
- `origin_shield_region` (String)


<a id="nestedatt--distribution_config--origins--s3_origin_config"></a>
### Nested Schema for `distribution_config.origins.s3_origin_config`

Optional:

- `origin_access_identity` (String)



<a id="nestedatt--distribution_config--restrictions"></a>
### Nested Schema for `distribution_config.restrictions`

Required:

- `geo_restriction` (Attributes) (see [below for nested schema](#nestedatt--distribution_config--restrictions--geo_restriction))

<a id="nestedatt--distribution_config--restrictions--geo_restriction"></a>
### Nested Schema for `distribution_config.restrictions.geo_restriction`

Required:

- `restriction_type` (String)

Optional:

- `locations` (List of String)



<a id="nestedatt--distribution_config--s3_origin"></a>
### Nested Schema for `distribution_config.s3_origin`

Required:

- `dns_name` (String)

Optional:

- `origin_access_identity` (String)


<a id="nestedatt--distribution_config--viewer_certificate"></a>
### Nested Schema for `distribution_config.viewer_certificate`

Optional:

- `acm_certificate_arn` (String)
- `cloudfront_default_certificate` (Boolean)
- `iam_certificate_id` (String)
- `minimum_protocol_version` (String)
- `ssl_support_method` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cloudfront_distribution.example <resource ID>
```