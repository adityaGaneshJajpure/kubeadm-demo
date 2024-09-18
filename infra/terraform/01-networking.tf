# Create a VPC
resource "aws_vpc" "k8s_vpc" {
  cidr_block           = var.vpc_cidr_block
  instance_tenancy     = "default"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "k8s_vpc"
  }
}

# Create private subnet
resource "aws_subnet" "k8s_private_subnet" {
  vpc_id                                      = aws_vpc.k8s_vpc.id
  cidr_block                                  = var.private_subnet_cidr_block
  map_public_ip_on_launch                     = false
  enable_resource_name_dns_a_record_on_launch = true
  private_dns_hostname_type_on_launch         = "resource-name"
  tags = {
    Name = "k8s_private_subnet"
  }
}

# Create public subnet 
resource "aws_subnet" "k8s_public_subnet" {
  vpc_id                                      = aws_vpc.k8s_vpc.id
  cidr_block                                  = var.public_subnet_cidr_block
  map_public_ip_on_launch                     = true
  enable_resource_name_dns_a_record_on_launch = true
  private_dns_hostname_type_on_launch         = "resource-name"
  tags = {
    Name = "k8s_public_subnet"
  }
}

# Create an internet gateway
resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.k8s_vpc.id
}

# Create a public route table for reaching the Internet
resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.k8s_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }

  route {
    ipv6_cidr_block = "::/0"
    gateway_id      = aws_internet_gateway.this.id
  }

  tags = {
    Name = "public_route_table"
  }
}

# Associate the public subnet to the public route table
resource "aws_route_table_association" "public_route_subnet_association" {
  subnet_id      = aws_subnet.k8s_public_subnet.id
  route_table_id = aws_route_table.public_route_table.id
}


# Create a private route table for the private subnet
resource "aws_route_table" "private_route_table" {
  vpc_id = aws_vpc.k8s_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    # the default route for the traffic originating 
    # from the private subnet goes to nat gateway 
    gateway_id = aws_nat_gateway.this.id
  }

  tags = {
    Name = "private_route_table"
  }
}

# Associate the private subnet to the private route table
resource "aws_route_table_association" "private_route_subnet_association" {
  subnet_id      = aws_subnet.k8s_private_subnet.id
  route_table_id = aws_route_table.private_route_table.id
}

#create an aws elastic ip for the NAT gateway
resource "aws_eip" "nat_gw_eip" {
  vpc        = true
  depends_on = [aws_internet_gateway.this]
}

# create the NAT gateway
resource "aws_nat_gateway" "this" {
  allocation_id     = aws_eip.nat_gw_eip.id
  subnet_id         = aws_subnet.k8s_public_subnet.id
  connectivity_type = "public"

  tags = {
    Name = "nat_gw"
  }
  # To ensure proper ordering, it is recommended to add an explicit dependency
  # on the Internet Gateway for the VPC.
  depends_on = [aws_internet_gateway.this]
}
