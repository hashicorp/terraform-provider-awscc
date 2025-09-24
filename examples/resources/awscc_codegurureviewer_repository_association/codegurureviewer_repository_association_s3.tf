resource "awscc_codegurureviewer_repository_association" "example2" {
  bucket_name = var.bucket_name
  name        = var.bucket_name
  type        = "S3Bucket"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

variable "bucket_name" {
  type = string
}
