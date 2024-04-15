terraform {
  backend "s3" {
    bucket         = "devops-directive-terraform-state-2"
    key            = "keygate/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform_locks_v2"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

provider "cloudflare" {
  api_token = var.cloudflare_token
}

variable "cloudflare_token" {
  description = "API token for Cloudflare"
  type        = string
  sensitive   = true
}

variable "github_private_key" {
  description = "Private SSH key for accessing the GitHub repository"
  type        = string
  sensitive   = true
}

locals {
  environment_name = terraform.workspace
}

module "keygate-api" {
  source           = "../keygate-api-tf"
  app_name         = "${local.environment_name}-keygate-api"
  cloudflare_token = var.cloudflare_token
  environment_name = local.environment_name
  region           = "us-east-1"
  ami              = "ami-0e731c8a588258d0d"
  instance_type    = "t2.micro"
  github_private_key = var.github_private_key
}

output "private_key" {
  value     = module.keygate-api.private_key
  sensitive = true
}