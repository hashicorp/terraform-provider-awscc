# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

locals {
  workflow_labels = {
    "provider"              = "844fba", # color:terraform (main)
    "needs-triage"          = "dc477d", # color:consul
    "enhancement"           = "844fba", # color:terraform (main)
    "new-resource"          = "8040c9", # color:terraform (link on white)
    "new-data-source"       = "ac72f0", # color:terraform (link on black)
    "bug"                   = "ec585d", # color:boundary
    "crash"                 = "ec585d", # color:boundary
    "breaking-change"       = "ec585d", # color:boundary
    "regression"            = "ec585d", # color:boundary
    "waiting-response"      = "d3353f", # color:darker boundary
    "tests"                 = "60dea9", # color:nomad
    "prerelease-tf-testing" = "60dea9", # color:nomad
    "technical-debt"        = "d1ebff", # color:terraform accent
    "proposal"              = "d1ebff", # color:terraform accent
    "documentation"         = "f4ecff", # color:terraform secondary
    "thinking"              = "f4ecff", # color:terraform secondary
    "question"              = "f4ecff", # color:terraform secondary
    "linter"                = "f4ecff", # color:terraform secondary
    "sweeper"               = "f4ecff", # color:terraform secondary
    "size/XS"               = "62d4dc", # color:lightest-darkest waypoint gradient
    "size/S"                = "4ec3ce", # color:lightest-darkest waypoint gradient
    "size/M"                = "3bb3c0", # color:lightest-darkest waypoint gradient
    "size/L"                = "27a2b2", # color:lightest-darkest waypoint gradient
    "size/XL"               = "1492a4", # color:lightest-darkest waypoint gradient
    "size/XXL"              = "008196", # color:lightest-darkest waypoint gradient
    "upstream-terraform"    = "1c7ada", # color:vagrant
    "upstream"              = "1c7ada", # color:vagrant
    "dependencies"          = "1c7ada", # color:vagrant
    "good first issue"      = "63d0ff", # color:packer
    "help wanted"           = "63d0ff", # color:packer
    "examples"              = "63d0ff", # color:packer
    "stale"                 = "828a90", # color:stale grey
    "new"                   = "828a90", # color:stale grey
    "windows"               = "828a90", # color:stale grey
    "github_actions"        = "828a90", # color:stale grey
    "partner"               = "ff9900", # color:yellow
    "external-maintainer"   = "63d0ff", # color:packer
    "prioritized"           = "d1ebff"  # color:terraform accent
  }
}

resource "github_issue_label" "workflow" {
  for_each = local.workflow_labels

  repository = "terraform-provider-awscc"
  name       = each.key
  color      = each.value
}
