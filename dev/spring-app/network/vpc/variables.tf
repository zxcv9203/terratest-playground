variable "vpc_cidr_block" {
  description = "default subnet cidr"
  default = "10.0.0.0/16"
  type = string
}

variable "public_subnet_cidr_block" {
  description = "public subnet cidr"
  default = "10.0.0.0/24"
  type = string
}

variable "private_subnet_cidr_block" {
  description = "private subnet cidr"
  default = "10.0.1.0/24"
  type = string
}