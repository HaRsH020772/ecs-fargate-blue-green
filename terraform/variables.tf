variable "project_name" {
  description = "Used for naming all resources"
  type        = string
  default     = "ecs-blue-green"
}

variable "region" {
  type    = string
  default = "ap-south-1"
}

variable "vpc_cidr" {
  type    = string
  default = "10.0.0.0/16"
}

variable "az_count" {
  type    = number
  default = 2
}

variable "container_name" {
  description = "Must match the name in pipeline/appspec.yaml"
  type        = string
  default     = "app"
}

variable "container_port" {
  type    = number
  default = 8080
}

variable "github_repo" {
  description = "GitHub repo in owner/repo format"
  type        = string
}

variable "github_branch" {
  type    = string
  default = "main"
}
