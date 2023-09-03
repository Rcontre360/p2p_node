
# Variables
APP_NAME := p2p-example
GO_FILES := $(wildcard *.go)
BUILD_DIR := bin
SRC_DIR := .

# Build the Go application
build: $(APP_NAME)

$(APP_NAME): $(GO_FILES)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Run the Go application
run:
	$(BUILD_DIR)/$(APP_NAME)

# Clean up build artifacts
clean:
	rm -rf $(BUILD_DIR)

codegen:
	protoc --go_out=$(SRC_DIR) --go-grpc_out=$(SRC_DIR) rpc/proto/schema.proto

