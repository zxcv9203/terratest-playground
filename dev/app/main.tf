provider "aws" {
  region = "us-west-1"
}

module "hello_world_app" {
  source = "../../module/app"

  server_text = var.server_text

  environment = "example"
  db_remote_state_bucket = var.db_remote_state_bucket
  db_remote_state_key = var.db_remote_state_key
  mysql_config = var.mysql_config

  instance_type = "t2.micro"
  min_size = 2
  max_size = 2
  enable_autoscaling = false
}