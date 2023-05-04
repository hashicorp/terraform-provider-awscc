resource "awscc_ecr_repository" "lifecycle_policy_example" {
  repository_name      = "example-ecr-repositry-lifecycle"
  image_tag_mutability = "MUTABLE"

  lifecycle_policy = {
    lifecycle_policy_text = <<EOF
        {
            "rules": [
                {
                    "rulePriority": 1,
                    "description": "Expire images older than 14 days",
                    "selection": {
                        "tagStatus": "untagged",
                        "countType": "sinceImagePushed",
                        "countUnit": "days",
                        "countNumber": 14
                    },
                    "action": {
                        "type": "expire"
                    }
                }
            ]
        }
        EOF
  }
}
