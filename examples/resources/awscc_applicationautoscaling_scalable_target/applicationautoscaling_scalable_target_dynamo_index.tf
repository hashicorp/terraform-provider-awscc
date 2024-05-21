resource "awscc_applicationautoscaling_scalable_target" "example" {
  max_capacity       = 100
  min_capacity       = 5
  resource_id        = "table/${var.dynamodb_table_name}/index/${var.index_name}"
  scalable_dimension = "dynamodb:index:ReadCapacityUnits"
  service_namespace  = "dynamodb"
}

variable "dynamodb_table_name" {
  type = string
}

variable "index_name" {
  type = string
}