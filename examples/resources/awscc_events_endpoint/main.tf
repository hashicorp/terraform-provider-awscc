# Event buses for routing configuration
resource "awscc_events_event_bus" "primary" {
  name = "primary-event-bus"
}

resource "awscc_events_event_bus" "secondary" {
  name = "secondary-event-bus"
}

# Create Route53 health check
resource "aws_route53_health_check" "endpoint" {
  fqdn              = "example.com"
  port              = 443
  type              = "HTTPS"
  resource_path     = "/health"
  failure_threshold = "3"
  request_interval  = "30"

  tags = {
    Name = "eventbridge-endpoint-health-check"
  }
}

# Create IAM role for EventBridge Endpoint
resource "aws_iam_role" "events_endpoint" {
  name = "EventBridgeEndpointRole"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "events.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "events_endpoint" {
  name = "EventBridgeEndpointPolicy"
  role = aws_iam_role.events_endpoint.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "events:PutEvents"
        ]
        Resource = [
          awscc_events_event_bus.primary.arn,
          awscc_events_event_bus.secondary.arn
        ]
      }
    ]
  })
}

# Create EventBridge Endpoint
resource "awscc_events_endpoint" "test" {
  name        = "test-endpoint"
  description = "Test EventBridge Endpoint"
  role_arn    = aws_iam_role.events_endpoint.arn

  event_buses = [
    {
      event_bus_arn = awscc_events_event_bus.primary.arn
    },
    {
      event_bus_arn = awscc_events_event_bus.secondary.arn
    }
  ]

  routing_config = {
    failover_config = {
      primary = {
        health_check = aws_route53_health_check.endpoint.arn
      }
      secondary = {
        route = "us-east-1"
      }
    }
  }
}