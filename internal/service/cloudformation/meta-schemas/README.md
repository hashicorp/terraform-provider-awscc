# AWS CloudFormation Resource Provider Definition Schema

This directory contains the JSON Schema documents which define the [AWS CloudFormation resource _meta-schema_](https://github.com/aws-cloudformation/cloudformation-resource-schema).

After download from GitHub the documents have been modified to add the `file://./` prefix to any [JSON Schema reference](https://json-schema.org/understanding-json-schema/structuring.html#ref) that resolves to a schema definition in this directory.