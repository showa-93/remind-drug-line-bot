locals {
  services = toset([
    "iam.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "iamcredentials.googleapis.com",
    "sts.googleapis.com"
  ])
}

resource "google_project_service" "enable_api" {
  for_each                   = local.services
  project                    = var.project
  service                    = each.value
  disable_dependent_services = true
}

resource "google_iam_workload_identity_pool" "bot_pool" {
  provider                  = google-beta
  project                   = var.project
  workload_identity_pool_id = "botpool"
  display_name              = "botpool"
  description               = "For GitHub Actions"
}

resource "google_iam_workload_identity_pool_provider" "bot_provider" {
  provider                           = google-beta
  project                            = var.project
  workload_identity_pool_id          = google_iam_workload_identity_pool.bot_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "botprovider"
  display_name                       = "botprovider"
  description                        = "For GitHub Actions"

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
  }

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

data "google_service_account" "terraform_sa" {
  account_id = var.tf_service_account
}

resource "google_service_account_iam_member" "terraform_sa" {
  service_account_id = data.google_service_account.terraform_sa.id
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.bot_pool.name}/attribute.repository/${var.repository}"
}
