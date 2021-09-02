module github.com/hashicorp/terraform-provider-awscc/tools

go 1.16

require (
	github.com/golangci/golangci-lint v1.42.0
	github.com/hashicorp/terraform-plugin-docs v0.4.1-0.20210901201438-295243212b7f // indirect
	mvdan.cc/gofumpt v0.1.1 // indirect
)

replace github.com/hashicorp/terraform-plugin-docs => github.com/ewbankkit/terraform-plugin-docs v0.4.1-0.20210902155645-4b65db6fb616
