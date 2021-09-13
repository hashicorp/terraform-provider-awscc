<!--- See what makes a good Pull Request at : https://github.com/hashicorp/terraform-provider-awscc/blob/main/docs/CONTRIBUTING.md --->

<!--- Please keep this note for the community --->

### Community Note

* Please vote on this pull request by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) to the original pull request comment to help the community and maintainers prioritize this request
* Please do not leave "+1" or other comments that do not add relevant new information or questions, they generate extra noise for pull request followers and do not help prioritize the request
* The resources and data sources in this provider are generated from the CloudFormation schema, so they can only support the actions that the underlying schema supports. For this reason submitted bugs should be limited to defects in the generation and runtime code of the provider. Customizing behavior of the resource, or noting a gap in behavior are not valid bugs and should be submitted as enhancements to AWS via the Cloudformation Open Coverage Roadmap.

<!--- Thank you for keeping this note for the community --->

<!--- If your PR fully resolves and should automatically close the linked issue, use Closes. Otherwise, use Relates --->
Relates OR Closes #0000
