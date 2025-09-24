resource "awscc_s3express_directory_bucket" "example" {
  bucket_name     = "example-bucket--use1-az4--x-s3"
  data_redundancy = "SingleAvailabilityZone"
  location_name   = "use1-az4"
}