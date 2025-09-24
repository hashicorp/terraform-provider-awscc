resource "awscc_ec2_launch_template" "with eip" {
  launch_template_data = {
    image_id = data.aws_ami.amazon_linux.id
    network_interfaces = [
      {
        device_index                = 0
        associate_public_ip_address = true
        delete_on_termination       = true
        groups = [
          "sg-example1",
          "sg-example2"
        ]
      }
    ]
    monitoring = {
      enabled = true
    }
    instance_type = "t2.micro"
  }
  launch_template_name = "with_eip"
}