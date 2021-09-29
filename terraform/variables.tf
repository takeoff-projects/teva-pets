variable "project_id" {
    type        = string
    description = "Google Cloud Platform Project ID"
    default = "roi-takeoff-user64"
}

variable "provider_region" {
  type    = string
  default = "us-central1"
}

variable "credentials" {
  type    = string
}