terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~> 5.60"
        }
    }
    backend "s3" {
        bucket = "terraform-remote-state20240730194456941500000001"
        kms_key_id = "cee1ea76-5ff9-452f-9aa9-83f807e40c77"
        key = "infra/remote_state/terraform.tfstate"
        region = "us-east-1"
        profile = "marmita-admin"
        dynamodb_table = "marmita_app_remote_state"
        encrypt = true
    }

    required_version = ">= 1.9.3"
}

provider "aws" {
    profile = module.env.aws_provider.profile
    region = module.env.aws_provider.region
}

