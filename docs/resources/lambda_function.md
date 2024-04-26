---
page_title: "awscc_lambda_function Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::Lambda::Function resource creates a Lambda function. To create a function, you need a deployment package https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html and an execution role https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html. The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
  You set the package type to Image if the deployment package is a container image https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html. For a container image, the code property must include the URI of a container image in the Amazon ECR registry. You do not need to specify the handler and runtime properties.
  You set the package type to Zip if the deployment package is a .zip file archive https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip. For a .zip file archive, the code property specifies the location of the .zip file. You must also specify the handler and runtime properties. For a Python example, see Deploy Python Lambda functions with .zip file archives https://docs.aws.amazon.com/lambda/latest/dg/python-package.html.
  You can use code signing https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html if your deployment package is a .zip file archive. To enable code signing for this function, specify the ARN of a code-signing configuration. When a user attempts to deploy a code package with UpdateFunctionCode, Lambda checks that the code package has a valid signature from a trusted publisher. The code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
  Note that you configure provisioned concurrency https://docs.aws.amazon.com/lambda/latest/dg/provisioned-concurrency.html on a AWS::Lambda::Version or a AWS::Lambda::Alias.
  For a complete introduction to Lambda functions, see What is Lambda? https://docs.aws.amazon.com/lambda/latest/dg/lambda-welcome.html in the Lambda developer guide.
---

# awscc_lambda_function (Resource)

