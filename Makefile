# Variables
DBS_COMPOSE_FILE = create_cluster_dbs.yml
INIT_CMD = ./cockroach --host=roach1:26357 init --insecure

# Targets
.PHONY: up init down clean

# Start the CockroachDB cluster
up:
	@echo "Starting CockroachDB cluster..."
	@docker compose -f $(DBS_COMPOSE_FILE) up -d
	@echo "Cluster started. Run 'make init' to initialize it if this is the first time."

# Initialize the cluster
init: up
	@echo "Initializing the CockroachDB cluster...If an error about it being already initialized appears, you can ignore it."
	@docker exec -it roach1 $(INIT_CMD)

# Initialize the cluster with predined tables
init-table: up
	@echo "Initializing the CockroachDB cluster with predefined tables..."
	@docker cp ./init.sql roach1:/cockroach/
	@docker exec -it roach1 ./cockroach sql --insecure --host=roach1:26257 -f /cockroach/init.sql
	@echo "Pet adoption tables initialized."

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

