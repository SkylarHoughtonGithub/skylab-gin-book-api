#!/bin/bash
set -e

# Use the first argument as APP_NAME, or fall back to "myapp" if not provided
NAME=$1

echo "Building $NAME..."

# Build Docker Image
docker build -t skylab-book-chameleon .

echo "$NAME built successfully!"