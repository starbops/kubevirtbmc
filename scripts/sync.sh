#!/bin/bash

set -euo pipefail

# Load .config.env
CONFIG_FILE="$(dirname "$0")/../.config.env"
if [[ -f "$CONFIG_FILE" ]]; then
    set -a
    source "$CONFIG_FILE"
    set +a
else
    echo ".config.env not found!"
    exit 1
fi

# Derived image tags
VIRTBMC_IMAGE_FULL="${REGISTRY_HOST}:${REGISTRY_PORT}/${VIRTBMC_IMAGE_NAME}:${VIRTBMC_IMAGE_TAG}"
CONTROLLER_IMAGE_FULL="${REGISTRY_HOST}:${REGISTRY_PORT}/${CONTROLLER_IMAGE_NAME}:${CONTROLLER_IMAGE_TAG}"

echo "Building Docker images..."
docker build -t "$VIRTBMC_IMAGE_FULL" -f Dockerfile.virtbmc .
docker build -t "$CONTROLLER_IMAGE_FULL" -f Dockerfile .

echo "Pushing images to kubevirtci registry..."
docker push "$VIRTBMC_IMAGE_FULL"
docker push "$CONTROLLER_IMAGE_FULL"


echo "Images built and loaded successfully:"
echo " - $VIRTBMC_IMAGE_FULL"
echo " - $CONTROLLER_IMAGE_FULL"
