DBNAME = gadgethub_db
CONTAINER_NAME = postgres-gadgethub
IMAGE = postgres:12.0-alpine
PORTS = 5432:5432
ENV_VARS = POSTGRES_PASSWORD=MP@TEST123 -e POSTGRES_USER=root
POSTGRESQL_URL = "postgresql://root:MP@TEST123@localhost:5432/gadgethub_db?sslmode=disable"


postgres:
	@if [ "$$(docker ps -aq -f name=$(CONTAINER_NAME))" ]; then \
		echo "Container $(CONTAINER_NAME) exists"; \
		if [ "$$(docker ps -q -f name=$(CONTAINER_NAME))" ]; then \
			echo "Docker $(CONTAINER_NAME) is running"; \
		else \
			echo "Docker $(CONTAINER_NAME) is stopped"; \
			docker start $(CONTAINER_NAME); \
		fi; \
	else \
		docker run --name $(CONTAINER_NAME) -p $(PORTS) -e POSTGRES_PASSWORD=MP@TEST123 -e POSTGRES_USER=root -d $(IMAGE); \
	fi

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DBNAME)

dropdb:
	docker exec -it $(CONTAINER_NAME)  dropdb $(DBNAME)

migrationup:
	migrate -database $(POSTGRESQL_URL) -path db/migration up

migrationdown:
	migrate -database $(POSTGRESQL_URL) -path db/migration down