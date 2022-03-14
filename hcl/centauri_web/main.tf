terraform {
  cloud {
    organization = "nimble"

    workspaces {
      name = "nimble-growth-37-centauri-hcl"
    }
  }
}

provider "aws" {
  region = var.region
}

module "alb" {
  source = "./modules/alb"

  name       = var.name
  vpc_id     = var.vpc_id
  subnet_ids = var.subnet_ids
}

module "rds" {
  source = "./modules/rds"

  db_identifier = var.name
  db_name       = var.db_name
  username      = var.db_username
  password      = var.db_password
}

module "log" {
  source = "./modules/log"

  name = var.name
}

module "ecs" {
  source = "./modules/ecs"

  image = var.docker_image

  region     = var.region
  name       = var.name
  subnet_ids = var.subnet_ids

  aws_alb_target_group_arn = module.alb.alb_target_group_arn
  aws_alb_dns              = module.alb.alb_dns

  cloudwatch_log_group_arn = module.log.cloudwatch_log_group_arn

  database_name     = var.db_name
  database_username = var.db_username
  database_password = var.db_password
  database_host     = module.rds.aws_db_instance_host
  database_port     = module.rds.aws_db_instance_port
}
