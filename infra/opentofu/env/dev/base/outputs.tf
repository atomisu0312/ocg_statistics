output "lambda_exec_role_arn" {
  description = "The ARN of the IAM role for the ocg-sample Lambda function."
  value       = module.lambda_exec_role.lambda_exec_role_arn
}

output "ecr_lambda_sample_repository_url" {
  description = "The URL of the ECR repository."
  value       = module.ecr_sample.ecr_repository_url
}
