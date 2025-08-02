module "neon_db" {
  source = "../../../modules/neon-db"

  neon_api_key    = var.neon_api_key
  neon_project_name = var.neon_project_name
}

module "secretmanager" {
  source = "../../../modules/secretmanager"

  secret_name        = "neon/${module.neon_db.neon_branch_name}/${module.neon_db.neon_database_name}/${module.neon_db.neon_role_name}"
  secret_description = "Neon SaaS access details for ${module.neon_db.neon_database_name}, ${module.neon_db.neon_role_name} @ ${module.neon_db.neon_branch_name}"
  tags = {
    project  = "demo"
    platform = "neon"
  }

  neon_endpoint_host = module.neon_db.neon_endpoint_host
  neon_role_name     = module.neon_db.neon_role_name
  neon_role_password = module.neon_db.neon_role_password
  neon_database_name = module.neon_db.neon_database_name

  iam_policy_name = "${module.neon_db.neon_branch_name}-${module.neon_db.neon_database_name}-${module.neon_db.neon_role_name}"
  iam_policy_path = "/neon/read-only/"
}
