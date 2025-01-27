# Terraform AWS Cloud Control Provider Resource Behavior

This document describes the behavior of resources implemented by the provider. In particular it describes in detail how [AWS CloudFormation resource types](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html) are defined as [Terraform resources](https://developer.hashicorp.com/terraform/language/resources).

We assume some familiarity with user-facing Terraform concepts like configuration, state, CLI workflow, etc. The Terraform website has documentation on these ideas.

## Resources

### Terraform Resources

Resources describe infrastructure objects such as virtual networks or compute instances.
By declaring a [Terraform resource block](https://developer.hashicorp.com/terraform/language/resources/syntax) and [applying that configuration](https://developer.hashicorp.com/terraform/language/resources/behavior), Terraform manages the lifecycle of the underlying infrastructure object so as to make its settings match the configuration.

Resource blocks declare a resource of a specific type with a specific local name. The local name is used solely to refer to that resource within its own [module](https://developer.hashicorp.com/terraform/language/modules), having no meaning outside the module's scope. For example

```terraform
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"
}
```

declares a resource of type `awscc_s3_bucket` with its `bucket_name` argument set to `"example-bucket"`. The resource's module-local name is `example`.

Resources are implemented in [providers](https://developer.hashicorp.com/terraform/language/providers), plugins which interact with infrastructure providers such as AWS.
A resource's [implementation](https://developer.hashicorp.com/terraform/plugin/framework/resources) defines a [schema](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas) which describes the resource's arguments, and methods (including CRUD methods) which define the resource's lifecycle management functionality.

Note that we will use the terms argument and attribute interchangeably from now on as Terraform plugins use [_attribute_](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes) and Terraform CLI uses [_argument_](https://developer.hashicorp.com/terraform/language/syntax/configuration#arguments) for the same concept.

### Terraform Data Sources

[Data sources](https://developer.hashicorp.com/terraform/language/data-sources) are a variant of resource intended to allow Terraform to reference external data. Unlike [managed resources](#terraform-resources), Terraform does not manage the underlying infrastructure object's lifecycle. Data sources are intended to have no side-effects.

For the purposes of this document we consider data sources to be similar to resources with only a Read method. We will call out differences where they are significant.

### AWS CloudFormation Resources

[CloudFormation resources](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html) are conceptually very similar to Terraform resources. However, there are some differences:

* [Resource type schemas](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html) are defined using [JSON Schema dialect](https://github.com/aws-cloudformation/cloudformation-resource-schema) and published in the [CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html).
* Resources are implemented via CRUDL handlers, responsible for interacting with underlying AWS (or third-party) services to manage infrastructure lifecycle. Handlers are invoked either via a [CloudFormation stack](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-whatis-concepts.html#cfn-concepts-stacks) or the [Cloud Control API](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/what-is-cloudcontrolapi.html).

## Interpretation Of CloudFormation Resource Schemas

During [generation](./generating.md) of the Terraform AWS Cloud Control Provider, all available CloudFormation resource schemas are downloaded from the CloudFormation registry and are cached in this GitHub repository (so as to have reproducible builds).
Unless suppressed, each CloudFormation resource schema is then used to generate

* A Terraform resource.
* A Terraform singular data source. A singular data source returns attributes of a single AWS object. A unique identifier is used to specify for which AWS object information is returned.
* A plural data source. A plural data source returns a list of the unique identifiers for every AWS object of the resource's type (in the [configured](https://registry.terraform.io/providers/hashicorp/awscc/latest/docs#schema) AWS account and Region).

### Resource Type Naming

CloudFormation resource [type names](https://github.com/aws-cloudformation/cloudformation-resource-schema?tab=readme-ov-file#resource-type-name) consist of three parts; an organization, service and resource (for example `AWS::EC2::Instance`).
Terraform type names are derived from the CloudFormation type name by lower casing the service part, [snake casing](https://en.wikipedia.org/wiki/Snake_case) the resource part and using `awscc_` as a prefix. The resource part is pluralized for any plural data source type name.

For example, the `AWS::EC2::Instance` CloudFormation resource type leads to the generation of the `awscc_ec2_instance` resource and `awscc_ec2_instance` and `awscc_ec2_instances` data sources.

### Resource Shape

The _shape_ of a resource defines the names, types and behaviors of its fields. Every [property](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-properties) in a CloudFormation resource schema corresponds to an argument in the Terraform schema.

#### Attribute Naming

A Terraform attribute's name is obtained by snake casing the corresponding CloudFormation property's name. For example a property named `GlobalReplicationGroupDescription` corresponds to an attribute named `global_replication_group_description`.

> [!NOTE]
> If a top-level attribute's name is one of the Terraform _meta-arguments_ [`count`](https://developer.hashicorp.com/terraform/language/meta-arguments/count), [`depends_on`](https://developer.hashicorp.com/terraform/language/meta-arguments/depends_on), [`for_each`](https://developer.hashicorp.com/terraform/language/meta-arguments/for_each), or [`lifecycle`](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) then generation of the Terraform resource (and data sources) is suppressed.
>
> If a top-level attribute's name is [`provider`](https://developer.hashicorp.com/terraform/language/meta-arguments/resource-provider) it is renamed to `provider_name`.

#### Attribute Types

A Terraform attribute's [type](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes#available-attribute-types) is derived from the corresponding CloudFormation property's [type](https://json-schema.org/understanding-json-schema/reference/type).

| CloudFormation Type | Terraform Type |
|---------------------|----------------|
| [`boolean`](https://json-schema.org/understanding-json-schema/reference/boolean) | [`Bool`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/bool) |
| [`integer`](https://json-schema.org/understanding-json-schema/reference/numeric#integer) | [`Int64`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/int64) |
| [`number`](https://json-schema.org/understanding-json-schema/reference/numeric#number) | [`Float64`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/float64) |
| [`string`](https://json-schema.org/understanding-json-schema/reference/string) | [`String`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/string)[^1] |
| [`array`](https://json-schema.org/understanding-json-schema/reference/array) | [`List`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/list) or [`Set`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/set)[^2][^3] |
| [`object`](https://json-schema.org/understanding-json-schema/reference/object) | [`Nested attribute`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes#nested-attribute-types) or [`Map`](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes/map)[^4] |

##### Array Types

| insertionOrder | uniqueItems | Terraform Type |
|----------------|-------------|----------------|
| `true` | `false` | `List` |
| `false` | `false` | `List` custom type with [semantic equality](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/custom#semantic-equality) |
| `true` | `true` | `List` with [`UniqueValues`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/listvalidator#UniqueValues) [validator](https://developer.hashicorp.com/terraform/plugin/framework/validation#attribute-validation) |
| `false` | `true` | `Set` |

[^1]: JSON Schema string properties with a [`format`](https://json-schema.org/understanding-json-schema/reference/string#format) value of [`"date-time"`](https://json-schema.org/understanding-json-schema/reference/string#dates-and-times) correspond to the Terraform [`RFC3339`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes#RFC3339) [custom type](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/custom).

[^2]: JSON Schema array properties correspond to either Terraform lists or sets depending on the values of [`uniqueItems`](https://json-schema.org/understanding-json-schema/reference/array#uniqueItems) and [`insertionOrder`](https://github.com/aws-cloudformation/cloudformation-resource-schema?tab=readme-ov-file#insertionorder).

[^3]: An array's [item type](https://json-schema.org/understanding-json-schema/reference/array#items) determines the Terraform list or set element type. See [Array Types](#array-types).

 [^4]: JSON Schema object properties with [pattern properties](https://json-schema.org/understanding-json-schema/reference/object#patternProperties) correspond to Terraform maps. Only the first pattern is considered.

#### Attribute Validation

##### Integer Validation

A JSON Schema integer property's [`minimum` and `maximum`](https://json-schema.org/understanding-json-schema/reference/numeric#range) values correspond to Terraform [`AtLeast`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/int64validator#AtLeast), [`AtMost`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/int64validator#AtMost) and [`Between`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/int64validator#Between) validators.

Any [`enum`](https://json-schema.org/understanding-json-schema/reference/enum) value corresponds to the Terraform [`OneOf`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/int64validator#OneOf) validator.

##### Number Validation

A JSON Schema number property's [`minimum` and `maximum`](https://json-schema.org/understanding-json-schema/reference/numeric#range) values correspond to Terraform [`AtLeast`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/float64validator#AtLeast), [`AtMost`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/float64validator#AtMost) and [`Between`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/float64validator#Between) validators.

##### String Validation

A JSON Schema string property's [`minLength` and `maxLength`](https://json-schema.org/understanding-json-schema/reference/string#length) values correspond to Terraform [`LengthAtLeast`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator#LengthAtLeast), [`LengthAtMost`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator#LengthAtMost) and [`LengthBetween`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator#LengthBetween) validators.

Any [`enum`](https://json-schema.org/understanding-json-schema/reference/enum) value corresponds to the Terraform [`OneOf`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator#OneOf) validator.

Any [`pattern`](https://json-schema.org/understanding-json-schema/reference/string#regexp) value corresponds to the Terraform [`RegexMatches`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator#RegexMatches) validator. If the pattern value is valid for [ECMA-262](https://262.ecma-international.org/) but not for [Go](https://github.com/google/re2/wiki/Syntax) then an empty pattern (`""`) is used in the validator, effectively allowing any string.

##### Array Validation

A JSON Schema array property's [`minItems` and `maxItems`](https://json-schema.org/understanding-json-schema/reference/array#length) values correspond to Terraform [`SizeAtLeast`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/listvalidator#SizeAtLeast), [`SizeAtMost`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/listvalidator#SizeAtMost) and [`SizeBetween`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/listvalidator#SizeBetween) validators (or their [equivalents for sets](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-validators/setvalidator)).

#### Attribute Behaviors

##### Default Values

A CloudFormation property's [`default`](https://json-schema.org/understanding-json-schema/reference/annotations) value corresponds to a Terraform attribute [plan modifier](https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification) which tailors the plan so that if the planned value is [`null`](https://developer.hashicorp.com/terraform/language/expressions/types#null) and there is a current value and the current value is the default then use the current value, else use the planned value.

##### Configurability

A Terraform attribute's _configurability_ defines how Terraform expects data to be set, whether from [configuration](https://developer.hashicorp.com/terraform/language/syntax/configuration) or in the provider's logic (such as an API response value). At least one of three schema attribute flags must be set to `true`:

* `Required`: The attribute must be configured to a [known](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/terraform-concepts#unknown-values), non-`null` value.
* `Optional`: The attribute may be configured to a known value or its value is `null`.
* `Computed`: The attribute's planned value is unknown and a known value must be set by provider logic.

The allowed combinations of these flags are `Required`-only, `Optional`-only, `Computed`-only (no value is allowed to be configured), and `Optional`+`Computed` (a value may be configured; if the configured value is `null`, a value may be set by provider logic).

A Terraform attribute's configurability is derived from the CloudFormation resource's [semantic properties](https://github.com/aws-cloudformation/cloudformation-resource-schema/blob/master/README.md#resource-semantics).

* If a CloudFormation property is [`required`](https://json-schema.org/understanding-json-schema/reference/object#required), the attribute is `Required`.
* If a CloudFormation property is not required and not in the `readOnlyProperties` list, the attribute is `Optional`.
* If a CloudFormation property is in the `readOnlyProperties` list, the attribute is `Computed`.
* If a CloudFormation property has a [default value](#default-values), the attribute is `Computed`. A required property with a default value is switches the attribute to `Optional`.
* All `Optional` attributes are marked as `Computed`. This is because CloudFormation only determines [drift](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html#what-is-drift) for property values that are explicitly set, whereas Terraform expects the value of an unset, non-`Computed` attribute to always be `null` (not present). AWS services will often return values that have not been specified as default values in the CloudFormation resource type schema for properties that are unset in configuration.

##### Immutability

If a CloudFormation property is in the `createOnlyProperties` list, the corresponding Terraform attribute is immutable. If the value of the attribute changes, in-place update is not possible and instead the resource is replaced for the change to occur.
The Terraform [`RequiresReplace`](https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#requiresreplace) plan modified is used for this behavior.

#### The `id` Attribute

Every Terraform schema generated from a CloudFormation resource schema includes a top-level attribute named `id`. This attribute's value uniquely identifies the underlying AWS resource in the configured AWS account and Region and is used for [Cloud Control API operations](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html) and in [acceptance tests](https://developer.hashicorp.com/terraform/plugin/framework/acctests#no-id-found-in-attributes). The `id` attribute is added during the Terraform resource generation process.

> [!NOTE]
> If the CloudFormation resource schema does define a top-level `Id` property then that property is mapped to a Terraform attribute named `<type>_id` (e.g. `flow_log_id` for the `awscc_ec2_flow_log` resource).

## Interaction With The Cloud Control API

The AWS Cloud Control API provides a standardized set of operations to create, read, update, delete, and list (CRUD-L) supported resources in an AWS account.

A _resource type_ represents an infrastructure artifact whose lifecycle can be managed through the API. Each resource type is defined by its _resource type schema_, a JSON Schema-compliant document published to the AWS CloudFormation registry.

To create or update a resource using the Cloud Control API a program must specify a JSON document representing the resource's properties and their values.
When the resource is being created, these values are the resource's _desired state_.
When the resource is being updated, the values are a [list of patches](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html#resource-operations-update-patch) between the resource's current state and its desired state.

Reading a resource returns a JSON document representing its _current state_.

The Terraform AWS Cloud Control Provider maps Terraform data to and from these JSON documents and call Cloud Control API methods.

### Create

The provider's [`Create`](https://developer.hashicorp.com/terraform/plugin/framework/resources/create) method is called from [`terraform apply`](https://developer.hashicorp.com/terraform/cli/commands/apply) when Terraform determines during planning that no resource with the configured module-local name exists or that an existing resource must be recreated. The provider converts the Terraform plan, which describes the expected values of the resource's attributes, into the Cloud Control desired state JSON document -- in the process reversing the snake casing [described above](#attribute-naming).

The resource's desired state document and the CloudFormation resource type name are passed to the [`CreateResource` API](https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_CreateResource.html) and the provider then [polls](https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html) for [completion of the operation](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html).

If the operation fails, the [error is returned](https://developer.hashicorp.com/terraform/plugin/framework/diagnostics) to the Terraform CLI.

If the operation succeeds, the resource's current state document is used to populate [unknown](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/terraform-concepts#unknown-values) attribute values, including the [`id attribute`](#the-id-attribute), and all attribute values are saved in [Terraform state](https://developer.hashicorp.com/terraform/language/state).

### Read

The provider's [`Read`](https://developer.hashicorp.com/terraform/plugin/framework/resources/read) method is called from [`terraform plan`](https://developer.hashicorp.com/terraform/cli/commands/plan), `terraform apply` and [`terraform refresh`](https://developer.hashicorp.com/terraform/cli/commands/refresh) to obtain the current state of the resource. The provider uses the value of the `id` attribute, stored in the resource's prior state, and the CloudFormation resource type name to call the [`GetResource` API](https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResource.html).

If the operation fails, the error is returned to the Terraform CLI.

If the operation succeeds, the resource's current state document is used to populate the Terraform current state. The values of attributes corresponding to properties in the [`writeOnlyProperties` list](https://github.com/aws-cloudformation/cloudformation-resource-schema?tab=readme-ov-file#resource-semantics) are copied from prior state to current state.

### Update

The provider's [`Update`](https://developer.hashicorp.com/terraform/plugin/framework/resources/update) method is called from `terraform apply` when Terraform determines during planning that an existing resource must be updated in-place. The provider converts the resource's prior state and planned new state into JSON documents and generates a [JSON Patch](https://datatracker.ietf.org/doc/html/rfc6902) document listing the operations required to reach the desired state from the current state of the resource.

The patch document and the CloudFormation resource type name are passed to the [`UpdateResource` API](https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_UpdateResource.html) and the provider then polls for completion of the operation.

If the operation fails, the error is returned.

If the operation succeeds, the resource's current state document is used to populate unknown attribute values and all attribute values are saved in Terraform state.

### Delete

The provider's [`Delete`](https://developer.hashicorp.com/terraform/plugin/framework/resources/delete) method is called from `terraform apply` when Terraform determines during planning that an existing resource has been removed from configuration or that an existing resource must be recreated.  The provider uses the value of the `id` attribute and the CloudFormation resource type name to call the [`DeleteResource` API](https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_DeleteResource.html) and then polls for completion of the operation.

If the operation fails, the error is returned to the Terraform CLI and Terraform keeps the resource under management.

If the operation succeeds, the resource is removed from state.

### List

The provider has no `List` method. Instead when a [plural data source's](#interpretation-of-cloudformation-resource-schemas) `Read` method is called the provider uses the CloudFormation resource type name to call the [`ListResources` API](https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_ListResources.html).

If the operation fails, the error is returned to the Terraform CLI.

If the operation succeeds, a list of primary identifiers for all resources is returned and stored in state as the value of the `ids` attribute.
