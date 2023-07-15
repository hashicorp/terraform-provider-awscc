# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "awscc_oam_link" "link_example" {
  resource_types  = ["AWS::CloudWatch::Metric", "AWS::Logs::LogGroup", "AWS::XRay::Trace"]
  sink_identifier = var.sinkArn
  label_template  = "$AccountName"
}