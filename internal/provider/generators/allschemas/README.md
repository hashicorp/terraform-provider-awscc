# All AWS Cloud Control API Schemas

Generate (to stdout) a sample configuration file for use with the [Schema Downloader](../schema/README.md) that lists all available AWS Cloud Control API schemas.

Note that valid AWS credentials must be available via [standard mechanisms](https://docs.aws.amazon.com/sdkref/latest/guide/environment-variables.html).

Official builds use the schemas from `us-east-1`.

#### Usage

```
export AWS_DEFAULT_REGION=us-east-1
export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
go run internal/provider/generators/allschemas/main.go
```
