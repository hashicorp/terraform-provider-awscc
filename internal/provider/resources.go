//go:generate go run generators/resource/main.go -resource aws_logs_log_group -cfschema ../service/cloudformation/schema-generator/testdata/aws_logs_log_group.cf-resource-schema.json -- aws_logs_log_group_gen.go

package provider
