resource "awscc_applicationautoscaling_scalable_target" "example" {
  max_capacity       = 100
  min_capacity       = 5
  resource_id        = "table/${var.dynamodb_table_name}"
  scalable_dimension = "dynamodb:table:WriteCapacityUnits"
  service_namespace  = "dynamodb"
}

variable "dynamodb_table_name" {
  type = string
}