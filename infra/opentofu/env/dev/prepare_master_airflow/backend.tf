terraform {
  backend "s3" {
    bucket = "atomisu-ocg-123456-terraform-state"
    key    = "dev/prepare_master_airflow/terraform.tfstate"
    region = "ap-northeast-1"
  }
}
