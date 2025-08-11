resource "aws_ssm_parameter" "currentid" {
  name  = var.parameter_name
  type  = "String"
  value = var.parameter_value
  tags = {
    Name = var.tag_base
  }
}
