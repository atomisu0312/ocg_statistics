# Terraform configuration for the base IAM role and policy.

# Data sources to dynamically get the current AWS Account ID and Region.
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Local variables for constructing ARNs.
locals {
  account_id = data.aws_caller_identity.current.account_id
  region     = data.aws_region.current.name

  # Construct the log group ARNs dynamically.
  log_group_arn_wildcard = "arn:aws:logs:${local.region}:${local.account_id}:*"
  log_stream_arn         = "arn:aws:logs:${local.region}:${local.account_id}:log-group:/aws/lambda/${var.lambda_sample_name}:*"
}


module "lambda_exec_role" {
  source = "../../../modules/role/lambda"

  lambda_exec_role_name = "${var.project}-lambda-exec-role"
  lambda_exec_role_policy_name = "${var.project}-lambda-policy"
  tag_base = var.tag_base
  log_group_arn_wildcard = local.log_group_arn_wildcard
  log_stream_arn = local.log_stream_arn
}

module "ecr_sample" {
  source = "../../../modules/ecr/sample"

  repo_name = var.ecr_lambda_sample_repo_name
  tag_base                    = var.tag_base
}