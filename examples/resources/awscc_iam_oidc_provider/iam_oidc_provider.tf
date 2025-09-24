data "tls_certificate" "tfc_certificate" {
  url = "https://app.terraform.io"
}

resource "awscc_iam_oidc_provider" "this" {
  thumbprint_list = [data.tls_certificate.tfc_certificate.certificates[0].sha1_fingerprint]
  client_id_list  = ["aws.workload.identity", ]
  url             = data.tls_certificate.tfc_certificate.url
  tags = [{
    key   = "Name"
    value = "IAM OIDC Provider"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    { key   = "Modified By"
      value = "AWSCC"
  }]
}