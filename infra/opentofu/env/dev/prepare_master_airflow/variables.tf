
variable "project" {
  description = "The name of the project."
  type        = string
}

variable "region" {
  description = "The AWS region."
  type        = string
}

variable "tag_base" {
  description = "The base tag for resources."
  type        = string
}

variable "mwaa_bucket_name" {
  description = "The name of the S3 bucket for MWAA."
  type        = string
}

variable "airflow_version" {
  description = "The version of Apache Airflow to use."
  type        = string
}

variable "internet_gateway_id" {
  description = "The ID of the Internet Gateway."
  type        = string
}

variable "private_subnet_cidrs" {
  description = "The private subnet CIDR blocks."
  type        = list(string)
}

variable "public_subnet_cidrs" {
  description = "The public subnet CIDR blocks."
  type        = list(string)
}

variable "vpc_id" {
  description = "The ID of the VPC."
  type        = string
}

variable "webserver_access_mode" {
  description = "The webserver access mode for MWAA."
  type        = string
}
