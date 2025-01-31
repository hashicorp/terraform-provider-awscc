# Get current account ID
data "aws_caller_identity" "current" {}

# Get current region
data "aws_region" "current" {}

resource "awscc_gamelift_container_group_definition" "example" {
  name                         = "example-container-group"
  operating_system             = "AMAZON_LINUX_2023"
  total_memory_limit_mebibytes = 1024
  total_vcpu_limit             = 2

  game_server_container_definition = {
    container_name     = "game-server"
    image_uri          = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/game-server:latest"
    server_sdk_version = "5.2.0"

    port_configuration = {
      container_port_ranges = [
        {
          from_port = 7000
          to_port   = 7000
          protocol  = "UDP"
        }
      ]
    }

    environment_override = [
      {
        name  = "GAME_PORT"
        value = "7000"
      }
    ]
  }

  support_container_definitions = [
    {
      container_name              = "monitoring"
      image_uri                   = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/monitoring:latest"
      essential                   = false
      memory_hard_limit_mebibytes = 512
      vcpu                        = 1

      health_check = {
        command      = ["/healthcheck.sh"]
        interval     = 60
        retries      = 5
        timeout      = 30
        start_period = 60
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}