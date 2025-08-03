variable "lambda_exec_role_name" {
  description = "The name of the Lambda function execution role."
  type        = string
}

variable "lambda_exec_role_policy_name" {
  description = "The name of the Lambda function execution role policy."
  type        = string
}

variable "tag_base" {
  description = "The base tag for the project."
  type        = string
}

variable "log_group_arn_wildcard" {
  description = "The ARN of the log group."
  type        = string
}

variable "log_stream_arn" {
  description = "The ARN of the log stream."
  type        = string
}