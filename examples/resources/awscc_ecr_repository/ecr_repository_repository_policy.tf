resource "awscc_ecr_repository" "repo_policy_example" {
  repository_name      = "example-ecr-repository-policy"
  image_tag_mutability = "MUTABLE"

  repository_policy_text = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Sid" : "CodeBuildAccess",
          "Effect" : "Allow",
          "Principal" : {
            "Service" : "codebuild.amazonaws.com"
          },
          "Action" : [
            "ecr:BatchGetImage",
            "ecr:GetDownloadUrlForLayer"
          ],
          "Condition" : {
            "ArnLike" : {
              "aws:SourceArn" : "arn:aws:codebuild:region:123456789012:project/project-name"
            },
            "StringEquals" : {
              "aws:SourceAccount" : "123456789012"
            }
          }
        }
      ]
    }
  )

}
