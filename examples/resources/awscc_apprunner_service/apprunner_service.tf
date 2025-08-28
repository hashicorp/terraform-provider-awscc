# App Runner Service
resource "awscc_apprunner_service" "example" {
  service_name = "example-apprunner-service"

  source_configuration = {
    auto_deployments_enabled = false

    image_repository = {
      image_configuration = {
        port = "8080"
        runtime_environment_variables = [
          {
            name  = "APP_ENV"
            value = "production"
          }
        ]
      }
      image_identifier      = "public.ecr.aws/aws-containers/hello-app-runner:latest"
      image_repository_type = "ECR_PUBLIC"
    }
  }

  instance_configuration = {
    cpu    = "1024"
    memory = "2048"
  }

  health_check_configuration = {
    protocol            = "HTTP"
    path                = "/"
    interval            = 10
    timeout             = 5
    healthy_threshold   = 1
    unhealthy_threshold = 5
  }

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-apprunner-service"
    }
  ]
}
