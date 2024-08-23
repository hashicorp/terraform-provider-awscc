
resource "awscc_ec2_network_insights_access_scope" "example" {
  match_paths = [{
    source = {
      resource_statement = {
        resources = [
          aws_vpc.example.id
        ]
      }
    }
    destination = {
      resource_statement = {
        resource_types = [
          "AWS::EC2::InternetGateway"
        ]
      }
    }
  }]

  tags = [{
    key   = "Name"
    value = "source-vpc-id-to-dest-igw"
  }]
}