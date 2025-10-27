variable "authn_image" {
  type    = string
  default = "pulap-authn:latest"
}

variable "authz_image" {
  type    = string
  default = "pulap-authz:latest"
}

variable "estate_image" {
  type    = string
  default = "pulap-estate:latest"
}

variable "admin_image" {
  type    = string
  default = "pulap-admin:latest"
}

job "pulap-services" {
  datacenters = ["dc1"]
  type        = "service"

  group "services" {
    count = 1

    restart {
      attempts = 5
      interval = "30m"
      delay    = "30s"
      mode     = "delay"
    }

    network {
      mode = "host"
    }

    task "authn" {
      driver = "docker"

      config {
        image        = var.authn_image
        network_mode = "host"
      }

      env {
        AUTHN_SERVER_PORT               = ":8082"
        AUTHN_DATABASE_PATH             = "/data/authn.db"
        AUTHN_DATABASE_MONGO_URL        = "mongodb://admin:password@127.0.0.1:27017/authn?authSource=admin"
        AUTHN_DATABASE_MONGO_DATABASE   = "authn"
      }

      resources {
        cpu    = 300
        memory = 256
      }
    }

    task "authz" {
      driver = "docker"

      config {
        image        = var.authz_image
        network_mode = "host"
      }

      env {
        AUTHZ_SERVER_PORT             = ":8083"
        AUTHZ_DATABASE_PATH           = "/data/authz.db"
        AUTHZ_DATABASE_MONGO_URL      = "mongodb://admin:password@127.0.0.1:27017/authz?authSource=admin"
        AUTHZ_DATABASE_MONGO_DATABASE = "authz"
        AUTHZ_AUTHN_URL               = "http://127.0.0.1:8082"
      }

      resources {
        cpu    = 300
        memory = 256
      }
    }

    task "estate" {
      driver = "docker"

      config {
        image        = var.estate_image
        network_mode = "host"
      }

      env {
        ESTATE_SERVER_PORT = ":8084"
        ESTATE_DATABASE_PATH = "/data/estate.db"
      }

      resources {
        cpu    = 300
        memory = 256
      }
    }

    task "admin" {
      driver = "docker"

      config {
        image        = var.admin_image
        network_mode = "host"
      }

      env {
        ADMIN_SERVER_PORT            = ":8081"
        ADMIN_SERVICES_AUTHN_URL     = "http://127.0.0.1:8082"
        ADMIN_SERVICES_AUTHZ_URL     = "http://127.0.0.1:8083"
        ADMIN_SERVICES_ESTATE_URL    = "http://127.0.0.1:8084"
        ADMIN_AUTH_SESSION_SECRET    = "change-this-in-production"
      }

      resources {
        cpu    = 300
        memory = 256
      }
    }
  }
}
