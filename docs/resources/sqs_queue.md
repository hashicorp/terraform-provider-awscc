---
page_title: "awscc_sqs_queue Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::SQS::Queue resource creates an SQS standard or FIFO queue.
  Keep the following caveats in mind:
  If you don't specify the FifoQueue property, SQS creates a standard queue.
  You can't change the queue type after you create it and you can't convert an existing standard queue into a FIFO queue. You must either create a new FIFO queue for your application or delete your existing standard queue and recreate it as a FIFO queue. For more information, see Moving from a standard queue to a FIFO queue https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-moving.html in the Developer Guide.If you don't provide a value for a property, the queue is created with the default value for the property.If you delete a queue, you must wait at least 60 seconds before creating a queue with the same name.To successfully create a new queue, you must provide a queue name that adheres to the limits related to queues https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html and is unique within the scope of your queues.
  For more information about creating FIFO (first-in-first-out) queues, see Creating an queue () https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/create-queue-cloudformation.html in the Developer Guide.
---

# awscc_sqs_queue (Resource)

The ``AWS::SQS::Queue`` resource creates an SQS standard or FIFO queue.
 Keep the following caveats in mind:
  +  If you don't specify the ``FifoQueue`` property, SQS creates a standard queue.
  You can't change the queue type after you create it and you can't convert an existing standard queue into a FIFO queue. You must either create a new FIFO queue for your application or delete your existing standard queue and recreate it as a FIFO queue. For more information, see [Moving from a standard queue to a FIFO queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-moving.html) in the *Developer Guide*. 
   +  If you don't provide a value for a property, the queue is created with the default value for the property.
  +  If you delete a queue, you must wait at least 60 seconds before creating a queue with the same name.
  +  To successfully create a new queue, you must provide a queue name that adheres to the [limits related to queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html) and is unique within the scope of your queues.
  
 For more information about creating FIFO (first-in-first-out) queues, see [Creating an queue ()](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/create-queue-cloudformation.html) in the *Developer Guide*.

## Example Usage

