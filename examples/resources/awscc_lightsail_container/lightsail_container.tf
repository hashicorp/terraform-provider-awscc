resource "awscc_lightsail_container" "example" {
  power        = "nano"
  scale        = "1"
  service_name = "example-service"
  container_service_deployment = {
    containers = [{
      container_name = "example-container"
      image          = "public.ecr.aws/nginx/nginx:latest"
      ports = [{
        port     = 80
        protocol = "HTTP"
      }]
    }]
    public_endpoint = {
      container_name = "example-container"
      container_port = 80
    }
  }
}
