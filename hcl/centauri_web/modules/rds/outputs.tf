output "aws_db_instance_host" {
  value = aws_db_instance.main.address
}

output "aws_db_instance_port" {
  value = aws_db_instance.main.port
}
