module github.com/hashicorp/terraform-provider-aws-cloudapi

go 1.15

require (
	github.com/aws/aws-sdk-go v1.38.37
	github.com/evanphx/json-patch v0.5.2 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.1.0
	github.com/hashicorp/aws-sdk-go-base v0.7.1
	github.com/hashicorp/go-getter v1.5.3
	github.com/hashicorp/hcl/v2 v2.10.0
	github.com/hashicorp/terraform-plugin-framework v0.1.0
	github.com/hashicorp/terraform-plugin-go v0.3.1
	github.com/hashicorp/terraform-plugin-log v0.1.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/iancoleman/strcase v0.1.3
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mitchellh/cli v1.1.2
)

replace github.com/aws/aws-sdk-go => github.com/hashicorp/aws-sdk-go-private v1.38.23-0.20210420184552-ae24b9862457
