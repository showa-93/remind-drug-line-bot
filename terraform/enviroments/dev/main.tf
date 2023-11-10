# ローカル変数
locals {
  env = "dev"
}

locals {
  services = toset([
    "cloudresourcemanager.googleapis.com",
    "logging.googleapis.com",
    "iam.googleapis.com",
    "secretmanager.googleapis.com",
    "storage.googleapis.com",
    "artifactregistry.googleapis.com",
    "run.googleapis.com",
    "firestore.googleapis.com"
  ])
}

resource "google_project_service" "enable_api" {
  for_each                   = local.services
  project                    = var.project
  service                    = each.value
  disable_dependent_services = true
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

module "workload" {
  source             = "../../resources/workload"
  repository         = var.repository
  project            = var.project
  region             = var.region
  zone               = var.zone
  tf_service_account = var.tf_service_account
}

module "firestore" {
  source  = "../../resources/firestore"
  project = var.project
  region  = var.region
}

module "api" {
  source     = "../../services/api"
  region     = var.region
  project    = var.project
  api_image  = var.api_image
  api_domain = var.api_domain
}
