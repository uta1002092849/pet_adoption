# Variables
PROTO_FILE = petAdoption.proto
DBS_COMPOSE_FILE = docker-compose.yml
INIT_CMD = ./cockroach --host=roach1:26357 init --insecure
PROTO_FILE = petAdoption.proto
GRPC_FOLDER = generated_grpc_code

# Targets
.PHONY: generate_grpc_python generate_grpc_go generate_grpc db-up db-init db-table-init db-table-show down clean service-up
# Start the CockroachDB cluster
db-up:
	@echo "Starting CockroachDB cluster..."
	@docker compose -f $(DBS_COMPOSE_FILE) up roach1 roach2 roach3 -d

# Initialize the cluster
db-init:
	@echo "Initializing the CockroachDB cluster..."
	@docker exec -it roach1 $(INIT_CMD)

# Initialize the cluster with predined tables
db-table-init:
	@echo "Initializing the CockroachDB cluster with predefined tables..."
	@docker cp ./init.sql roach1:/cockroach/
	@docker exec -it roach1 ./cockroach sql --insecure --host=roach1:26257 -f /cockroach/init.sql
	@echo "Pet adoption tables initialized."

# Show tables in the cluster
db-table-show:
	@echo "Showing tables in the CockroachDB cluster..."
	@docker exec -it roach1 ./cockroach sql --insecure --host=roach1:26257 -e "SELECT * FROM petadoptiondb.client_schema.pet;"

# Stop the cluster
down:
	@echo "Stopping the CockroachDB cluster..."
	@docker compose -f $(DBS_COMPOSE_FILE) down
	@echo "Cluster stopped."

# Clean up containers, volumes, and networks
clean:
	@echo "Cleaning up the Docker Compose environment..."
	@docker compose -f $(DBS_COMPOSE_FILE) down -v
	@echo "Environment cleaned."

generate_grpc_python:
	@echo "Generating gRPC code in Python..."
	@mkdir -p ./server
	@python -m grpc_tools.protoc -I. --python_out=./server/ --grpc_python_out=./server/ $(PROTO_FILE)
	@echo "gRPC code for python generated under ./server"

generate_grpc_go:
	@echo "Make sure you have all independent for go grpc installed first!"
	# @go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	# @go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	# @export PATH="$PATH:$(go env GOPATH)/bin"
	@echo "Generating gRPC code in Go..."
	@mkdir -p ./client
	@protoc --go_out=./client --go-grpc_out=./client $(PROTO_FILE)
	@echo "gRPC code for Go generated under ./client"

generate_grpc: generate_grpc_python generate_grpc_go

service-up:
	@echo "Starting the client and server..."
	@docker compose -f $(DBS_COMPOSE_FILE) up server client -d
	@echo "Web interface available at http://localhost:8081"

service-down:
	@echo "Stopping the client and server..."
	@docker compose -f $(DBS_COMPOSE_FILE) down server client
	@echo "Client and server stopped."

