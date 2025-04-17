# Reference to current AWS account
data "aws_caller_identity" "current" {}

# Create an instance of B2BI capability
resource "awscc_b2bi_capability" "example" {
  name = "example-edi-capability"
  type = "edi"

  configuration = {
    edi = {
      capability_direction = "INBOUND"
      transformer_id       = "example-transformer-id"

      input_location = {
        bucket_name = awscc_s3_bucket.input.id
        key         = "input/"
      }

      output_location = {
        bucket_name = awscc_s3_bucket.output.id
        key         = "output/"
      }

      type = {
        x12_details = {
          transaction_set = "X12_850"
          version         = "VERSION_4010"
        }
      }
    }
  }

  instructions_documents = local.instruction_docs

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket" "input" {
  bucket_name    = "b2bi-example-input-${data.aws_caller_identity.current.account_id}"
  access_control = "Private"

  tags = [{
    key   = "Name"
    value = "B2BI Input Bucket"
  }]
}

resource "awscc_s3_bucket" "output" {
  bucket_name    = "b2bi-example-output-${data.aws_caller_identity.current.account_id}"
  access_control = "Private"

  tags = [{
    key   = "Name"
    value = "B2BI Output Bucket"
  }]
}

resource "awscc_s3_bucket" "docs" {
  bucket_name    = "b2bi-example-docs-${data.aws_caller_identity.current.account_id}"
  access_control = "Private"

  tags = [{
    key   = "Name"
    value = "B2BI Documents Bucket"
  }]
}

resource "aws_s3_object" "example_instructions" {
  bucket = awscc_s3_bucket.docs.id
  key    = "instructions/guide.pdf"
  source = "/dev/null" # Create empty file
}

locals {
  instruction_docs = [{
    bucket_name = awscc_s3_bucket.docs.id
    key         = aws_s3_object.example_instructions.key
  }]
}