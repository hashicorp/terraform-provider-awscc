resource "aws_greengrassv2_component_version" "MyGreengrassComponentVersion-NestedLambda" {
  lambda_function {
    component_name    = "MyLambdaComponent"
    component_version = "1.0.0"
    lambda_arn        = "arn:aws:lambda:<region>:<account>:function:<LambdaFunctionName>:<LambdaVersion>"

    component_dependencies {
      dependency_type     = "runtime"
      version_requirement = ">=1.0.0"
    }

    component_lambda_parameters {
      environment_variables = {
        ENV_VAR1 = "value1"
        ENV_VAR2 = "value2"
      }

      event_sources {
        topic = "topic1"
        type  = "sns"
      }

      exec_args                   = ["arg1", "arg2"]
      input_payload_encoding_type = "json"

      linux_process_params {
        isolation_mode {
          devices {
            add_group_owner = true
            path            = "/dev/random"
            permission      = "rw"
          }
          memory_size_in_kb = 1024
          mount_ro_sysfs    = false

          volumes {
            add_group_owner  = false
            destination_path = "/mnt/data"
            permission       = "rw"
            source_path      = "/data"
          }
        }
        container_params {
          environment_variables = {
            CONTAINER_ENV_VAR1 = "container_value1"
            CONTAINER_ENV_VAR2 = "container_value2"
          }
          image_uri         = "123456789012.dkr.ecr.us-west-2.amazonaws.com/my-container-image:latest"
          memory_size_in_kb = 2048
          vcpu_count        = 2
        }
      }

      max_idle_time_in_seconds  = 300
      max_instances_count       = 3
      max_queue_size            = 10
      pinned                    = false
      status_timeout_in_seconds = 60
      timeout_in_seconds        = 30
    }

    component_platforms {
      attributes = {
        platform_version = "1.2.3"
        architecture     = "arm64"
      }
      name = "Linux"
    }
  }

  tags = {
    Environment = "Production"
    Project     = "GreengrassProject-Nested"
  }
}
