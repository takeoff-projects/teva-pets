variable "project_id" {
    type        = string
    description = "Google Cloud Platform Project ID"
}

variable "provider_region" {
  type    = string
  default = "us-central1"
}

variable "credentials" {
  type    = string
}