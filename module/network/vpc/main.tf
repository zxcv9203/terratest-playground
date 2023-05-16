# 사용하고자 하는 VPC
resource "aws_vpc" "my_vpc" {
  cidr_block = var.vpc_cidr_block
}

# 서브넷 생성 (public)
resource "aws_subnet" "public" {
  vpc_id = aws_vpc.my_vpc.id
  cidr_block = var.public_subnet_cidr_block
}

# 서브넷 생성 (private)
resource "aws_subnet" "private" {
  vpc_id = aws_vpc.my_vpc.id
  cidr_block = var.private_subnet_cidr_block

}

# 서브넷에서 인터넷 연결을 위한 Internet gateway 생성
resource "aws_internet_gateway" "my_vpc" {
  vpc_id = aws_vpc.my_vpc.id
}

# NAT 게이트웨이에서 사용할 EIP
resource "aws_eip" "nat" {
  vpc = true
}

# private 서브넷을 Internet과 연결하기 위해 사용하는 NAT gateway
resource "aws_nat_gateway" "my_vpc" {
  allocation_id = aws_eip.nat.id
  subnet_id = aws_subnet.public.id
}

# Internet gateway로 연결되는 public route table
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.my_vpc.id

  route {
    cidr_block = local.all_ips
    gateway_id = aws_internet_gateway.my_vpc.id
  }
}

# NAT gateway로 연결되는 private route table
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.my_vpc.id

  route {
    cidr_block = local.all_ips
    nat_gateway_id = aws_nat_gateway.my_vpc.id
  }
}

# 생성한 route table을 public 서브넷에 연결
resource "aws_route_table_association" "public" {
  subnet_id = aws_subnet.public.id
  route_table_id = aws_route_table.public.id
}

# 생성한 route table을 private 서브넷에 연결
resource "aws_route_table_association" "private" {
  subnet_id = aws_subnet.private.id
  route_table_id = aws_route_table.private.id
}