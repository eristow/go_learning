terraform {
  backend "s3" {
    bucket       = "eristow-go-learning-terraform-state"
    key          = "go_learning.tfstate"
    region       = "us-east-1"
    use_lockfile = true
    encrypt      = true
  }

  required_version = "~> 1.11.2"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}
