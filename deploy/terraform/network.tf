# Partner settlement systems call the quotes API over the private link
# instead of the public endpoint.
variable "api_security_group_id" {
  description = "Security group attached to the quotes API load balancer"
  type        = string
  default     = "sg-0b6d1c2f4e8a93d75"
}

variable "partner_cidr" {
  description = "CIDR block of the partner network"
  type        = string
  default     = "203.0.113.0/24"
}

resource "aws_vpc_security_group_ingress_rule" "quotes_api_partner_https" {
  security_group_id = var.api_security_group_id
  description       = "HTTPS from the partner network"
  cidr_ipv4         = var.partner_cidr
  ip_protocol       = "tcp"
  from_port         = 443
  to_port           = 443

  tags = {
    Service     = "backend"
    Environment = "production"
  }
}
