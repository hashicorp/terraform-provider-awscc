module github.com/hashicorp/terraform-provider-aws-cloudapi

go 1.15

require (
	github.com/aws/aws-sdk-go v1.38.37 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.1.0
	github.com/hashicorp/hcl/v2 v2.10.0
	github.com/hashicorp/terraform-plugin-docs v0.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/iancoleman/strcase v0.1.3
)

replace github.com/aws/aws-sdk-go => github.com/hashicorp/aws-sdk-go-private f-cloudapi-20210420
