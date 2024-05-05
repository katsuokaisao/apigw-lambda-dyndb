resource "aws_ecr_repository" "go_lambda" {
  name                 = "go-lambda"
  image_tag_mutability = "MUTABLE"

  encryption_configuration {
    encryption_type = "AES256"
  }

  image_scanning_configuration {
    scan_on_push = "false"
  }
}

locals {
  ecr-lifecycle-policy = {
    rules = [
      {
        action = {
          type = "expire"
        }
        description  = "最新のイメージを5つだけ残す"
        rulePriority = 1
        selection = {
          countNumber   = 5
          countType     = "imageCountMoreThan"
          tagStatus     = "tagged"
          tagPrefixList = ["latest"]
        }
      },
    ]
  }
}

resource "aws_ecr_lifecycle_policy" "go_lambda" {
  repository = aws_ecr_repository.go_lambda.name
  policy     = jsonencode(local.ecr-lifecycle-policy)
}