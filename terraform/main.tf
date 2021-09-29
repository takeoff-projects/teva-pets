
terraform {
  required_version = ">= 0.14"

  required_providers {
    # Cloud Run support was added on 3.3.0
    google = ">= 3.3"
  }
}
provider "google" {
  credentials = file(var.credentials)
  project = var.project_id
  region  = var.provider_region
}
resource "google_project_service" "run_api" {
  service = "run.googleapis.com"
  disable_on_destroy = true
}

resource "google_project_service" "firestore" {
  service = "firestore.googleapis.com"
  disable_on_destroy = false
}

resource "google_firestore_document" "pet_default" {
  collection  = "Pet"
  document_id = "pet-default"
  fields      = "{\"petname\":{\"stringValue\":\"Onyx\"}, \"owner\":{\"stringValue\":\"Olesia\"}, \"likes\":{\"integerValue\":\"50\"}, \"image\":{\"stringValue\":\"https://d3544la1u8djza.cloudfront.net/APHI/Blog/2020/07-23/How+Much+Does+It+Cost+to+Have+a+Cat+_+ASPCA+Pet+Insurance+_+black+cat+with+yellow+eyes+peeking+out-min.jpg\"}}"
}

resource "google_cloud_run_service" "run_service" {
  name = "go-pets-teva"
  location = var.provider_region

  template {
    spec {
      containers {
        env {
          name = "GOOGLE_CLOUD_PROJECT"
          value = var.project_id
        }
        env {
          name = "GOOGLE_APPLICATION_CREDENTIALS"
          value = var.credentials
        }
        image = "gcr.io/${var.project_id}/go-pets-teva:latest"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  # Waits for the Cloud Run API to be enabled
  depends_on = [google_project_service.run_api]
}
# Allow unauthenticated users to invoke the service
resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.run_service.name
  location = google_cloud_run_service.run_service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
# Display the service URL
output "service_url" {
  value = google_cloud_run_service.run_service.status[0].url
}