resource "awscc_ecr_pull_through_cache_rule" "terraform_ecr_pull_through_cache_rule" {
  ecr_repository_prefix = "ecr-public"
  upstream_registry_url = "public.ecr.aws"
}