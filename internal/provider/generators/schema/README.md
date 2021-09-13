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
The file contains 3 types of block.

#### `defaults` Block

The `defaults` block specifies defaults which can be overriden by each resource schema.
Each configuration file must contain a single `defaults` block.

```hcl
defaults {
  # Schema cache directory. Required.
  schema_cache_directory = "../service/cloudformation/schemas"

  # Prefix for Terraform type names. Optional.
  # The default is to use the label from the resource_schema block as the type name.
  terraform_type_name_prefix = "awscc"
}
```

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

  # Path to the CloudFormation schema file.
  # Optional.
  # The default value combines the `defaults.schema_cache_directory` value with the CloudFormation type name.
  cloudformation_schema_path = "../service/cloudformation/schemas/ec2-instance.json"

  # Whether or not to suppress Terraform resource generation.
  # Optional.
  # The default value is false - A Terraform resource is generated.
  suppress_resource_generation = true

  # Whether or not to suppress Terraform singular data source generation.
  # Optional.
  # The default value is false - A Terraform singular data source is generated.
  suppress_singular_data_source_generation = true

  # Whether or not to suppress Terraform plural data source generation.
  # Optional.
  # The default value is false - A Terraform plural data source is generated.
  suppress_plural_data_source_generation = true
}
```
