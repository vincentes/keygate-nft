variable "region" {
    description = "Default region for provider"
    type = string
    default = "us-east-1"
}

variable "ami" {
    description = "Amazon machine image to use for EC2 instances"
    type = string
    default = "ami-0e731c8a588258d0d"
}

variable "instance_type" {
    description = "Type of EC2 instance to launch"
    type = string
    default = "t2.micro"
}

variable "cloudflare_token" {
    description = "API token for Cloudflare"
    type = string
    sensitive = true
}

variable "cloudflare_zone_id" {
    description = "Cloudflare zone to use"
    type = string
    default = "e350bb55fbd9dfd268d47463724fb913"
}

variable "environment_name" {
    description = "Deployment environment (dev/staging/production)"
    type = string
    default = "dev"
}

variable "app_name" {
    description = "Name of the application"
    type = string
    default = "web-app"
}