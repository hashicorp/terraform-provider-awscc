# A Macie session must exist before creating a findings filter
resource "awscc_macie_session" "example" {
  status                       = "ENABLED"
  finding_publishing_frequency = "ONE_HOUR"
}

# Define a findings filter that automatically archives findings from a specific S3 bucket
resource "awscc_macie_findings_filter" "example" {
  depends_on = [awscc_macie_session.example]

  name        = "example-bucket-filter"
  description = "Automatically archive findings for the example bucket"
  action      = "ARCHIVE"
  position    = 1

  finding_criteria = {
    criterion = {
      "resourcesAffected.s3Bucket.name" = {
        eq = ["example-data-bucket"]
      }
    }
  }

  tags = [
    {
      key   = "Environment"
      value = "Example"
    },
    {
      key   = "Name"
      value = "example-filter"
    }
  ]
}
