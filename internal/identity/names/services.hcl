# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# This file defines services that exist in AWS.
# Services are inserted into this file when an exception to the normal identity generator process is necessary.

## example service definition with all available options

# service "example" {
#   is_global = true
#
#   # resource is a repeatable block
#   resource "awscc_example_thing2" {
#     has_mutable_identity = true
#   }
#
#   resource "awscc_example_thing1" {
#      has_mutable_identity = true
#   }
# }

service "bcmdataexports" {
  is_global = true
}

service "billing" {
  is_global = true
}

service "billingconductor" {
  is_global = true
}

service "budgets" {
  is_global = true
}

service "ce" {
  is_global = true
}

service "cloudfront" {
  is_global = true
}

service "cur" {
  is_global = true
}

service "globalaccelerator" {
  is_global = true
}

service "iam" {
  is_global = true
}

service "networkmanager" {
  is_global = true
}

service "notifications" {
  is_global = true
}

service "notificationscontacts" {
  is_global = true
}

service "organizations" {
  is_global = true
}

service "rolesanywhere" {
  is_global = true
}

service "route53" {
  is_global = true
}

service "route53recoverycontrol" {
  is_global = true
}

service "route53recoveryreadiness" {
  is_global = true
}

service "shield" {
  is_global = true
}