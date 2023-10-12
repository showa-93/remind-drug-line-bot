resource "google_cloud_run_v2_service" "api" {
  name     = "api"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    max_instance_request_concurrency = 100
    timeout                          = "30s"

    scaling {
      min_instance_count = 0
      max_instance_count = 2
    }

    containers {
      image = var.api_image
      startup_probe {
        failure_threshold = 1
        period_seconds    = 240
        timeout_seconds   = 240
        tcp_socket {
          port = 8080
        }
      }
      resources {
        cpu_idle          = true
        limits            = { "memory" : "512Mi", "cpu" : "1" }
        startup_cpu_boost = true
      }
    }
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers"
    ]
  }
}

resource "google_cloud_run_v2_service_iam_policy" "noauth" {
  location    = google_cloud_run_v2_service.api.location
  project     = google_cloud_run_v2_service.api.project
  name        = google_cloud_run_v2_service.api.name
  policy_data = data.google_iam_policy.noauth.policy_data
}

resource "google_cloud_run_domain_mapping" "domain-mapping" {
  location = var.region
  name     = var.api_domain

  metadata {
    namespace = var.project
  }

  spec {
    route_name = google_cloud_run_v2_service.api.name
  }
}
