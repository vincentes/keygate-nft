
resource "cloudflare_record" "api" {
    zone_id = var.cloudflare_zone_id
    type = "CNAME"
    name = var.environment_name == "production" ? "api" : "${var.environment_name}-api"
    value = aws_lb.load_balancer.dns_name
    proxied = true
}

resource "cloudflare_record" "root" {
    count = var.environment_name == "production" ? 1 : 0
    zone_id = var.cloudflare_zone_id
    type = "CNAME"
    name = "@"
    value = aws_lb.load_balancer.dns_name
    proxied = true
}

resource "cloudflare_record" "www" {
    count = var.environment_name == "production" ? 1 : 0
    zone_id = var.cloudflare_zone_id
    type = "CNAME"
    name = "www"
    value = aws_lb.load_balancer.dns_name
    proxied = true
}