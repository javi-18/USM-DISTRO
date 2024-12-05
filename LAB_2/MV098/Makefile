# Variables
DOCKER_COMPOSE = docker-compose

# Construir las imágenes Docker
build:
	$(DOCKER_COMPOSE) up --build

# Levantar los contenedores
up:
	$(DOCKER_COMPOSE) up

# Detener los contenedores
down:
	$(DOCKER_COMPOSE) down

# Limpiar los contenedores y volúmenes de Docker
clean:
	$(DOCKER_COMPOSE) down --volumes --rmi all
