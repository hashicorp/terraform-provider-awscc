# Route53 CIDR Collection
resource "awscc_route53_cidr_collection" "test" {
  name = "example-cidr-collection"

  locations = [
    {
      location_name = "west-coast"
      cidr_list     = ["10.0.0.0/16", "172.16.0.0/16"]
    },
    {
      location_name = "east-coast"
      cidr_list     = ["192.168.0.0/16", "172.17.0.0/16"]
    }
  ]
}