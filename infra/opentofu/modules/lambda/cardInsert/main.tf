resource "aws_lambda_function" "lambda_cardinsert" {
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
  
  environment {
    variables = {
      PG_DB_USER = var.environment.PG_DB_USER
      PG_DB_PASSWORD = var.environment.PG_DB_PASSWORD
      PG_DB_HOST_PORT = var.environment.PG_DB_HOST_PORT
      PG_DB_NAME = var.environment.PG_DB_NAME
    }
  }

  tags = {
    project = var.tag_base
  }
}