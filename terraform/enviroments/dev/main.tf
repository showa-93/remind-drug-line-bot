# ローカル変数
locals {
  env = "dev"
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

module "api" {
  source     = "../../services/api"
  region     = var.region
  project    = var.project
  api_image  = var.api_image
  api_domain = var.api_domain
}