The ``AWS::Lambda::Function`` resource creates a Lambda function. To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html). The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
 You set the package type to ``Image`` if the deployment package is a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html). For a container image, the code property must include the URI of a container image in the Amazon ECR registry. You do not need to specify the handler and runtime properties. 
 You set the package type to ``Zip`` if the deployment package is a [.zip file archive](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip). For a .zip file archive, the code property specifies the location of the .zip file. You must also specify the handler and runtime properties. For a Python example, see [Deploy Python Lambda functions with .zip file archives](https://docs.aws.amazon.com/lambda/latest/dg/python-package.html).
 You can use [code signing](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html) if your deployment package is a .zip file archive. To enable code signing for this function, specify the ARN of a code-signing configuration. When a user attempts to deploy a code package with ``UpdateFunctionCode``, Lambda checks that the code package has a valid signature from a trusted publisher. The code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
 Note that you configure [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/provisioned-concurrency.html) on a ``AWS::Lambda::Version`` or a ``AWS::Lambda::Alias``.
 For a complete introduction to Lambda functions, see [What is Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/lambda-welcome.html) in the *Lambda developer guide.*

## Example Usage

### Basic example
To create a AWS lambda function with basic details
```terraform
resource "awscc_iam_role" "main" {
  description = "AWS IAM role for lambda function"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

data "archive_file" "main" {
  type        = "zip"
  source_file = "main.py"
  output_path = "lambda_function_payload.zip"
}

resource "awscc_lambda_function" "main" {
  function_name = "lambda_function_name"
  description   = "AWS Lambda function"
  code = {
    zip_file = data.archive_file.main.output_path
  }
  package_type  = "Zip"
  handler       = "main.lambda_handler"
  runtime       = "python3.10"
  timeout       = "300"
  memory_size   = "128"
  role          = awscc_iam_role.main.arn
  architectures = ["arm64"]
  environment = {
    variables = {
      MY_KEY_1 = "MY_VALUE_1"
      MY_KEY_2 = "MY_VALUE_2"
    }
  }
}
```

### Lambda Layer example
To create a AWS lambda function using lambda layers
```terraform
resource "awscc_iam_role" "main" {
  description = "AWS IAM role for lambda function"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

data "archive_file" "main" {
  type        = "zip"
  source_file = "main.py"
  output_path = "lambda_function_payload.zip"
}

resource "aws_lambda_layer_version" "lambda_layer" {
  filename   = "lambda_layer_payload.zip"
  layer_name = "lambda_layer_name"

  compatible_runtimes = ["python3.10"]
}

resource "awscc_lambda_function" "main" {
  function_name = "lambda_function_name"
  description   = "AWS Lambda function"
  code = {
    zip_file = data.archive_file.main.output_path
  }
  handler       = "main.lambda_handler"
  runtime       = "python3.10"
  layers        = [aws_lambda_layer_version.lambda_layer.arn]
  timeout       = "300"
  memory_size   = "128"
  role          = awscc_iam_role.main.arn
  architectures = ["arm64"]
  environment = {
    variables = {
      MY_KEY_1 = "MY_VALUE_1"
      MY_KEY_2 = "MY_VALUE_2"
    }
  }
}
```

### Ephemeral storage example
To create a AWS lambda function using Ephemeral storage
```terraform
resource "awscc_iam_role" "main" {
  description = "AWS IAM role for lambda function"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

data "archive_file" "main" {
  type        = "zip"
  source_file = "main.py"
  output_path = "lambda_function_payload.zip"
}

resource "awscc_lambda_function" "main" {
  function_name = "lambda_function_name"
  description   = "AWS Lambda function"
  code = {
    zip_file = data.archive_file.main.output_path
  }
  handler       = "main.lambda_handler"
  runtime       = "python3.10"
  timeout       = "300"
  memory_size   = "128"
  role          = awscc_iam_role.main.arn
  architectures = ["arm64"]
  environment = {
    variables = {
      MY_KEY_1 = "MY_VALUE_1"
      MY_KEY_2 = "MY_VALUE_2"
    }
  }
  ephemeral_storage = {
    size = 10240 # Min 512 MB and the Max 10240 MB
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `code` (Attributes) The code for the function. (see [below for nested schema](#nestedatt--code))
- `role` (String) The Amazon Resource Name (ARN) of the function's execution role.

### Optional

- `architectures` (List of String) The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is ``x86_64``.
- `code_signing_config_arn` (String) To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
- `dead_letter_config` (Attributes) A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see [Dead-letter queues](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq). (see [below for nested schema](#nestedatt--dead_letter_config))
- `description` (String) A description of the function.
- `environment` (Attributes) Environment variables that are accessible from function code during execution. (see [below for nested schema](#nestedatt--environment))
- `ephemeral_storage` (Attributes) The size of the function's ``/tmp`` directory in MB. The default value is 512, but it can be any whole number between 512 and 10,240 MB. (see [below for nested schema](#nestedatt--ephemeral_storage))
- `file_system_configs` (Attributes List) Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) resource, you must also specify a ``DependsOn`` attribute to ensure that the mount target is created or updated before the function.
 For more information about using the ``DependsOn`` attribute, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html). (see [below for nested schema](#nestedatt--file_system_configs))
- `function_name` (String) The name of the Lambda function, up to 64 characters in length. If you don't specify a name, CFN generates one.
 If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
- `handler` (String) The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Lambda programming model](https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).
- `image_config` (Attributes) Configuration values that override the container image Dockerfile settings. For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms). (see [below for nested schema](#nestedatt--image_config))
- `kms_key_arn` (String) The ARN of the KMSlong (KMS) customer managed key that's used to encrypt your function's [environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption). When [Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html) is activated, Lambda also uses this key is to encrypt your function's snapshot. If you deploy your function using a container image, Lambda also uses this key to encrypt your function when it's deployed. Note that this is not the same key that's used to protect your container image in the Amazon Elastic Container Registry (Amazon ECR). If you don't provide a customer managed key, Lambda uses a default service key.
- `layers` (List of String) A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.
- `logging_config` (Attributes) The function's Amazon CloudWatch Logs configuration settings. (see [below for nested schema](#nestedatt--logging_config))
- `memory_size` (Number) The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB. Note that new AWS accounts have reduced concurrency and memory quotas. AWS raises these quotas automatically based on your usage. You can also request a quota increase.
- `package_type` (String) The type of deployment package. Set to ``Image`` for container image and set ``Zip`` for .zip file archive.
- `reserved_concurrent_executions` (Number) The number of simultaneous executions to reserve for the function.
- `runtime` (String) The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.
 The following list includes deprecated runtimes. For more information, see [Runtime deprecation policy](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy).
- `runtime_management_config` (Attributes) Sets the runtime management configuration for a function's version. For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html). (see [below for nested schema](#nestedatt--runtime_management_config))
- `snap_start` (Attributes) The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting. (see [below for nested schema](#nestedatt--snap_start))
- `tags` (Attributes Set) A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function. (see [below for nested schema](#nestedatt--tags))
- `timeout` (Number) The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).
- `tracing_config` (Attributes) Set ``Mode`` to ``Active`` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html). (see [below for nested schema](#nestedatt--tracing_config))
- `vpc_config` (Attributes) For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see [Configuring a Lambda function to access resources in a VPC](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html). (see [below for nested schema](#nestedatt--vpc_config))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `snap_start_response` (Attributes) The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting. (see [below for nested schema](#nestedatt--snap_start_response))

<a id="nestedatt--code"></a>
### Nested Schema for `code`

Optional:

- `image_uri` (String) URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.
- `s3_bucket` (String) An Amazon S3 bucket in the same AWS-Region as your function. The bucket can be in a different AWS-account.
- `s3_key` (String) The Amazon S3 key of the deployment package.
- `s3_object_version` (String) For versioned objects, the version of the deployment package object to use.
- `zip_file` (String) (Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, CFN places it in a file named ``index`` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html). This zip file cannot exceed 4MB. For the ``Handler`` property, the first part of the handler identifier must be ``index``. For example, ``index.handler``.
  For JSON, you must escape quotes and special characters such as newline (``\n``) with a backslash.
 If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ([cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html)) that simplifies sending responses. See [Using Lambda with CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for details.


<a id="nestedatt--dead_letter_config"></a>
### Nested Schema for `dead_letter_config`

Optional:

- `target_arn` (String) The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.


<a id="nestedatt--environment"></a>
### Nested Schema for `environment`

Optional:

- `variables` (Map of String) Environment variable key-value pairs. For more information, see [Using Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html).


<a id="nestedatt--ephemeral_storage"></a>
### Nested Schema for `ephemeral_storage`

Required:

- `size` (Number) The size of the function's ``/tmp`` directory.


<a id="nestedatt--file_system_configs"></a>
### Nested Schema for `file_system_configs`

Required:

- `arn` (String) The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
- `local_mount_path` (String) The path where the function can access the file system, starting with ``/mnt/``.


<a id="nestedatt--image_config"></a>
### Nested Schema for `image_config`

Optional:

- `command` (List of String) Specifies parameters that you want to pass in with ENTRYPOINT. You can specify a maximum of 1,500 parameters in the list.
- `entry_point` (List of String) Specifies the entry point to their application, which is typically the location of the runtime executable. You can specify a maximum of 1,500 string entries in the list.
- `working_directory` (String) Specifies the working directory. The length of the directory string cannot exceed 1,000 characters.


<a id="nestedatt--logging_config"></a>
### Nested Schema for `logging_config`

Optional:

- `application_log_level` (String) Set this property to filter the application logs for your function that Lambda sends to CloudWatch. Lambda only sends application logs at the selected level of detail and lower, where ``TRACE`` is the highest level and ``FATAL`` is the lowest.
- `log_format` (String) The format in which Lambda sends your function's application and system logs to CloudWatch. Select between plain text and structured JSON.
- `log_group` (String) The name of the Amazon CloudWatch log group the function sends logs to. By default, Lambda functions send logs to a default log group named ``/aws/lambda/<function name>``. To use a different log group, enter an existing log group or enter a new log group name.
- `system_log_level` (String) Set this property to filter the system logs for your function that Lambda sends to CloudWatch. Lambda only sends system logs at the selected level of detail and lower, where ``DEBUG`` is the highest level and ``WARN`` is the lowest.


<a id="nestedatt--runtime_management_config"></a>
### Nested Schema for `runtime_management_config`

Required:

- `update_runtime_on` (String) Specify the runtime update mode.
  +   *Auto (default)* - Automatically update to the most recent and secure runtime version using a [Two-phase runtime version rollout](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-two-phase). This is the best choice for most customers to ensure they always benefit from runtime updates.
  +   *FunctionUpdate* - LAM updates the runtime of you function to the most recent and secure runtime version when you update your function. This approach synchronizes runtime updates with function deployments, giving you control over when runtime updates are applied and allowing you to detect and mitigate rare runtime update incompatibilities early. When using this setting, you need to regularly update your functions to keep their runtime up-to-date.
  +   *Manual* - You specify a runtime version in your function configuration. The function will use this runtime version indefinitely. In the rare case where a new runtime version is incompatible with an existing function, this allows you to roll back your function to an earlier runtime version. For more information, see [Roll back a runtime version](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-rollback).
  
  *Valid Values*: ``Auto`` | ``FunctionUpdate`` | ``Manual``

Optional:

- `runtime_version_arn` (String) The ARN of the runtime version you want the function to use.
  This is only required if you're using the *Manual* runtime update mode.


<a id="nestedatt--snap_start"></a>
### Nested Schema for `snap_start`

Required:

- `apply_on` (String) Set ``ApplyOn`` to ``PublishedVersions`` to create a snapshot of the initialized execution environment when you publish a function version.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)

Optional:

- `value` (String)


<a id="nestedatt--tracing_config"></a>
### Nested Schema for `tracing_config`

Optional:

- `mode` (String) The tracing mode.


<a id="nestedatt--vpc_config"></a>
### Nested Schema for `vpc_config`

Optional:

- `ipv_6_allowed_for_dual_stack` (Boolean) Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack subnets.
- `security_group_ids` (List of String) A list of VPC security group IDs.
- `subnet_ids` (List of String) A list of VPC subnet IDs.


<a id="nestedatt--snap_start_response"></a>
### Nested Schema for `snap_start_response`

Read-Only:

- `apply_on` (String) When set to ``PublishedVersions``, Lambda creates a snapshot of the execution environment when you publish a function version.
- `optimization_status` (String) When you provide a [qualified Amazon Resource Name (ARN)](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using), this response element indicates whether SnapStart is activated for the specified function version.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_lambda_function.example <resource ID>
```