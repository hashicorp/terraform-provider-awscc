module github.com/hashicorp/terraform-provider-aws-cloudapi

go 1.15

require (
	github.com/aws/aws-sdk-go v1.31.9 // indirect
	github.com/aws/aws-sdk-go-v2 v1.7.1
	github.com/aws/aws-sdk-go-v2/config v1.5.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.7.0
	github.com/evanphx/json-patch v0.5.2 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go v0.2.0
	github.com/hashicorp/go-getter v1.5.3
	github.com/hashicorp/go-hclog v0.16.1 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0
	github.com/hashicorp/terraform-plugin-framework v0.2.0
	github.com/hashicorp/terraform-plugin-go v0.3.1
	github.com/hashicorp/terraform-plugin-log v0.1.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mitchellh/cli v1.1.2
)

replace github.com/aws/aws-sdk-go-v2/service/cloudformation => github.com/hashicorp/aws-sdk-go-v2-service-cloudformation-private v0.0.0-20210723211334-303716fc7a60

replace github.com/hashicorp/terraform-plugin-framework => github.com/ewbankkit/terraform-plugin-framework v0.2.1-0.20210726180547-0f2fe477353e
