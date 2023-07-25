resource "awscc_ec2_volume" "example" {
  availability_zone = "us-west-2a"
  size              = 10
}

resource "aws_instance" "web" {
  ami               = "ami-08541bb85074a743a"
  instance_type     = "t3.micro"
  availability_zone = "us-west-2a"

  tags = {
    Name = "HelloWorld"
  }
}

resource "awscc_ec2_volume_attachment" "this" {
  instance_id = aws_instance.web.id
  volume_id   = awscc_ec2_volume.example.id
  device      = "/dev/sdh"
  depends_on  = [aws_instance.web, awscc_ec2_volume.example]
}