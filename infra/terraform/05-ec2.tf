resource "aws_key_pair" "sshkey" {
  public_key = file(var.ssh_public_key_path)
  key_name   = var.key_name
}

locals {
  # ec2_instances local variable contains all the configuration
  # attributes of the EC2 instances in a map form
  ec2_instances = {
    k8s_master_node = {
      ami      = var.master_node_ami
      key_name = var.key_name
      network_interface = {
        network_interface_id = aws_network_interface.master_node_iface.id
        device_index         = 0
      }
      instance_type        = var.master_node_instance_type
      user_data            = data.cloudinit_config.master_node.rendered
      iam_instance_profile = aws_iam_instance_profile.ssm_iam_profile.name
    }
    k8s_worker_node01 = {
      ami      = var.worker_node_ami
      key_name = var.key_name
      network_interface = {
        network_interface_id = aws_network_interface.worker_node01_iface.id
        device_index         = 0
      }
      instance_type        = var.worker_node_instance_type
      user_data            = data.cloudinit_config.k8s_worker_node01.rendered
      iam_instance_profile = aws_iam_instance_profile.ssm_iam_profile.name
    }
    k8s_worker_node02 = {
      ami      = var.worker_node_ami
      key_name = var.key_name
      network_interface = {
        network_interface_id = aws_network_interface.worker_node02_iface.id
        device_index         = 0
      }
      instance_type        = var.worker_node_instance_type
      user_data            = data.cloudinit_config.k8s_worker_node02.rendered
      iam_instance_profile = aws_iam_instance_profile.ssm_iam_profile.name
    }
    nginx_server = {
      ami      = var.nginx_server_ami
      key_name = var.key_name
      network_interface = {
        network_interface_id = aws_network_interface.nginx_server_iface.id
        device_index         = 0
      }
      instance_type        = var.nginx_server_instance_type
      user_data            = data.cloudinit_config.nginx_server.rendered
      iam_instance_profile = aws_iam_instance_profile.ssm_iam_profile.name
    }
  }
}


resource "aws_instance" "ec2_nodes" {
  for_each = local.ec2_instances
  ami      = each.value.ami
  key_name = each.value.key_name
  network_interface {
    network_interface_id = each.value.network_interface.network_interface_id
    device_index         = each.value.network_interface.device_index
  }
  instance_type        = each.value.instance_type
  user_data            = each.value.user_data
  iam_instance_profile = each.value.iam_instance_profile
  tags = {
    Name = "${each.key}"
  }
}
