resource "aws_lambda_function" "lambda_sample" {
  function_name = var.lambda_name
  package_type  = "Image"
  role          = var.role_arn
  image_uri     = var.image_uri

  memory_size   = 1024
  timeout       = 240
  architectures = ["arm64"]

  ephemeral_storage {
    size = 512
  }

  tags = {
    project = var.tag_base
  }
}