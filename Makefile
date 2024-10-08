# Variables
PROTO_FILE = petAdoption.proto

# Targets
.PHONY: generate_grpc_python generate_grpc_go generate_grpc

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