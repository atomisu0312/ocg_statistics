variable "tag_base" {
  description = "The base tag for resources."
  type        = string
}

variable "source_bucket_arn" {
  description = "The ARN of the S3 bucket for MWAA."
  type        = string
}

variable "airflow_version" {
  description = "The version of Apache Airflow to use."
  type        = string
}

variable "account_id" {
  description = "The AWS account ID."
  type        = string
}

variable "region" {
  description = "The AWS region."
  type        = string
}

variable "internet_gateway_id" {
  description = "The ID of the Internet Gateway."
  type        = string
}

variable "private_subnet_cidrs" {
  description = "List of private subnet CIDR blocks."
  type        = list(string)
}

variable "public_subnet_cidrs" {
  description = "List of public subnet CIDR blocks."
  type        = list(string)
}

variable "vpc_id" {
  description = "The ID of the VPC."
  type        = string
}

variable "webserver_access_mode" {
  description = "The webserver access mode for MWAA."
  type        = string
  default     = "PUBLIC_ONLY"
}
