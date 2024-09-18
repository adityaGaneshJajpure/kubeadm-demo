resource "local_file" "ansible_inventory" {
  content  = <<-EOT
      [nginx_server]   
      ${aws_instance.ec2_nodes["nginx_server"].id}
      [master_node]
      ${aws_instance.ec2_nodes["k8s_master_node"].id}
      [worker_nodes]
      ${aws_instance.ec2_nodes["k8s_worker_node01"].id}
      ${aws_instance.ec2_nodes["k8s_worker_node02"].id}
    EOT
  filename = "../ansible/inventory"
}
