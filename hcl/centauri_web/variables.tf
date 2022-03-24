variable "name" {}
variable "region" {}
variable "vpc_id" {}
variable "subnet_ids" {
  type = list(string)
}

# RDS
variable "db_name" {}
variable "db_username" {}
variable "db_password" {}

# ECS
variable "docker_image" {}
variable "app_secret_key_base" {}
