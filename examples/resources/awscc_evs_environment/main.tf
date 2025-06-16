# Create a VPC for the EVS environment
resource "awscc_ec2_vpc" "evs" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "evs-vpc"
  }]
}

# Create a subnet for service access
resource "awscc_ec2_subnet" "service_access" {
  vpc_id     = awscc_ec2_vpc.evs.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "evs-service-access"
  }]
}

# Create a security group for EVS service access
resource "awscc_ec2_security_group" "evs_service" {
  group_name        = "evs-service-access"
  group_description = "Security group for EVS service access"
  vpc_id            = awscc_ec2_vpc.evs.id
  security_group_ingress = [{
    from_port   = 443
    to_port     = 443
    ip_protocol = "tcp"
    cidr_ip     = "0.0.0.0/0"
  }]
  security_group_egress = [{
    from_port   = -1
    to_port     = -1
    ip_protocol = "-1"
    cidr_ip     = "0.0.0.0/0"
  }]
  tags = [{
    key   = "Name"
    value = "evs-service-sg"
  }]
}

# Create an SSH key pair for the hosts
resource "awscc_ec2_key_pair" "evs_hosts" {
  key_name = "evs-hosts-key"
  tags = [{
    key   = "Name"
    value = "evs-hosts-key"
  }]
}

# Create the EVS Environment
resource "awscc_evs_environment" "example" {
  environment_name = "example-evs"
  site_id          = "examplesite"
  vpc_id           = awscc_ec2_vpc.evs.id
  vcf_version      = "VCF-5.2.1"

  service_access_subnet_id = awscc_ec2_subnet.service_access.id
  terms_accepted           = true

  connectivity_info = {
    private_route_server_peerings = ["10.0.0.1", "10.0.0.2"]
  }

  license_info = {
    solution_key = "ABCD1-EFGH2-IJKL3-MNOP4-QRST5"
    vsan_key     = "VSAN1-VSAN2-VSAN3-VSAN4-VSAN5"
  }

  vcf_hostnames = {
    cloud_builder = "cloudbuilder"
    nsx           = "nsx"
    nsx_edge_1    = "nsxedge1"
    nsx_edge_2    = "nsxedge2"
    nsx_manager_1 = "nsxmanager1"
    nsx_manager_2 = "nsxmanager2"
    nsx_manager_3 = "nsxmanager3"
    sddc_manager  = "sddcmanager"
    v_center      = "vcenter"
  }

  service_access_security_groups = {
    security_groups = [awscc_ec2_security_group.evs_service.id]
  }

  initial_vlans = {
    vmk_management = {
      cidr = "10.0.10.0/24"
    }
    vm_management = {
      cidr = "10.0.11.0/24"
    }
    v_san = {
      cidr = "10.0.12.0/24"
    }
    v_motion = {
      cidr = "10.0.13.0/24"
    }
    v_tep = {
      cidr = "10.0.14.0/24"
    }
    edge_v_tep = {
      cidr = "10.0.15.0/24"
    }
    nsx_up_link = {
      cidr = "10.0.16.0/24"
    }
    hcx = {
      cidr = "10.0.17.0/24"
    }
    expansion_vlan_1 = {
      cidr = "10.0.18.0/24"
    }
    expansion_vlan_2 = {
      cidr = "10.0.19.0/24"
    }
  }

  # Required host configuration (must have exactly 4 hosts)
  hosts = [
    {
      instance_type = "i4i.metal"
      host_name     = "evs-host-1"
      key_name      = awscc_ec2_key_pair.evs_hosts.key_name
    },
    {
      instance_type = "i4i.metal"
      host_name     = "evs-host-2"
      key_name      = awscc_ec2_key_pair.evs_hosts.key_name
    },
    {
      instance_type = "i4i.metal"
      host_name     = "evs-host-3"
      key_name      = awscc_ec2_key_pair.evs_hosts.key_name
    },
    {
      instance_type = "i4i.metal"
      host_name     = "evs-host-4"
      key_name      = awscc_ec2_key_pair.evs_hosts.key_name
    }
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}