// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scheduler_test

import (
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSchedulerSchedule_ecsParameters_tags(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Scheduler::Schedule", "awscc_scheduler_schedule", "test")
	rName := td.RandomName()

	td.ResourceTestWithTestCase(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"awscc": {
						Source:            "hashicorp/awscc",
						VersionConstraint: "1.89.0",
					},
					"aws": {
						Source:            "hashicorp/aws",
						VersionConstraint: "6.51.0",
					},
				},
				Config: acctest.ConfigCompose(
					testAccAWSScheduleConfig_base(&td, rName),
					testAccAWSScheduleConfig_old(&td, rName),
				),
				// error is returned from CloudControl API.
				// object is serialized as incorrect type.
				ExpectError: regexache.MustCompile(`(?s)Model validation failed.*expected type: JSONObject, found: JSONArray`),
			},
			{
				ExternalProviders: map[string]resource.ExternalProvider{
				"aws": {
					Source:            "hashicorp/aws",
					VersionConstraint: "6.51.0",
				},
			},
				Config: acctest.ConfigCompose(
					testAccAWSScheduleConfig_base(&td, rName),
					testAccAWSScheduleConfig_new(&td, rName),
				),
				ProtoV6ProviderFactories: td.ProviderFactories(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(td.ResourceName, tfjsonpath.New("target").AtMapKey("ecs_parameters").AtMapKey("tags").AtSliceIndex(0).AtMapKey("key"), knownvalue.StringExact("environment")),
					statecheck.ExpectKnownValue(td.ResourceName, tfjsonpath.New("target").AtMapKey("ecs_parameters").AtMapKey("tags").AtSliceIndex(0).AtMapKey("value"), knownvalue.StringExact("Prod")),
				},
			},
		},
	})
}

func testAccAWSScheduleConfig_base(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource "aws_iam_role" "scheduler_role" {
  name = "%[3]s-scheduler"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "scheduler.amazonaws.com"
        },
        Effect = "Allow",
        Sid    = ""
      }
    ]
  })
}

# IAM role for the ECS task
resource "aws_iam_role" "ecs_task_execution_role" {
  name = "%[3]s-ecs"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Effect = "Allow",
        Sid    = ""
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_cluster" "test" {
  name = %[3]q
}

# VPC and networking
resource "aws_vpc" "test" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}

resource "aws_subnet" "test1" {
  vpc_id            = aws_vpc.test.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"
}

resource "aws_subnet" "test2" {
  vpc_id            = aws_vpc.test.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"
}

resource "aws_security_group" "ecs_task" {
  name        = %[3]q
  description = "Security group for ECS Fargate tasks"
  vpc_id      = aws_vpc.test.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_ecs_task_definition" "test" {
  family                   = %[3]q
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([
    {
      name      = %[3]q
      image     = "public.ecr.aws/amazonlinux/amazonlinux:latest"
      essential = true
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = "/ecs/test-task"
          "awslogs-region"        = "us-west-2"
          "awslogs-stream-prefix" = "ecs"
        }
      }
    }
  ])
}

resource "aws_iam_role_policy" "ecs_execution" {
  name = %[3]q
  role = aws_iam_role.scheduler_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action   = ["ecs:RunTask"],
        Resource = aws_ecs_task_definition.test.arn,
        Effect   = "Allow"
      },
      {
        Action   = ["iam:PassRole"],
        Resource = aws_iam_role.ecs_task_execution_role.arn,
        Effect   = "Allow"
      }
    ]
  })
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}

func testAccAWSScheduleConfig_old(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
# EventBridge Scheduler Schedule using the AWSCC provider
resource %[1]q %[2]q {
  name                         = %[3]q
  description                  = "Test schedule created with Terraform"
  schedule_expression          = "cron(0 12 * * ? *)" # Daily at 12:00 PM
  schedule_expression_timezone = "America/Los_Angeles"

  flexible_time_window = {
    mode                      = "FLEXIBLE"
    maximum_window_in_minutes = 30 # Allow 30-minute execution window
  }

  target = {
    arn      = aws_ecs_cluster.test.arn
    role_arn = aws_iam_role.scheduler_role.arn

    ecs_parameters = {
      task_definition_arn    = aws_ecs_task_definition.test.arn
      launch_type            = "FARGATE"
      task_count             = 1
      network_configuration = {
        awsvpc_configuration = {
          assign_public_ip = "DISABLED"
          subnets          = [aws_subnet.test1.id, aws_subnet.test2.id]
          security_groups  = [aws_security_group.ecs_task.id]
        }
      }

      tags = [
        ["this"]
      ]
    }

    retry_policy = {
      maximum_retry_attempts       = 3
      maximum_event_age_in_seconds = 3600
    }
  }

  state = "ENABLED" # Schedule is active
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}

func testAccAWSScheduleConfig_new(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
# EventBridge Scheduler Schedule using the AWSCC provider
resource %[1]q %[2]q {
  name                         = %[3]q
  description                  = "Test schedule created with Terraform"
  schedule_expression          = "cron(0 12 * * ? *)" # Daily at 12:00 PM
  schedule_expression_timezone = "America/Los_Angeles"

  flexible_time_window = {
    mode                      = "FLEXIBLE"
    maximum_window_in_minutes = 30 # Allow 30-minute execution window
  }

  target = {
    arn      = aws_ecs_cluster.test.arn
    role_arn = aws_iam_role.scheduler_role.arn

    ecs_parameters = {
      task_definition_arn    = aws_ecs_task_definition.test.arn
      launch_type            = "FARGATE"
      task_count             = 1
      network_configuration = {
        awsvpc_configuration = {
          assign_public_ip = "DISABLED"
          subnets          = [aws_subnet.test1.id, aws_subnet.test2.id]
          security_groups  = [aws_security_group.ecs_task.id]
        }
      }

      tags = [
        {
          key   = "environment"
          value = "Prod"
        }
      ]
    }

    retry_policy = {
      maximum_retry_attempts       = 3
      maximum_event_age_in_seconds = 3600
    }
  }

  state = "ENABLED" # Schedule is active
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
