output "cloudwatch_log_group_arn" {
  value = aws_cloudwatch_log_group.app.arn
}

output "cloudwatch_log_group_name" {
  value = aws_cloudwatch_log_group.app.name
}
