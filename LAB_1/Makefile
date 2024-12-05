DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = docker-compose.yaml

build:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) build

up: build
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d

down:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

.PHONY: build up down