#!/bin/bash
set -e

NAME=$1

echo "Building $NAME..."

# Build binary
go build -o $NAME ./cmd/server

echo "$NAME built successfully!"