To create a SQS queue with tags
```terraform
resource "awscc_sqs_queue" "terraform_awscc_queue" {
  queue_name                        = "terraform-awscc-queue-example"
  delay_seconds                     = 90
  maximum_message_size              = 2048
  message_retention_period          = 86400
  receive_message_wait_time_seconds = 10
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

To create a SQS FIFO queue
```terraform
resource "awscc_sqs_queue" "terraform_awscc_queue_fifo" {
  queue_name                  = "terraform-awscc-queue-example.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

To create a SQS High-throughput FIFO queue
```terraform
resource "awscc_sqs_queue" "terraform_awscc_queue_high_throughput" {
  queue_name            = "terraform-awscc-queue-high-throughput-example.fifo"
  fifo_queue            = true
  deduplication_scope   = "messageGroup"
  fifo_throughput_limit = "perMessageGroupId"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

To create a SQS Dead-letter queue
```terraform
resource "awscc_sqs_queue" "terraform_awscc_queue" {
  queue_name                        = "terraform-awscc-queue-example"
  delay_seconds                     = 90
  maximum_message_size              = 2048
  message_retention_period          = 86400
  receive_message_wait_time_seconds = 10
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_sqs_queue" "terraform_awscc_queue_deadletter" {
  queue_name = "terraform-awscc-queue-deadletter-example"
  redrive_allow_policy = jsonencode({
    redrivePermission = "byQueue",
    sourceQueueArns   = [awscc_sqs_queue.terraform_awscc_queue.arn]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

To create a SQS queue using Server-side encryption (SSE)
```terraform
resource "awscc_sqs_queue" "terraform_awscc_queue_sse" {
  queue_name              = "terraform-awscc-queue-sse-example"
  sqs_managed_sse_enabled = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

To create a SQS queue using KMS key
```terraform
resource "awscc_sqs_queue" "terraform_awscc_queue_kms" {
  queue_name                        = "terraform-awscc-queue-kms-example"
  kms_master_key_id                 = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `content_based_deduplication` (Boolean) For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the ``ContentBasedDeduplication`` attribute for the ``CreateQueue`` action in the *API Reference*.
- `deduplication_scope` (String) For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level. Valid values are ``messageGroup`` and ``queue``.
 To enable high throughput for a FIFO queue, set this attribute to ``messageGroup`` *and* set the ``FifoThroughputLimit`` attribute to ``perMessageGroupId``. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.
- `delay_seconds` (Number) The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of ``0`` to ``900`` (15 minutes). The default value is ``0``.
- `fifo_queue` (Boolean) If set to true, creates a FIFO queue. If you don't specify this property, SQS creates a standard queue. For more information, see [Amazon SQS FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html) in the *Developer Guide*.
- `fifo_throughput_limit` (String) For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are ``perQueue`` and ``perMessageGroupId``.
 To enable high throughput for a FIFO queue, set this attribute to ``perMessageGroupId`` *and* set the ``DeduplicationScope`` attribute to ``messageGroup``. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.
- `kms_data_key_reuse_period_seconds` (Number) The length of time in seconds for which SQS can reuse a data key to encrypt or decrypt messages before calling KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
  A shorter time period provides better security, but results in more calls to KMS, which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Developer Guide*.
- `kms_master_key_id` (String) The ID of an AWS Key Management Service (KMS) for SQS, or a custom KMS. To use the AWS managed KMS for SQS, specify a (default) alias ARN, alias name (for example ``alias/aws/sqs``), key ARN, or key ID. For more information, see the following:
  +   [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Developer Guide* 
  +   [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *API Reference* 
  +   [Request Parameters](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters) in the *Key Management Service API Reference* 
  +   The Key Management Service (KMS) section of the [Security best practices for Key Management Service](https://docs.aws.amazon.com/kms/latest/developerguide/best-practices.html) in the *Key Management Service Developer Guide*
- `maximum_message_size` (Number) The limit of how many bytes that a message can contain before SQS rejects it. You can specify an integer value from ``1,024`` bytes (1 KiB) to ``262,144`` bytes (256 KiB). The default value is ``262,144`` (256 KiB).
- `message_retention_period` (Number) The number of seconds that SQS retains a message. You can specify an integer value from ``60`` seconds (1 minute) to ``1,209,600`` seconds (14 days). The default value is ``345,600`` seconds (4 days).
- `queue_name` (String) A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the ``.fifo`` suffix. For more information, see [Amazon SQS FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html) in the *Developer Guide*.
 If you don't specify a name, CFN generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *User Guide*. 
  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
- `receive_message_wait_time_seconds` (Number) Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Developer Guide*.
- `redrive_allow_policy` (String) The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object. The parameters are as follows:
  +   ``redrivePermission``: The permission type that defines which source queues can specify the current queue as the dead-letter queue. Valid values are:
  +   ``allowAll``: (Default) Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.
  +   ``denyAll``: No source queues can specify this queue as the dead-letter queue.
  +   ``byQueue``: Only queues specified by the ``sourceQueueArns`` parameter can specify this queue as the dead-letter queue.
  
  +   ``sourceQueueArns``: The Amazon Resource Names (ARN)s of the source queues that can specify this queue as the dead-letter queue and redrive messages. You can specify this parameter only when the ``redrivePermission`` parameter is set to ``byQueue``. You can specify up to 10 source queue ARNs. To allow more than 10 source queues to specify dead-letter queues, set the ``redrivePermission`` parameter to ``allowAll``.
- `redrive_policy` (String) The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object. The parameters are as follows:
  +   ``deadLetterTargetArn``: The Amazon Resource Name (ARN) of the dead-letter queue to which SQS moves messages after the value of ``maxReceiveCount`` is exceeded.
  +   ``maxReceiveCount``: The number of times a message is received by a consumer of the source queue before being moved to the dead-letter queue. When the ``ReceiveCount`` for a message exceeds the ``maxReceiveCount`` for a queue, SQS moves the message to the dead-letter-queue.
  
  The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a standard queue must also be a standard queue.
   *JSON* 
  ``{ "deadLetterTargetArn" : String, "maxReceiveCount" : Integer }`` 
  *YAML* 
  ``deadLetterTargetArn : String`` 
  ``maxReceiveCount : Integer``
- `sqs_managed_sse_enabled` (Boolean) Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (for example, [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html) or [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)). When ``SqsManagedSseEnabled`` is not defined, ``SSE-SQS`` encryption is enabled by default.
- `tags` (Attributes List) The tags that you attach to this queue. For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *User Guide*. (see [below for nested schema](#nestedatt--tags))
- `visibility_timeout` (Number) The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.
 Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
 For more information about SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Developer Guide*.

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `queue_url` (String)

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
  to = awscc_sqs_queue.example
  id = "queue_url"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_sqs_queue.example "queue_url"
```