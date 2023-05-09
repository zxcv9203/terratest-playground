provider "aws" {
  region = "us-west-1"
}

module "hello_world_app" {
  source = "../../module/app"

  server_text = "Hello, World"
  environment = "example"

  mysql_config = var.mysql_config

  instance_type = "t2.micro"
  min_size = 2
  max_size = 2
  enable_autoscaling = false
}