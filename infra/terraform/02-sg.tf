resource "aws_security_group" "public" {
  name        = "allow_web_traffic"
  description = "Allow web traffic"
  vpc_id      = aws_vpc.k8s_vpc.id


  ingress {
    description      = "HTTPS"
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  ingress {
    description      = "HTTP"
    from_port        = 80
    to_port          = 80
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "public"
  }
}

resource "aws_security_group" "master_node" {
  name        = "allow_ctrl_plane_traffic"
  description = "Allow inbound traffic needed for k8s control plane to operate"
  vpc_id      = aws_vpc.k8s_vpc.id

  ingress {
    description = "Weavenet control port"
    from_port   = 6783
    to_port     = 6783
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Weavenet data ports"
    from_port   = 6783
    to_port     = 6784
    protocol    = "udp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Allow all icmp traffic from the private subnet"
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Allow all icmp traffic from the public subnet"
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = [aws_subnet.k8s_public_subnet.cidr_block]
  }

  ingress {
    description = "HTTPS from public subnet"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_public_subnet.cidr_block]
  }

  ingress {
    description = "HTTP from public subnet"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_public_subnet.cidr_block]
  }

  ingress {
    description = "Kubernetes API server"
    from_port   = 6443
    to_port     = 6443
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "etcd server client API"
    from_port   = 2379
    to_port     = 2380
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Kubelet API"
    from_port   = 10250
    to_port     = 10250
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "kube_scheduler"
    from_port   = 10259
    to_port     = 10259
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "kube_controller_manager"
    from_port   = 10257
    to_port     = 10257
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "allow_ctrl_plane_traffic"
  }
}

resource "aws_security_group" "worker_node" {
  name        = "allow_worker_nodes_traffic"
  description = "Allow inbound traffic needed for the worker nodes to operate"
  vpc_id      = aws_vpc.k8s_vpc.id

  ingress {
    description = "Allow all icmp traffic"
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Weavenet control port"
    from_port   = 6783
    to_port     = 6783
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Weavenet data ports"
    from_port   = 6783
    to_port     = 6784
    protocol    = "udp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "Kubelet API"
    from_port   = 10250
    to_port     = 10250
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }

  ingress {
    description = "NodePort Services"
    from_port   = 30000
    to_port     = 32767
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_public_subnet.cidr_block]
  }

  ingress {
    description = "NodePort Services"
    from_port   = 30000
    to_port     = 32767
    protocol    = "tcp"
    cidr_blocks = [aws_subnet.k8s_private_subnet.cidr_block]
  }



  # The rule below allows kubernetes master node to ssh to worker nodes. 
  # The k8s cluster itself does not need this rule to operate. This rule was
  # added, to work on a CKA exam scenario where I needed to perform etcd backup
  # using the etcdctl from one of the worker nodes which means that the etcd 
  # pki certificates had to be copied from the master node to the worker node 
  # using scp (scp's default port is tcp 22) 
  ingress {
    description = "Allow ssh only from the control plane node of the cluster"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = formatlist("%s/32", var.master_node_ip_address)
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "allow_worker_nodes_traffic"
  }
}
