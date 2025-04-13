variable "aws_region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "us-east-1"
}

variable "db_username" {
  description = "The username for the database"
  type        = string
  default     = "docker"
}

variable "db_password" {
  description = "The password for the database"
  type        = string
  sensitive   = true
}

# variable "gitlab_deploy_token_username" {
#   description = "GitLab deploy token username for accessing the repository"
#   type        = string
#   sensitive   = true
# }

# variable "gitlab_deploy_token_password" {
#   description = "GitLab deploy token password for accessing the repository"
#   type        = string
#   sensitive   = true
# }
