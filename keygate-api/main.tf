terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~> 3.0"
        }
        cloudflare = {
            source = "cloudflare/cloudflare"
            version = "~> 4.0"
        }
    }
}

provider "cloudflare" {
    api_token = var.cloudflare_token
}

provider "aws" {
    region = var.region
}