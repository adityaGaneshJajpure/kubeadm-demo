resource "aws_network_interface" "master_node_iface" {
  subnet_id       = aws_subnet.k8s_private_subnet.id
  private_ips     = [var.master_node_ip_address]
  security_groups = [aws_security_group.master_node.id]

  tags = {
    Name = "master_node_iface"
  }
}

resource "aws_network_interface" "worker_node01_iface" {
  subnet_id       = aws_subnet.k8s_private_subnet.id
  private_ips     = [var.worker_node01_ip_address]
  security_groups = [aws_security_group.worker_node.id]

  tags = {
    Name = "worker_node01_iface"
  }
}

resource "aws_network_interface" "worker_node02_iface" {
  subnet_id       = aws_subnet.k8s_private_subnet.id
  private_ips     = [var.worker_node02_ip_address]
  security_groups = [aws_security_group.worker_node.id]

  tags = {
    Name = "worker_node02_iface"
  }
}

# create the interface of nginx server
resource "aws_network_interface" "nginx_server_iface" {
  subnet_id       = aws_subnet.k8s_public_subnet.id
  security_groups = [aws_security_group.public.id]

  tags = {
    Name = "nginx_server_iface"
  }
}
