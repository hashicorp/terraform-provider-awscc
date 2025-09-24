resource "awscc_emr_security_configuration" "example" {
  name = "example"
  security_configuration = jsonencode({
    "EncryptionConfiguration" : {
      "AtRestEncryptionConfiguration" : {
        "S3EncryptionConfiguration" : {
          "EncryptionMode" : "SSE-S3"
        },
        "LocalDiskEncryptionConfiguration" : {
          "EncryptionKeyProviderType" : "AwsKms",
          "AwsKmsKey" : awscc_kms_key.example.key_id
        }
      },
      "EnableInTransitEncryption" : true,
      "EnableAtRestEncryption" : true,
      "InTransitEncryptionConfiguration" : {
        "TLSCertificateConfiguration" : {
          "CertificateProviderType" : "PEM",
          "S3Object" : "arn:aws:s3:::config_bucket/certs.zip"
        }
      }
    },
    "InstanceMetadataServiceConfiguration" : {
      "MinimumInstanceMetadataServiceVersion" : 2,
      "HttpPutResponseHopLimit" : 1
    }
  })
}


