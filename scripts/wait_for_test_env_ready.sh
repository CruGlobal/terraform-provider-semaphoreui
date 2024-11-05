#!/usr/bin/env bash

set -eu

# get the container name
SEMAPHORE_CONTAINER_ID=$(docker container ls -a --filter "name=semaphore" --filter "status=running" --format "{{.ID}}")

echo "Waiting for the Semaphore container $SEMAPHORE_CONTAINER_ID to be ready..."
until [ "$(docker inspect -f {{.State.Health.Status}} "$SEMAPHORE_CONTAINER_ID")" == "healthy" ]; do
    sleep 1;
    printf "."
done;
echo "Semaphore container is ready!"