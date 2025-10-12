resource "aws_lambda_function" "lambda_idcheck" {
  function_name = var.lambda_name
  package_type  = "Image"
  role          = var.role_arn
  image_uri     = var.image_uri

  memory_size   = 128
  timeout       = 3
  architectures = ["arm64"]

  ephemeral_storage {
    size = 512
  }
  
  environment {
    variables = {
      ENVIRONMENT = "production"
      LOG_LEVEL   = "info"
    }
  }

  tags = {
    project = var.tag_base
  }
}