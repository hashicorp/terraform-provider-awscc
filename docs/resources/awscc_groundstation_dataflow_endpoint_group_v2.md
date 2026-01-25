# Ground Station Dataflow Endpoint Group V2

This example creates a Ground Station dataflow endpoint group V2 with uplink and downlink endpoints for satellite communication.

## Example Usage

```hcl
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
```

## Argument Reference

See the [AWS Provider CloudControlApi documentation](https://registry.terraform.io/providers/hashicorp/awscc/latest/docs/resources/groundstation_dataflow_endpoint_group_v2) for detailed argument descriptions.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `dataflow_endpoint_group_v2_id` - The ID of the Ground Station Dataflow Endpoint Group V2
* `arn` - The ARN of the Ground Station Dataflow Endpoint Group V2