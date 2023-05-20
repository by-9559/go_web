#!/bin/bash

IMAGE_NAME="war7ng/go-web"
UTC_DATE_TIME=$(date +%Y%m%d-%H%M%S)
IMAGE_VERSION="v2-${UTC_DATE_TIME}"

go build main.go

echo "Building Docker image ${IMAGE_NAME}:${IMAGE_VERSION}"
docker build -t "${IMAGE_NAME}:${IMAGE_VERSION}" .

echo "Pushing Docker image ${IMAGE_NAME}:${IMAGE_VERSION}"
docker push "${IMAGE_NAME}:${IMAGE_VERSION}"
echo "${IMAGE_VERSION}"
