terraform {
  required_version = ">= 1.2.9"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.20.1"
    }
  }

  backend "s3" {
    bucket = "${var.tf_state_bucket}"
    key    = "${var.application}/${var.environment}"
    region = "${var.region}"
  }
}
