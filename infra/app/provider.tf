terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.60"
    }
  }
  backend "s3" {
    bucket         = "terraform-remote-state20240808170101034800000001"
    kms_key_id     = "43eaa8e3-6fe5-452a-88c9-fd8b25c0e07d"
    key            = "infra/app/terraform.tfstate"
    region         = "us-east-1"
    profile        = "marmita-admin"
    dynamodb_table = "marmita_app_remote_state"
    encrypt        = true
  }

  required_version = ">= 1.9.3"
}

provider "aws" {
  profile = module.env.aws_provider.profile
  region  = module.env.aws_provider.region

  default_tags {
    tags = {
      App        = "marmita"
      ManagedBy  = "Terraform"
      SourceCode = "infra/app/provider.tf"
    }
  }
}

