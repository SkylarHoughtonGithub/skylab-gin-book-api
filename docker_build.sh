#!/bin/bash
set -e

NAME=$1

echo "Building $NAME..."

# Build Docker Image
docker build -t skylab-book-chameleon .

echo "$NAME built successfully!"