# Image Builder Component for installing httpd
resource "awscc_imagebuilder_component" "httpd" {
  name     = "install-httpd-component"
  platform = "Linux"
  version  = "1.0.0"

  description = "Component to install and configure Apache HTTP Server"
  data = yamlencode({
    schemaVersion = "1.0"
    phases = [{
      name = "build"
      steps = [{
        name   = "InstallHttpd"
        action = "ExecuteBash"
        inputs = {
          commands = [
            "yum install -y httpd",
            "systemctl enable httpd",
            "systemctl start httpd"
          ]
        }
      }]
    }]
  })

  supported_os_versions = ["Amazon Linux 2"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}