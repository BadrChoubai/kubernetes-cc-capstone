terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

module "cc_eks" {
  source = "./modules/eks"
}

module "cc_ecr_repository" {
  source = "./modules/ecr"

  ecr_repo_name = "cc_ecr_repository"
}
