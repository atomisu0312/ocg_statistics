output "neon_endpoint_host" {
  description = "Host of the Neon endpoint"
  value       = neon_endpoint.this.host
}

output "neon_role_name" {
  description = "Name of the Neon role"
  value       = neon_role.this.name
}

output "neon_role_password" {
  description = "Password of the Neon role"
  value       = neon_role.this.password
  sensitive   = true
}

output "neon_database_name" {
  description = "Name of the Neon database"
  value       = neon_database.this.name
}

output "neon_branch_name" {
  description = "Name of the Neon branch"
  value       = neon_branch.this.name
}
