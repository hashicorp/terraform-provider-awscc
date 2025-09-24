resource "awscc_lightsail_disk" "example" {
  disk_name         = "example-disk"
  size_in_gb        = "32"
  availability_zone = "us-east-1a"
}

resource "awscc_lightsail_instance" "example" {
  blueprint_id      = "amazon_linux_2023"
  bundle_id         = "nano_3_0"
  instance_name     = "example-instance"
  availability_zone = "us-east-1a"
  hardware = {
    disks = [{
      disk_name = awscc_lightsail_disk.example.disk_name
      path      = "/dev/xvdf"
    }]
  }
  networking = {
    ports = [
      {
        access_direction = "inbound"
        access_type      = "public"
        cidrs            = ["0.0.0.0/0"]
        from_port        = 22
        ipv_6_cidrs      = ["::/0"]
        protocol         = "tcp"
        to_port          = 22
      },
      {
        access_direction = "inbound"
        access_type      = "public"
        cidrs            = ["0.0.0.0/0"]
        from_port        = 8080
        ipv_6_cidrs      = ["::/0"]
        protocol         = "tcp"
        to_port          = 8080
      }
    ]
  }
  user_data = <<-EOF
    #!/bin/sh
    end=$((SECONDS+60))
    while [ ! -e /dev/xvdf ]
    do
      if [ $SECONDS -gt $end ]; then
        echo "Timeout waiting for /dev/xvdf."
        exit 1
      fi
      echo "Waiting for /dev/xvdf..."
      sleep 5
    done
    file -s /dev/xvdf
    mkfs -t xfs /dev/xvdf
    mkdir /var/www
    mount /dev/xvdf /var/www
    echo "/dev/xvdf /var/www xfs defaults 0 0" >> /etc/fstab
    dnf install -y httpd
    sed -i 's/Listen 80/Listen 8080/' /etc/httpd/conf/httpd.conf
    echo "Hello, World!" > /var/www/html/index.html
    systemctl start httpd
    systemctl enable httpd
    usermod -a -G apache ec2-user
    chown -R ec2-user:apache /var/www
    chmod 2775 /var/www
    find /var/www -type d -exec sudo chmod 2775 {} \;
    find /var/www -type f -exec sudo chmod 0664 {} \;
  EOF
}
