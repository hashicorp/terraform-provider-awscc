resource "awscc_xray_sampling_rule" "example" {
  sampling_rule = {
    rule_name      = "example"
    fixed_rate     = 0.05
    host           = "MyHost"
    http_method    = "GET"
    priority       = 9999
    reservoir_size = 1
    resource_arn   = "*"
    service_name   = "MyServiceName"
    service_type   = "MyServiceType"
    url_path       = "*"
    version        = 1
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}

