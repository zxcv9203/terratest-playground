output "vpc_id" {
  value = aws_vpc.my_vpc.id
  description = "The Id of the VPC"
}

output "public_subnet_id" {
  value = aws_subnet.public.id
  description = "The Id of public subnet ip"
}

output "private_subnet_id" {
  value = aws_subnet.private.id
  description = "The Id of private subnet ip"
}