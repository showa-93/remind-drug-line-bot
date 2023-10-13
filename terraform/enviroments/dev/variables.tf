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
  default = "{{DOCKER_REPOSITORY}}@{{DIGEST}}"
}

variable "api_domain" {
  type    = string
  default = "api.showa.dev"
}
