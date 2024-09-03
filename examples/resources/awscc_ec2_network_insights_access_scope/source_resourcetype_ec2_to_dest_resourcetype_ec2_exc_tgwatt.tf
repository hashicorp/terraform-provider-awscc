resource "awscc_ec2_network_insights_access_scope" "example" {
  match_paths = [{
    source = {
      resource_statement = {
        resource_types = [
          "AWS::EC2::Instance"
        ]
      }
      packet_header_statement = {
        source_addresses = [aws_subnet.source.cidr_block]
        protocols        = ["tcp"]
      }
    }
    destination = {
      resource_statement = {
        resource_types = [
          "AWS::EC2::Instance"
        ]
      }
      packet_header_statement = {
        destination_addresses = [aws_subnet.dest.cidr_block]
        protocols             = ["tcp"]
        destination_ports     = [80]
      }
    }
  }]

  exclude_paths = [{
    through_resources = [{
      resource_statement = {
        resource_types = [
          "AWS::EC2::TransitGatewayAttachment"
        ]
      }
    }]
  }]

  tags = [{
    key   = "Name"
    value = "source-ec2-tcp-to-dest-ec2-tcp-exc-tgw-att"
  }]
}