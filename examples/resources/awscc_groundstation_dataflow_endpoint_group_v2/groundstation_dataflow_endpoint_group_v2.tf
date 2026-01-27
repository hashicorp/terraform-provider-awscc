resource "awscc_groundstation_dataflow_endpoint_group_v2" "example" {
  endpoints = [{
    downlink_aws_ground_station_agent_endpoint = {
      name = "example-downlink-endpoint"
      dataflow_details = {
        agent_connection_details = {
          agent_ip_and_port_address = {
            socket_address = {
              name = "192.0.2.10"
              port_range = {
                minimum = 55550
                maximum = 55559
              }
            }
            mtu = 1500
          }
          egress_address_and_port = {
            socket_address = {
              name = "192.0.2.20"
              port = 55560
            }
            mtu = 1500
          }
        }
      }
    }
    }, {
    uplink_aws_ground_station_agent_endpoint = {
      name = "example-uplink-endpoint"
      dataflow_details = {
        agent_connection_details = {
          agent_ip_and_port_address = {
            socket_address = {
              name = "192.0.2.30"
              port_range = {
                minimum = 55570
                maximum = 55579
              }
            }
            mtu = 1500
          }
          ingress_address_and_port = {
            socket_address = {
              name = "192.0.2.40"
              port = 55580
            }
            mtu = 1500
          }
        }
      }
    }
  }]

  contact_pre_pass_duration_seconds  = 120
  contact_post_pass_duration_seconds = 60

  tags = [{
    key   = "Environment"
    value = "Example"
    }, {
    key   = "Name"
    value = "example-dataflow-endpoint-group-v2"
  }]
}

output "dataflow_endpoint_group_v2_id" {
  description = "The ID of the Ground Station Dataflow Endpoint Group V2"
  value       = awscc_groundstation_dataflow_endpoint_group_v2.example.dataflow_endpoint_group_v2_id
}

output "dataflow_endpoint_group_v2_arn" {
  description = "The ARN of the Ground Station Dataflow Endpoint Group V2"
  value       = awscc_groundstation_dataflow_endpoint_group_v2.example.arn
}