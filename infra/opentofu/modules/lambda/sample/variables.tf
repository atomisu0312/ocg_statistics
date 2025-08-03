variable "lambda_sample_name" {
  description = "The name of the Lambda function."
  type        = string
}

variable "tag_base" {
  description = "The base tag for the project."
  type        = string
}

variable "role_arn" {
  description = "The ARN of the IAM role for the Lambda function."
  type        = string
}

variable "image_uri" {
  description = "The URI of the ECR image for the Lambda function."
  type        = string
}