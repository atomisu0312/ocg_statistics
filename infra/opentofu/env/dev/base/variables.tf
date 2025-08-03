variable "lambda_sample_name" {
  description = "The name of the Lambda function."
  type        = string
}

variable "region" {
  description = "The region of the Lambda function."
  type        = string
}

variable "project" {
  description = "The project name."
  type        = string
}

variable "tag_base" {
  description = "The base tag for the project."
  type        = string
}

variable "ecr_lambda_sample_repo_name" {
  description = "The name of the ECR repository for the Lambda function."
  type        = string
}