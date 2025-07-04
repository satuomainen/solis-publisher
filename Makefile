# Go parameters
GO=go
# GOFLAGS=-ldflags="-s -w"
GOFLAGS=

# Directories
CMD_DIR=./cmd

# Find all directories containing main.go files
CMD_DIRS=$(shell find $(CMD_DIR) -type f -name 'main.go' -exec dirname {} \; | sort -u)

# Build targets for each directory
BUILD_TARGETS=$(addprefix build-,$(CMD_DIRS))

# Default target
all: $(BUILD_TARGETS)

# Build target for each directory
$(BUILD_TARGETS): build-%:
	@echo "Building $*..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build $(GOFLAGS) -o out/$* ./$*

.PHONY: all $(BUILD_TARGETS)

clean:
	rm -r out
