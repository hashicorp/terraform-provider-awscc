resource "awscc_accessanalyzer_analyzer" "this" {
  analyzer_name = "example"
  type          = "ACCOUNT"
}