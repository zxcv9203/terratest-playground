variable "vpc_cidr_block" {
  description = "default subnet cidr"
  type = string
}

variable "public_subnet_cidr_block" {
  description = "public subnet cidr"
  type = string
}

variable "private_subnet_cidr_block" {
  description = "private subnet cidr"
  type = string
}

locals {
  all_ips = "0.0.0.0/0"
}