resource "awscc_location_map" "example" {
  map_name = "example"

  configuration = {
    style = "VectorHereExplore"
  }
}