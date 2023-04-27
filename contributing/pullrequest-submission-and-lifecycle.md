# Pull Request Submission and Lifecycle

- [Pull Request Lifecycle](#pull-request-lifecycle)
- [Branch Prefixes](#branch-prefixes)
- [Common Review Items](#common-review-items)
    - [Go Coding Style](#go-coding-style)

We appreciate direct contributions to the provider codebase. Here's what to
expect:

- For pull requests that follow the guidelines, we will proceed to reviewing
  and merging, following the provider team's review schedule. There may be some
  internal or community discussion needed before we can complete this.
- Pull requests that don't follow the guidelines will be commented with what
  they're missing. The person who submits the pull request or another community
  member will need to address those requests before they move forward.

## Pull Request Lifecycle

1. [Fork the GitHub repository](https://docs.github.com/en/get-started/quickstart/fork-a-repo),
   modify the code, and [create a pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork).
   You are welcome to submit your pull request for commentary or review before
   it is fully completed by creating a [draft pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests#draft-pull-requests)
   or adding `[WIP]` to the beginning of the pull request title.
   Please include specific questions or items you'd like feedback on.

1. Once you believe your pull request is ready to be reviewed, ensure the
   pull request is not a draft pull request by [marking it ready for review](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/changing-the-stage-of-a-pull-request)
   or removing `[WIP]` from the pull request title if necessary, and a
   maintainer will review it.

1. One of Terraform's provider team members will look over your contribution and
   either approve it or provide comments letting you know if there is anything
   left to do. We'll try give you the opportunity to make the required changes yourself, but in some cases we may perform the changes ourselves if it makes sense to (minor changes, or for urgent issues).  We do our best to keep up with the volume of PRs waiting for
   review, but it may take some time depending on the complexity of the work.

1. Once all outstanding comments have been addressed, your
   contribution will be merged! Merged PRs will be included in the next
   Terraform release.

1. In some cases, we might decide that a PR should be closed without merging.
   We'll make sure to provide clear reasoning when this happens.

## Branch Prefixes

We try to use a common set of branch name prefixes when submitting pull requests. Prefixes give us an idea of what the branch is for. For example, `td-staticcheck-st1008` would let us know the branch was created to fix a technical debt issue, and `d-improve-contributing-guide` would indicate enhancements to the contributing guide. These are the prefixes we currently use:

- f = feature
- b = bug fix
- d = documentation
- t = tests
- td = technical debt
- v = dependencies

Conventions across non-AWS providers varies so when working with other providers please check the names of previously created branches and conform to their standard practices.

## Common Review Items

The Terraform AWS Provider follows common practices to ensure consistent and
reliable implementations across all resources in the project. While there may be
older resource and testing code that predates these guidelines, new submissions
are generally expected to adhere to these items to maintain Terraform Provider
quality. For any guidelines listed, contributors are encouraged to ask any
questions and community reviewers are encouraged to provide review suggestions
based on these guidelines to speed up the review and merge process.

### Go Coding Style

All Go code is automatically checked for compliance with various linters, such as `gofmt`. These tools can be installed using the `GNUMakefile` in this repository.

```console
% cd terraform-provider-awscc
% make tools
```

Check your code with the linters:

```console
% make lint
```

`gofmt` will also fix many simple formatting issues for you. The Makefile includes a target for this:

```console
% make fmt
```

The import statement in a Go file follows these rules:

1. Import declarations are grouped into a maximum of three groups with the following order:
    - Standard packages (also called short import path or built-in packages)
    - Third-party packages (also called long import path packages)
    - Local packages
1. Groups are separated by a single blank line
1. Packages within each group are alphabetized

Check your imports:

```console
% make importlint
```

For greater detail, the following Go language resources provide common coding preferences that may be referenced during review, if not automatically handled by the project's linting tools.

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
