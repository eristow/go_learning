output "db_instance_endpoint" {
  value = aws_db_instance.app_db.endpoint
}

output "db_instance_port" {
  value = aws_db_instance.app_db.port
}

output "db_instance_username" {
  value = aws_db_instance.app_db.username
}


output "db_connection_string" {
  value     = local.db_connection_string
  sensitive = true
}

output "ecs_cluster_name" {
  value = aws_ecs_cluster.go_learning_cluster.name
}

output "ecs_service_frontend_name" {
  value = aws_ecs_service.go_learning_frontend.name
}

output "ecs_service_backend_name" {
  value = aws_ecs_service.go_learning_backend.name
}

output "alb_dns_name" {
  description = "DNS name of the load balancer"
  value       = aws_lb.app_lb.dns_name
}
