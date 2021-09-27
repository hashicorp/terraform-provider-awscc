# aws-sdk-go-base

An opinionated [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2) library for consistent authentication configuration between projects plus additional helper functions. This library was originally started in [HashiCorp Terraform](https://github.com/hashicorp/terraform), migrated with the [Terraform AWS Provider](https://github.com/terraform-providers/terraform-provider-aws) during the Terraform 0.10 Core and Provider split, and now is offered as a separate library to allow easier dependency management in the Terraform ecosystem.

**NOTE:** This library is not currently designed or intended for usage outside the [Terraform S3 Backend](https://www.terraform.io/docs/backends/types/s3.html) and the [Terraform AWS Provider](https://www.terraform.io/docs/providers/aws/index.html).

This project publishes two Go modules, `aws-sdk-go-base/v2` and `aws-sdk-go-base/v2/awsv1shim/v2`.
The module `aws-sdk-go-base/v2` returns configuration compatible with the [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2).
In order to assist with migrating large code bases using the AWS SDK for Go v1, the module `aws-sdk-go-base/v2/awsv1shim/v2` takes the AWS SDK for Go v2 configuration and returns configuration for the AWS SDK for Go v1.

## Requirements

* [Go](https://golang.org/doc/install) 1.16

## Development

Running `make test` will test both `aws-sdk-go-base/v2` and `aws-sdk-go-base/v2/awsv1shim/v2`.
To test individual cases, `go test` works as well, but be aware that it only works in the current module.
To test both modules, run:

```sh
$ go test -v ./...
$ cd v2/awsv1shim && go test -v ./...
```

Code is validated using
[`golangci-lint`](https://github.com/golangci/golangci-lint) for general code quality,
[`impi`](https://github.com/pavius/impi) for import block formatting, and
[Semgrep](https://semgrep.dev) to validate additional rules.

`golangci-lint` and `impi` are Go tools, and can be installed using either `make tools` or installing the Go packages.
Installing the packages from the `tools` directory will ensure that you are using the expected version.
`Semgrep` can be installed as described in [the documentation](https://semgrep.dev/docs/getting-started/) or using a Docker container.

Validation can be run using `make lint` to run `golangci-lint` and `impi`.
`make semgrep` will run Semgrep using a Docker container.

If running linters directly, be aware that `golangci-lint` will only run for the current module.
To validate both modules, run:

```sh
$ golangci-lint run ./...
$ cd v2/awsv1shim && golangci-lint run ./...
```

## Release Process

The two modules can be released separately.
If changes are only made to `awsv1shim`, `aws-sdk-go-base` should **not** be released.
However, if changes are made to `aws-sdk-go-base`, both modules should be released.
If `aws-sdk-go-base` is being released, it should be released first, since `awsv1shim` depends on it.

### Releasing `aws-sdk-go-base`

* Push a new version tag to GitHub, in the form `vX.Y.Z`
* Close the associated GitHub milestone
* Create a release on GitHub

### Releasing `awsv1shim`

* If needed, update the version of `aws-sdk-go-base` to the version published above
    * Update the version in the `go.mod` file
    * Run `go mod tidy`
* Push a new version tag to GitHub, in the form `v2/awsv1shim/vX.Y.Z`
* Close the associated GitHub milestone
* Create a release on GitHub
