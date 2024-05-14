resource "awscc_opensearchserverless_vpc_endpoint" "example" {
  name               = "example-endpoint"
  security_group_ids = var.security_group_ids
  vpc_id             = var.vpc_id
  subnet_ids         = var.subnet_ids
}

variable "vpc_id" {
  type = string
}

variable "security_group_ids" {
  type = list(string)
}

variable "subnet_ids" {
  type = list(string)
}
