# For complete examples for all PCS resources, please refer https://github.com/aws-samples/aws-hpc-recipes/blob/main/recipes/pcs/terraform_awscc/assets/modules/compute/main.tf

resource "awscc_pcs_queue" "example" {
  name       = "normal"
  cluster_id = awscc_pcs_cluster.main.cluster_id
  compute_node_group_configurations = [
    { compute_node_group_id = awscc_pcs_compute_node_group.compute-st.compute_node_group_id },
    { compute_node_group_id = awscc_pcs_compute_node_group.compute-dy.compute_node_group_id }
  ]

  slurm_configuration = {
    slurm_custom_settings = [
      {
        parameter_name  = "Default"
        parameter_value = "YES"
      },
      {
        parameter_name  = "MaxTime"
        parameter_value = "48:00:00"
      },
    ]
  }


}
