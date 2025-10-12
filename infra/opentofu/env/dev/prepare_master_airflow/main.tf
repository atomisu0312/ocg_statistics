
data "aws_caller_identity" "current" {}

data "terraform_remote_state" "base" {
  backend = "s3"
  config = {
    bucket = "atomisu-ocg-123456-terraform-state"
    key    = "dev/base/terraform.tfstate"
    region = var.region
  }
}

locals {
  account_id = data.aws_caller_identity.current.account_id
}

module "s3_for_mwaa" {
  source = "../../../modules/s3/mwaa"
  mwaa_bucket_name = var.mwaa_bucket_name
  tag_base = var.tag_base
}

module "airflow" {
  source = "../../../modules/airflow"

  tag_base = var.tag_base
  source_bucket_arn = module.s3_for_mwaa.mwaa_bucket_arn
  airflow_version = "2.10.3"
  account_id = local.account_id
  region = var.region
  internet_gateway_id = var.internet_gateway_id 
  private_subnet_cidrs = var.private_subnet_cidrs
  public_subnet_cidrs  = var.public_subnet_cidrs
  vpc_id               = var.vpc_id
  webserver_access_mode = var.webserver_access_mode
}
