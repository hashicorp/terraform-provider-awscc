# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"

  tags = [{
    key = "Name"
  value = "My bucket" }]

}