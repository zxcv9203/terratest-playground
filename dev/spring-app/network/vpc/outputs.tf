output "vpc_id" {
  value = module.vpc.vpc_id
  description = "The Id of the VPC"
}

output "public_subnet_id" {
  value = module.vpc.public_subnet_id
  description = "The Id of public subnet ip"
}

output "private_subnet_id" {
  value = module.vpc.private_subnet_id
  description = "The Id of private subnet ip"
}