#!/bin/bash
set -e

# Use the first argument as APP_NAME, or fall back to "myapp" if not provided
NAME=$1

echo "Building $NAME..."

# Build binary
go build -o $NAME ./cmd/server

echo "$NAME built successfully!"