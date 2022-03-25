terraform {
  cloud {
    organization = "nimble"

    workspaces {
      name = "nimble-growth-37-centauri-ecr"
    }
  }
}

provider "aws" {
  region = var.region
}

resource "aws_ecr_repository" "app" {
  name = var.name
}

resource "aws_ecr_lifecycle_policy" "foopolicy" {
  repository = aws_ecr_repository.app.name

  policy = <<EOF
{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Delete build",
            "selection": {
                "countType": "imageCountMoreThan",
                "countNumber": 5,
                "tagStatus": "tagged",
                "tagPrefixList": [
                    "development-sit-",
                    "development-uat-"
                ]
            },
            "action": {
                "type": "expire"
            }
        },
        {
            "rulePriority": 2,
            "description": "Delete untagged images",
            "selection": {
                "countType": "sinceImagePushed",
                "countNumber": 1,
                "tagStatus": "untagged",
                "countUnit": "days"
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
EOF
}
