resource "awscc_rekognition_collection" "example" {
  collection_id = "example"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]

}