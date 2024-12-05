DOCKER_COMPOSE_VM097 := docker-compose.vm097.yml
DOCKER_COMPOSE_VM098 := docker-compose.vm098.yml
DOCKER_COMPOSE_VM099 := docker-compose.vm099.yml
DOCKER_COMPOSE_VM100 := docker-compose.vm100.yml

.PHONY: buildvm097 buildvm098 buildvm099 buildvm100 startvm097 startvm098 startvm099 startvm100 logsvm097 logsvm098 logsvm099 logsvm100 stopvm097 stopvm098 stopvm099 stopvm100

buildvm097:
	@echo "Construyendo imagen para VM097..."
	docker-compose -f $(DOCKER_COMPOSE_VM097) build

buildvm098:
	@echo "Construyendo imagen para VM098..."
	docker-compose -f $(DOCKER_COMPOSE_VM098) build

buildvm099:
	@echo "Construyendo imagen para VM099..."
	docker-compose -f $(DOCKER_COMPOSE_VM099) build

buildvm100:
	@echo "Construyendo imagen para VM100..."
	docker-compose -f $(DOCKER_COMPOSE_VM100) build

startvm097: buildvm097
	@echo "Desplegando Broker en VM097..."
	docker-compose -f $(DOCKER_COMPOSE_VM097) up -d

startvm098: buildvm098
	@echo "Desplegando Jayce en VM098..."
	docker-compose -f $(DOCKER_COMPOSE_VM098) up -d

startvm099: buildvm099
	@echo "Desplegando Supervisores en VM099..."
	docker-compose -f $(DOCKER_COMPOSE_VM099) up -d

startvm100: buildvm100
	@echo "Desplegando los Servidores en VM100..."
	docker-compose -f $(DOCKER_COMPOSE_VM100) up -d

logsvm097:
	@echo "Mostrando logs de VM097..."
	docker-compose -f $(DOCKER_COMPOSE_VM097) logs -f

logsvm098:
	@echo "Mostrando logs de VM098..."
	docker-compose -f $(DOCKER_COMPOSE_VM098) logs -f

logsvm099:
	@echo "Mostrando logs de VM099..."
	docker-compose -f $(DOCKER_COMPOSE_VM099) logs -f

logsvm100:
	@echo "Mostrando logs de VM100..."
	docker-compose -f $(DOCKER_COMPOSE_VM100) logs -f

stopvm097:
	@echo "Deteniendo todos los contenedores de VM097..."
	docker-compose -f $(DOCKER_COMPOSE_VM097) down

stopvm098:
	@echo "Deteniendo todos los contenedores de VM098..."
	docker-compose -f $(DOCKER_COMPOSE_VM098) down

stopvm099:
	@echo "Deteniendo todos los contenedores de VM099..."
	docker-compose -f $(DOCKER_COMPOSE_VM099) down

stopvm100:
	@echo "Deteniendo todos los contenedores de VM100..."
	docker-compose -f $(DOCKER_COMPOSE_VM100) down


