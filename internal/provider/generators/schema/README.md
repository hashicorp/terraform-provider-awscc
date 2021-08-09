# AWS CloudFormation Resource Schema Downloader

Download and verify AWS CloudFormation resource schemas.

This tool

* Reads an [HCL](https://github.com/hashicorp/hcl) configuration file listing the CloudFormation resource type schemas included in the provider
* Checks whether a local copy of each configured schema is available and if not, downloads the schema from the CloudFormation registry and stores it locally
* Verifies that each schema conforms to the [correct meta-schema](https://github.com/aws-cloudformation/cloudformation-cli/blob/master/src/rpdk/core/data/schema/provider.definition.schema.v1.json)
* Generates a Go source file that contains instructions for [`go generate`](https://blog.golang.org/generate) to drive the [Terraform Resource Schema Generator](../resource/README.md)

Run `go run internal/provider/generators/schema/main.go --help` to see all options.

Note that valid AWS credentials must be available via [standard mechanisms](https://docs.aws.amazon.com/sdkref/latest/guide/environment-variables.html) to download a resource type schema from the [CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html).

### Configuration File

The `-config` command-line argument specifies a configuration file.
The file contains 2 types of block.

#### `meta_schema` Block

The `meta_schema` block specifies the details of the CloudFormation resource meta-schema.
Each configuration file must contain a single `meta_schema` block.

```hcl
meta_schema {
  # Path to the meta-schema file. Required.
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}
```

#### `resource_schema` Block

Each `resource_schema` block specifies the details of a single CloudFormation resource schema.
Each configuration file contains zero or more `resource_schema` blocks.

```hcl
resource_schema "aws_ec2_instance" {
  # CloudFormation type name. Required.
  cloudformation_type_name = "AWS::EC2::Instance"
}
```
