module github.com/hashicorp/aws-sdk-go-base

require (
	github.com/aws/aws-sdk-go v1.31.9
	github.com/aws/aws-sdk-go-v2 v1.8.0
	github.com/aws/aws-sdk-go-v2/config v1.6.0
	github.com/aws/aws-sdk-go-v2/credentials v1.3.2
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.4.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.6.1
	github.com/aws/smithy-go v1.7.0
	github.com/google/go-cmp v0.5.6
	github.com/hashicorp/go-cleanhttp v0.5.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
)

go 1.16

replace github.com/aws/aws-sdk-go-v2/credentials => github.com/gdavison/aws-sdk-go-v2/credentials v1.2.2-0.20210811194025-146c1ad6c3b2
