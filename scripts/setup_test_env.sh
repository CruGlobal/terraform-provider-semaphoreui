#!/usr/bin/env bash

set -eu
SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

# Get the mysql docker container ID
MYSQL_CONTAINER_ID=$(docker container ls -a --filter "name=mysql" --filter "status=running" --format "{{.ID}}")

# Insert the api_token into the database
docker exec $MYSQL_CONTAINER_ID sh -c 'exec mysql --user=semaphore --password=semaphore --database=semaphore -e "INSERT INTO user__token (user_id, id, created) VALUES (1, \"'$SEMAPHOREUI_API_TOKEN'\", NOW());"' >/dev/null 2>&1