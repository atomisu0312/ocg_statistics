# Terraform configuration for the ocg-sample Lambda function.

data "aws_caller_identity" "current" {}

# Data source to read the outputs from the 'base' state file.
data "terraform_remote_state" "base" {
  backend = "s3"
  config = {
    bucket = "${var.project}-terraform-state"
    key    = "dev/base/terraform.tfstate"
    region = var.region
  }
}

# Local variables for constructing the ECR image URI.
locals {
  image_tag = "latest"
  account_id = data.aws_caller_identity.current.account_id

  # Construct the full ECR image URI dynamically using the created repository's URL.
  image_uri_sample = "${data.terraform_remote_state.base.outputs.ecr_lambda_sample_repository_url}:${local.image_tag}"
  image_uri_idcheck = "${data.terraform_remote_state.base.outputs.ecr_lambda_idcheck_repository_url}:${local.image_tag}"
}

module "lambda_sample" {
  source = "../../../modules/lambda/sample"

  lambda_name = var.lambda_sample_name
  tag_base = var.tag_base
  role_arn = data.terraform_remote_state.base.outputs.lambda_exec_role_arn
  image_uri = local.image_uri_sample
}

module "lambda_idcheck" {
  source = "../../../modules/lambda/idcheck"

  lambda_name = var.lambda_idcheck_name
  tag_base = var.tag_base
  role_arn = data.terraform_remote_state.base.outputs.lambda_exec_role_arn
  image_uri = local.image_uri_idcheck
}

module "parameter_current_id" {
  source = "../../../modules/parameterstore/currentid"

  parameter_name = var.parameter_current_id_name
  parameter_value = var.parameter_current_id_value
  tag_base = var.tag_base
}

module "parameter_delta_id" {
  source = "../../../modules/parameterstore/deltaid"

  parameter_name = var.parameter_delta_id_name
  parameter_value = var.parameter_delta_id_value
  tag_base = var.tag_base
}

