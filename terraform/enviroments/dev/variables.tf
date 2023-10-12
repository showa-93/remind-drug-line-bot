variable "region" {
  type    = string
  default = "{{REGION}}"
}

variable "zone" {
  type    = string
  default = "{{ZONE}}"
}

variable "project" {
  type    = string
  default = "{{PROJECT_ID}}"
}

variable "api_image" {
  type    = string
  default = "{{DOCKER_REPOSITORY}}@{{DIGEST}}"
}

variable "api_domain" {
  type    = string
  default = "api.showa.dev"
}
