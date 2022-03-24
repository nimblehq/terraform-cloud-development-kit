variable "instance_type" {
  default = "db.t3.small"
}

variable "engine" {
  default = "postgres"
}

variable "engine_version" {
  default = "14.2"
}

variable "allocated_storage" {
  default = 5
}

variable "max_allocated_storage" {
  default = 10
}

variable "port" {
  default = "5432"
}

variable "db_identifier" {}
variable "username" {}
variable "password" {}
variable "db_name" {}
