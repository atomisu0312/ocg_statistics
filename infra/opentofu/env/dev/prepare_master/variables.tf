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

variable "lambda_cardinsert_name" {
  description = "The name of the Lambda function."
  type        = string
}

variable "lambda_cardinsert_env_pg_db_user" {
  description = "The name of the Lambda function."
  type        = string
}

variable "lambda_cardinsert_env_pg_db_password" {
  description = "The name of the Lambda function."
  type        = string
}

variable "lambda_cardinsert_env_pg_db_host_port" {
  description = "The name of the Lambda function."
  type        = string
}

variable "lambda_cardinsert_env_pg_db_name" {
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


