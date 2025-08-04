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