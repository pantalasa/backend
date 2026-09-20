terraform {
  required_version = ">= 1.4.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

module "transactions" {
  source = "git::ssh://git@github.com/pantalasa/transactions-s3.git//terraform?ref=v1.0.0"

  name_prefix = "pantalasa-backend"
  tags = {
    Service     = "backend"
    Environment = "production"
    Compliance  = "SOC2"
  }
}

resource "aws_dynamodb_table" "api_rate_limits" {
  name         = "pantalasa-backend-rate-limits"
  billing_mode = "PAY_PER_REQUEST"

  deletion_protection_enabled = true

  attribute {
    name = "client_id"
    type = "S"
  }

  hash_key = "client_id"

  lifecycle {
    prevent_destroy = true
  }

  tags = {
    Service     = "backend"
    Environment = "production"
  }
}
