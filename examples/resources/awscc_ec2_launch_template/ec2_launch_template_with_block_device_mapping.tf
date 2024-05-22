resource "awscc_ec2_launch_template" "block_device" {
  launch_template_data = {
    image_id = data.aws_ami.amazon_linux.id
    block_device_mappings = [
      {
        ebs = {
          volume_size         = 30
          volume_type         = "gp3"
          delete_on_terminate = true
          encrypted           = true
        }
        device_name = "/dev/xvdcz"
      }
    ]
    monitoring = {
      enabled = true
    }
    instance_type = "t2.micro"
  }
  launch_template_name = "with_block_device"
}