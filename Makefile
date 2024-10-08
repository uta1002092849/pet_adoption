# Variables
DBS_COMPOSE_FILE = create_cluster_dbs.yml
INIT_CMD = ./cockroach --host=roach1:26357 init --insecure
PROTO_FILE = petAdoption.proto
GRPC_FOLDER = generated_grpc_code

# Targets
.PHONY: up init down clean generate_grpc_python generate_grpc_go clean_grpc clean_grpc_python clean_grpc_go

# Start the CockroachDB cluster
up:
	@echo "Starting CockroachDB cluster..."
	@docker compose -f $(DBS_COMPOSE_FILE) up -d
	@echo "Cluster started. Run 'make init' to initialize it if this is the first time."

# Initialize the cluster
init: up
	@echo "Initializing the CockroachDB cluster...If an error about it being already initialized appears, you can ignore it."
	@echo "Sleeping for 5 seconds to allow the cluster to start..."
	@sleep 5
	@docker exec -it roach1 $(INIT_CMD)

# Initialize the cluster with predined tables
init-table: up
	@echo "Initializing the CockroachDB cluster with predefined tables..."
	@docker cp ./init.sql roach1:/cockroach/
	@docker exec -it roach1 ./cockroach sql --insecure --host=roach1:26257 -f /cockroach/init.sql
	@echo "Pet adoption tables initialized."

# Show tables in the cluster
show-tables:
	@echo "Showing tables in the CockroachDB cluster..."
	@docker exec -it roach1 ./cockroach sql --insecure --host=roach1:26257 -e "SELECT * FROM petadoptiondb.client_schema.pet;"

# Stop the cluster
down:
	@echo "Stopping the CockroachDB cluster..."
	@docker compose -f $(DBS_COMPOSE_FILE) down
	@echo "Cluster stopped."

generate_grpc_python:
	@echo "Don't run with sudo to avoid permission issues."
	@echo "Generating gRPC code in Python..."
	@mkdir -p $(GRPC_FOLDER)/python/
	@python -m grpc_tools.protoc -I. --python_out=$(GRPC_FOLDER)/python/ --grpc_python_out=$(GRPC_FOLDER)/python/ $(PROTO_FILE)
	@touch $(GRPC_FOLDER)/python/__init__.py
	@echo "gRPC code for python generated under /$(GRPC_PYTHON_FOLDER)/python."

generate_grpc_go:
	@echo "Installing protoc-gen-go..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@export PATH="$PATH:$(go env GOPATH)/bin"
	@echo "Generating gRPC code in Go..."
	@mkdir -p $(GRPC_FOLDER)
	@protoc --go_out=./$(GRPC_FOLDER) --go-grpc_out=./$(GRPC_FOLDER) $(PROTO_FILE)
	@echo "gRPC code for Go generated under $(GRPC_FOLDER)/go."

generate_grpc: generate_grpc_python generate_grpc_go

clean_grpc:
	@echo "Cleaning up generated gRPC code..."
	@rm -rf $(GRPC_FOLDER)
	@echo "gRPC code cleaned."

# Clean up containers, volumes, and networks
clean:
	@echo "Cleaning up the Docker Compose environment..."
	@docker compose -f $(DBS_COMPOSE_FILE) down -v
	@echo "Environment cleaned."

