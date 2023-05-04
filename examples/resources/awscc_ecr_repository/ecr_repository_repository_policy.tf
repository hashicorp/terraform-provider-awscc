resource "awscc_ecr_repository" "repo_policy_example" {
  repository_name      = "example-ecr-repository-repo"
  image_tag_mutability = "MUTABLE"

  repository_policy_text = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Sid" : "DenyPull",
          "Effect" : "Deny",
          "Principal" : "*",
          "Action" : [
            "ecr:BatchGetImage",
            "ecr:GetDownloadUrlForLayer"
          ]
        }
      ]
    }
  )

}
