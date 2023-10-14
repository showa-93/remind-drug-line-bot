variable "repository" {
  type = string
}

variable "region" {
  type = string
}

variable "zone" {
  type = string
}

variable "project" {
  type = string
}

variable "api_image" {
  type    = string
  default = "{{DOCKER_REPOSITORY}}/api@{{DIGEST}}"
}

variable "api_domain" {
  type    = string
  default = "api.showa.dev"
}

variable "tf_service_account" {
  type = string
}
