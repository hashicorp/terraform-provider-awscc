---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_neptunegraph_private_graph_endpoint Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::NeptuneGraph::PrivateGraphEndpoint
---

# awscc_neptunegraph_private_graph_endpoint (Data Source)

Data Source schema for AWS::NeptuneGraph::PrivateGraphEndpoint



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `graph_identifier` (String) The auto-generated Graph Id assigned by the service.
- `private_graph_endpoint_identifier` (String) PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.

 For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
- `security_group_ids` (List of String) The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
- `subnet_ids` (List of String) The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
- `vpc_endpoint_id` (String) VPC endpoint that provides a private connection between the Graph and specified VPC.
- `vpc_id` (String) The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
