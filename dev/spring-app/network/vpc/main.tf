provider "aws" {
  region = "ap-northeast-2"
}

module "vpc" {
  source = "../../../../module/network/vpc"

  vpc_cidr_block = var.vpc_cidr_block
  public_subnet_cidr_block = var.public_subnet_cidr_block
  private_subnet_cidr_block = var.private_subnet_cidr_block
}
