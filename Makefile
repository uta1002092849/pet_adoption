# Variables
PROTO_FILE = petAdoption.proto

# Targets
.PHONY: generate_grpc_python generate_grpc_go generate_grpc

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