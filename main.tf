terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "2.21.0"
    }
  }
}

provider "docker" {
    host = "tcp://localhost:8090"
}

resource "docker_image" "database" {
  name = "postgres:alpine"
  keep_locally = false
}

resource "docker_image" "api" {
  name = "matheusfenolio/go-api:main"
  keep_locally = false
}

resource "docker_network" "network" {
    name = "go_api_net"
}

resource "docker_container" "database" {
  image = docker_image.database.image_id
  name = "postgres_go_api"
  env = [
    "POSTGRES_USER=api_user",
    "POSTGRES_PASSWORD=4p!User",
    "POSTGRES_DB=go_api"
  ]
  ports {
    internal = 5432
    external = 5432
  }
  networks_advanced {
    name = docker_network.network.id
  }
  provisioner "local-exec" {
      command = "echo Waiting database to be ready && sleep 10"
  }
}

resource "docker_container" "api" {
    image = docker_image.api.image_id
    name = "go_api"
    env = [
        "HOST=${docker_container.database.name}",
        "USERNAME=api_user",
        "PASSWORD=4p!User",
        "DATABSE=go_api",
        "DB_PORT=5432",
        "SDSLMODE=disable",
        "TIMEZONE=UTC"
    ]
    ports {
        internal = 8080
        external = 8080
    }
    networks_advanced {
        name = docker_network.network.id
    }
    depends_on = [
      docker_container.database
    ]
}