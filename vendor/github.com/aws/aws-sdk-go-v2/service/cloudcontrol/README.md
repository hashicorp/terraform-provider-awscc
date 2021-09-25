# AWS SDK for Go v2 Cloud Control Service

[AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2)-compatible Cloud Control service client for private feature development.

### Usage

```
go mod edit --replace github.com/aws/aws-sdk-go-v2/service/cloudcontrol=github.com/hashicorp/aws-sdk-go-v2-service-cloudformation-private
GOPRIVATE=github.com/hashicorp/aws-sdk-go-v2-service-cloudformation-private go get -d github.com/aws/aws-sdk-go-v2/service/cloudcontrol
```
