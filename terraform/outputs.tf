output "alb_url" {
  description = "Application URL — test with curl after pipeline runs"
  value       = "http://${aws_lb.main.dns_name}"
}

output "ecr_repository_url" {
  description = "Push bootstrap image here before pipeline first runs"
  value       = aws_ecr_repository.app.repository_url
}

output "ecs_task_execution_role_arn" {
  description = "Paste into pipeline/taskdef.json → executionRoleArn"
  value       = aws_iam_role.ecs_task_execution.arn
}

output "ecs_task_role_arn" {
  description = "Paste into pipeline/taskdef.json → taskRoleArn"
  value       = aws_iam_role.ecs_task.arn
}

output "codestar_connection_arn" {
  description = "Authorize this in AWS console after apply (PENDING → AVAILABLE)"
  value       = aws_codestarconnections_connection.github.arn
}

output "codestar_connection_status" {
  description = "Must be AVAILABLE before pipeline can trigger from GitHub"
  value       = aws_codestarconnections_connection.github.connection_status
}

output "codepipeline_name" {
  value = aws_codepipeline.main.name
}

output "ecs_cluster_name" {
  value = aws_ecs_cluster.main.name
}

output "ecs_service_name" {
  value = aws_ecs_service.app.name
}
