resource "aws_ecr_repository" "lambda_sample" {
  name                 = var.ecr_lambda_sample_repo_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    project = var.tag_base
  }
}
