# For complete examples for all PCS resources, please refer https://github.com/aws-samples/aws-hpc-recipes/blob/main/recipes/pcs/terraform_awscc/assets/modules/compute/main.tf

resource "awscc_pcs_compute_node_group" "login" {
  name       = "login"
  cluster_id = awscc_pcs_cluster.main.cluster_id
  custom_launch_template = {
    template_id = aws_launch_template.login.id
    version     = aws_launch_template.login.latest_version
  }
  iam_instance_profile_arn = var.instance_profile_arn
  instance_configs = [{
    instance_type = var.pcs_cng_login_instance_type
  }]
  scaling_configuration = {
    min_instance_count = 1
    max_instance_count = 1
  }
  subnet_ids = [var.access_subnet_id]
  ami_id     = var.pcs_cng_ami_id

}
