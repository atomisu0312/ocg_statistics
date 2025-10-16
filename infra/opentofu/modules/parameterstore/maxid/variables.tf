variable "parameter_name" {
  description = "The name of the SSM parameter."
  type        = string
}

variable "parameter_value" {
  description = "The value of the SSM parameter."
  type        = string
}

variable "tag_base" {
  description = "The base tag for the project."
  type        = string
}
