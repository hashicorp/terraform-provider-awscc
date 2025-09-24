resource "awscc_ecr_pull_through_cache_rule" "ecr_pull_through_cache_rule_example" {
  ecr_repository_prefix = "ecr-public"
  upstream_registry_url = "public.ecr.aws"
}