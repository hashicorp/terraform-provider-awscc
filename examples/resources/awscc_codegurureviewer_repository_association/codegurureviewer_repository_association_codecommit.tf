resource "awscc_codegurureviewer_repository_association" "example" {
  name = var.repo_name
  type = "CodeCommit"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

variable "repo_name" {
  type = string
}
