resource "aws_ecr_repository" "ecr" {
  name                 = var.repo_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    project = var.tag_base
  }
}
