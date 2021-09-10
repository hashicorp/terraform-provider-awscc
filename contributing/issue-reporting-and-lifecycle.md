# Issue Reporting and Lifecycle

<!-- TOC depthFrom:2 -->

- [Issue Reporting Checklists](#issue-reporting-checklists)
    - [Bug Reports](https://github.com/hashicorp/terraform-provider-awscc/issues/new?template=Bug_Report.md)
    - [Feature Requests](https://github.com/hashicorp/terraform-provider-awscc/issues/new?labels=enhancement&template=Feature_Request.md)
    - [Questions](https://github.com/hashicorp/terraform-provider-awscc/issues/new?labels=question&template=Question.md)
- [Issue Lifecycle](#issue-lifecycle)

<!-- /TOC -->

## Issue Reporting Checklists

We welcome issues of all kinds including feature requests, bug reports, and
general questions. Below you'll find checklists with guidelines for well-formed
issues of each type.

### [Bug Reports](https://github.com/hashicorp/terraform-provider-awscc/issues/new?template=Bug_Report.md)

- [ ] __Test against latest release__: Make sure you test against the latest
   released version. It is possible we already fixed the bug you're experiencing.

- [ ] __Search for possible duplicate reports__: It's helpful to keep bug
   reports consolidated to one thread, so do a quick search on existing bug
   reports to check if anybody else has reported the same thing. You can [scope
      searches by the label "bug"](https://github.com/hashicorp/terraform-provider-awscc/issues?q=is%3Aopen+is%3Aissue+label%3Abug) to help narrow things down.

- [ ] __Include steps to reproduce__: Provide steps to reproduce the issue,
   along with your `.tf` files, with secrets removed, so we can try to
   reproduce it. Without this, it makes it much harder to fix the issue.

- [ ] __For panics, include `crash.log`__: If you experienced a panic, please
   create a [gist](https://gist.github.com) of the *entire* generated crash log
   for us to look at. Double check no sensitive items were in the log.

### [Feature Requests](https://github.com/hashicorp/terraform-provider-awscc/issues/new?labels=enhancement&template=Feature_Request.md)

- [ ] __Ensure coverage status__: The generated provider will only support functionality available in the CloudFormation schema provided by AWS. To inflence that prioritization please refer to the [CloudFormation Open Coverage](https://github.com/aws-cloudformation/cloudformation-coverage-roadmap/projects/1) roadmap. At this time we will not accept PR's which add or modify functionality in generated resources. There may however be a bug in the generation code which prevents a resource from being available in the provider. Typically we will be aware of this issue and will track it publically, but if you are seeing an issue in generation that does not have an issue attached please feel free to submit a bug report.

- [ ] __Search for possible duplicate requests__: It's helpful to keep requests
   consolidated to one thread, so do a quick search on existing requests to
   check if anybody else has reported the same thing. You can [scope searches by
      the label "enhancement"](https://github.com/hashicorp/terraform-provider-awscc/issues?q=is%3Aopen+is%3Aissue+label%3Aenhancement) to help narrow things down.

- [ ] __Include a use case description__: In addition to describing the
   behavior of the feature you'd like to see added, it's helpful to also lay
   out the reason why the feature would be important and how it would benefit
   Terraform users.

### [Questions](https://github.com/hashicorp/terraform-provider-awscc/issues/new?labels=question&template=Question.md)

- [ ] __Search for answers in Terraform documentation__: We're happy to answer
   questions in GitHub Issues, but it helps reduce issue churn and maintainer
   workload if you work to [find answers to common questions in the
   documentation](https://www.terraform.io/docs/providers/aws/index.html). Oftentimes Question issues result in documentation updates
   to help future users, so if you don't find an answer, you can give us
   pointers for where you'd expect to see it in the docs.

## Issue Lifecycle

1. The issue is reported.

2. The issue is verified and categorized by a Terraform collaborator.
   Categorization is done via GitHub labels. We generally use a two-label
   system of (1) issue/PR type, and (2) section of the codebase. Type is
   one of "bug", "enhancement", "documentation", or "question", and section
   is usually the AWS service name.

3. An initial triage process determines whether the issue is critical and must
    be addressed immediately, or can be left open for community discussion.

4. The issue is addressed in a pull request or commit. The issue number will be
   referenced in the commit message so that the code that fixes it is clearly
   linked.

5. The issue is closed. Sometimes, valid issues will be closed because they are
   tracked elsewhere or non-actionable. The issue is still indexed and
   available for future viewers, or can be re-opened if necessary.
