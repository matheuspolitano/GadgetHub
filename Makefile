DBNAME = gadgethub_db
CONTAINER_NAME = postgres-gadgethub
IMAGE = postgres:12.0-alpine
PORTS = 5432:5432
ENV_VARS = POSTGRES_PASSWORD=secret -e POSTGRES_USER=root
POSTGRESQL_URL = "postgresql://root:secret@localhost:5432/gadgethub_db?sslmode=disable"


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
		docker run --name $(CONTAINER_NAME) -p $(PORTS) -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d $(IMAGE); \
	fi

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DBNAME)

dropdb:
	docker exec -it $(CONTAINER_NAME)  dropdb $(DBNAME)

migrateup:
	migrate -verbose -database $(POSTGRESQL_URL) -path pkg/db/migration up

migratedown:
	migrate -database $(POSTGRESQL_URL) -path pkg/db/migration down

sqlc:
	sqlc generate

test:
	go test -v ./...

proto1:
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

relay1:
	docker run -d --network host webhookrelay/webhookrelayd:latest \
	-k b9d0c5fd-31ac-4d33-9106-33e2e64d7134 \
	-s crs2eWyuOwgO \
	-b testMeta

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

run:
	go run main.go	