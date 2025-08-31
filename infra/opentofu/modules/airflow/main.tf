module "airflow" {
  source = "idealo/mwaa/aws"
  version = "3.2.1"

  account_id = var.account_id
  environment_name = var.tag_base
  internet_gateway_id = var.internet_gateway_id
  private_subnet_cidrs = var.private_subnet_cidrs
  public_subnet_cidrs = var.public_subnet_cidrs
  region = var.region
  source_bucket_arn = var.source_bucket_arn
  vpc_id = var.vpc_id
  airflow_version = var.airflow_version
  webserver_access_mode = var.webserver_access_mode
}