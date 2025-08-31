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

variable "ecr_repo_name" {
  description = "The name of the ECR repository."
  type        = string
}

variable "lambda_sample_name" {
  description = "The name of the Lambda function."
  type        = string
}

variable "lambda_idcheck_name" {
  description = "The name of the Lambda function."
  type        = string
}
variable "parameter_current_id_name" {
  description = "The name of the SSM parameter."
  type        = string 
}

variable "parameter_current_id_value" {
  description = "The value of the SSM parameter."
  type        = string
}

variable "parameter_delta_id_name" {
  description = "The name of the SSM parameter."
  type        = string 
}

variable "parameter_delta_id_value" {
  description = "The value of the SSM parameter."
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
