variable "secret_name" {
  description = "Name of the Secrets Manager secret"
  type        = string
}

variable "secret_description" {
  description = "Description of the Secrets Manager secret"
  type        = string
}

variable "tags" {
  description = "Tags for the Secrets Manager secret"
  type        = map(string)
  default     = {}
}

variable "neon_endpoint_host" {
  description = "Host of the Neon endpoint"
  type        = string
}

variable "neon_role_name" {
  description = "Name of the Neon role"
  type        = string
}

variable "neon_role_password" {
  description = "Password of the Neon role"
  type        = string
  sensitive   = true
}

variable "neon_database_name" {
  description = "Name of the Neon database"
  type        = string
}

variable "iam_policy_name" {
  description = "Name of the IAM policy"
  type        = string
}

variable "iam_policy_path" {
  description = "Path for the IAM policy"
  type        = string
  default     = "/"
}
