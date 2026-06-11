<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Generating the Provider

This document describes the process of generating the Terraform AWS Cloud Control Provider from [CloudFormation resource type schemas](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html).

<!--mdtoc: begin-->
* [1. Setup](#1-setup)
* [2. Schema Refresh](#2-schema-refresh)
* [3. Track New Schemas](#3-track-new-schemas)
* [4. Generate Resources](#4-generate-resources)
* [5. Generate Data Sources](#5-generate-data-sources)
* [6. Build](#6-build)
* [7. Smoke Test](#7-smoke-test)
* [8. Generate Documentation](#8-generate-documentation)
* [9. Update CHANGELOG](#9-update-changelog)
* [10. Open Pull Request](#10-open-pull-request)
<!--mdtoc: end-->

## 1. Setup

Pull the latest from the GitHub repository, and update the tools:

```sh
git switch main
git pull
```

```sh
make tools
```

Create a new branch:

```sh
make newbranch
```

> [!TIP]
> `date -I` formats the date using `YYYY-MM-DD`.

The CloudFormation resource type schemas in the `us-east-1` region are used to generate the provider.
Ensure that this is the configured default AWS Region:

```sh
export AWS_DEFAULT_REGION=us-east-1
```

Also ensure that valid AWS credentials are set in the environment.

## 2. Schema Refresh

Delete all existing resource schemas:

```sh
make cleanschemas
```

Currently updates to some schemas must be suppressed as they have changes which prevent Terraform schema generation (or they no longer exist and are pending major version removal):

```sh
make suppressions
```

> [!NOTE]
> When adding files to the list (`internal/update/suppressions_checkout.txt`), add one file per line.
> Be sure to open a GitHub issue with details any time a new entry is added to the list.

Download the current versions of the schemas.

```sh
make schemas
```

This operation uses the CloudFormation API and may be rate-limited.
**Repeat until there are no errors.**
If `git status` shows that schemas have been deleted then there were errors during download.

> [!NOTE]
> When an update to an existing schema causes an new failure, the schema may need to be added to the suppression list.
> Once added, re-run the suppression step above to ensure schema retrieval completes without issue.

Commit the updated schemas:

```sh
make commitrefresh
```

## 3. Track New Schemas

Generate a list of all currently available CloudFormation resource type schemas (this may take `~11` minutes):

```sh
make biglister
```

> [!TIP]
> There may be a bit of output to `stdout` during this generation.
> It can be ignored as it pertains to the current version of the schema (e.g. `Quicksight::Analysis`) and we are using the previous version for the resource generation.

Compare this latest schemas with the previous list.
The experimental command below attempts handle this automatically.

```sh
make bigdiffer
```

If this fails, manually enter the date of the last release in a `diff` command.
For example,

```sh
diff internal/provider/generators/allschemas/available_schemas.2021-11-17.hcl internal/provider/generators/allschemas/available_schemas.$(date -I).hcl
```

The result will look similar to the following.
Manually add each new resource type to `internal/provider/all_schemas.hcl`.

> [!TIP]
> Angle-bracket direction matters.
> The usual marker is right-pointing (`>`), which means the resource was added.
> Occasionally, you may see left-pointing (`<`), as below, which means the resource was removed.

```text
1c1
< # 1359 CloudFormation resource types schemas are available for use with the Cloud Control API.
---
> # 1368 CloudFormation resource types schemas are available for use with the Cloud Control API
248a249,253
> suppress_plural_data_source_generation = true
> }
>
> resource_schema "aws_apigatewayv2_stage" {
> cloudformation_type_name               = "AWS::ApiGatewayV2::Stage"
372a378,381
> resource_schema "aws_appstream_stack" {
> cloudformation_type_name = "AWS::AppStream::Stack"
> }
>
1424a1434,1438
> suppress_plural_data_source_generation = true
> }
>
> resource_schema "aws_customerprofiles_recommender" {
> cloudformation_type_name               = "AWS::CustomerProfiles::Recommender"
2317,2320d2330
< resource_schema "aws_emr_cluster" {
<   cloudformation_type_name = "AWS::EMR::Cluster"
< }
<
4022a4033,4036
> resource_schema "aws_novaact_workflow_definition" {
> cloudformation_type_name = "AWS::NovaAct::WorkflowDefinition"
> }
>
4082a4097,4100
> resource_schema "aws_omics_configuration" {
> cloudformation_type_name = "AWS::Omics::Configuration"
> }
>
5154a5173,5176
> }
>
> resource_schema "aws_sagemaker_model" {
> cloudformation_type_name = "AWS::SageMaker::Model"
5234a5257,5264
> }
>
> resource_schema "aws_securityagent_agent_space" {
> cloudformation_type_name = "AWS::SecurityAgent::AgentSpace"
> }
>
> resource_schema "aws_securityagent_application" {
> cloudformation_type_name = "AWS::SecurityAgent::Application"
5236a5267,5275
> resource_schema "aws_securityagent_pentest" {
> cloudformation_type_name               = "AWS::SecurityAgent::Pentest"
> suppress_plural_data_source_generation = true
> }
>
> resource_schema "aws_securityagent_target_domain" {
> cloudformation_type_name = "AWS::SecurityAgent::TargetDomain"
> }
>
```

Once the new entries are included in `all_schemas.hcl`, re-run the `schemas` Make target to download any new resource type schemas:

```sh
make schemas
```

Downloading the schemas also validates them.
If there are any errors, suppress resource generation in `all_schemas.hcl` (`suppress_resource_generation = true`) and open an issue with details about the suppression reason and relevant sections of the resource schema.
See an example issue [here](https://github.com/hashicorp/terraform-provider-awscc/issues/2070).

Commit the new schemas:

```sh
make commitschemas
```

## 4. Generate Resources

Generate all the Terraform resource schemas:

```sh
make resources
```

If there are any errors, add `suppress_resource_generation = true` to the relevant block in `all_schemas.hcl` and re-run the command above.

Commit the Terraform resource schemas:

```sh
make commitresources
```

## 5. Generate Data Sources

Generate all the Terraform data source schemas:

```sh
make singular-data-sources plural-data-sources
```

If there are any errors, add `suppress_plural_data_source_generation = true` and/or `suppress_singular_data_source_generation = true` to the relevant block in `all_schemas.hcl` and re-run the command above.

You may also need to re-run `make schemas`.

Commit the Terraform data source schemas:

```sh
make commitdatas
```

## 6. Build

Ensure that the provider builds:

```sh
make build
```

## 7. Smoke Test

Run a smoke test (all tests should pass):

```sh
make smoke
```

To manually run smoke tests instead:

```sh
% TF_LOG=ERROR make testacc PKG_NAME=internal/aws/logs TESTARGS='-run=TestAccAWSLogsLogGroup_\|TestAccAWSLogsLogGroupDataSource_' ACCTEST_PARALLELISM=3
```

Something similar to this is an expected error:

```text
2025-08-14T16:23:46.085-0400 [ERROR] sdk.proto: Response contains error diagnostic: tf_rpc=ReadDataSource tf_data_source_type=awscc_logs_log_group diagnostic_detail="After attempting to read the data source, the API returned a resource not found error for the id provided. Original Error: couldn't find resource" diagnostic_severity=ERROR diagnostic_summary="AWS Data Source Not Found" tf_req_id=8d016cb0-8616-bfc8-6f5f-0de67318a3b4 tf_provider_addr=registry.terraform.io/hashicorp/awscc tf_proto_version=6.9
```

## 8. Generate Documentation

Build the documentation.
Ensure that you are using Terraform `v1.14` or above for list resource support.

```sh
make docs-all
```

## 9. Update CHANGELOG

`CHANGELOG` entries can be generated from the names of the new `*.md` files.
New files can be listed using:

```sh
git ls-files --others --exclude-standard
```

Manually copy these into `CHANGELOG.md` and format appropriately.
`version/VERSION` should also be updated to match the `CHANGELOG`.

Finally, commit the documentation:

```sh
make commitdocs
```

## 10. Open Pull Request

Once all steps are complete, open a pull request any verify all CI checks pass.
Once merged a new release can be completed.
