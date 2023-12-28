# filename: main.tf
#####################################################################
# providers
#####################################################################
provider "aws" {
  region = var.aws_region
  default_tags {
    tags = local.tags
  }
  dynamic "assume_role" {
    for_each = var.aws_role_assume ? tolist(["1"]) : []
    content {
      role_arn = var.aws_role
    }
  }
}

#####################################################################
# backend
#####################################################################
terraform {
  backend "s3" {
    bucket               = "templater-example-state"
    key                  = "terraform.tfstate"
    region               = "us-east-1"
    workspace_key_prefix = "test"
    encrypt              = true
  }
}