# Adding Examples to Generated Documentation

The Cloud Control provider is 100% generated from the AWS CloudFormation Schema *including* the documentation. This means that at the resource/attribute level any changes must be made in the CloudFormation Schema and cannot be modified within the provider itself. However examples can be added to the generated documentation in order to provide working examples or best practices when using the provider.

The Terraform AWS Cloud Control provider uses [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) to generate the documentation found in the [Terraform Registry](https://registry.terraform.io/providers/hashicorp/awscc/latest/docs). Full details of the tool can be found on its GitHub repository. In summary, the tool uses the resource schema to generate and populate Terraform Registry compatible documentation.

## Adding a Single Example to a Resource or Data Source

The default template used by [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) will already look for and include an example if an appropriately named file exists in a known path.

To include an example for the `awscc_amplify_app` resource, simply add an example file to the following path: `examples/resources/awscc_amplify_app/resource.tf`

To include an example for the `awscc_amplify_app` data source, simply add an example file to the following path: `examples/data-sources/awscc_amplify_app/data-source.tf`

## Adding Multiple Examples to a Resource or Data Source

In order to include multiple examples, a new template must be provided to configure the format in which the examples should be displayed.

1. Create a new file in `template/resources` or `template/datasources` named the same as the resource/data source you want to add the template for. For example, to add a new template for the `awscc_ssmcontacts_plan` resource, you would create a file named `template/resources/ssm_contacts_plan.md.tmpl`.
2. As a starting point you can view the default template used by `terraform-plugin-go` [here](https://github.com/hashicorp/terraform-plugin-docs/blob/2385169af97e6ac8bb69446e70d4d4d4db74c9fc/internal/provider/template.go#L217)
3. Make any changes to the template you wish (though for consistency most parts of the template should remain unchanged).
4. For each example, add a file in the examples folder: eg. `examples/resources/awscc_ssm_contacts_plan/example-one.tf`
5. Include `{{ tffile (printf "examples/resources/%s/example-one.tf" .Name)}}` so that [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) knows to pull in that examples when the documentation is generated.

## Customize the Position of a Single Example

In order to change the position of a single example, you will need to add a new template file for the resource as described [above](#adding-multiple-examples-to-a-resource-or-data-source). Then move the following stanza to change its location:

```go
{{ if .HasExample -}}
## Example Usage

{{ printf "{{tffile %q}}" .ExampleFile }}
{{- end }}
```

## Generating the Documentation and Raising a Pull Request

1. Run `make docs` to regenerate all of the documentation. You should now be able to see the changed docs by running `git status`.
2. Review the docs to make sure the layout is what you are expecting.
3. Raise a pull request including any new templates/examples and the generated documentation. IMPORTANT: Documentation Generation is expected to be done locally, and changes committed to the repository. Once a PR is is raised, a check will run to compare the source in your branch to the output of the `make docs` command and will fail if they are not identical.
