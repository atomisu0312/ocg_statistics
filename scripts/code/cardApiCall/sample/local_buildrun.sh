#!/bin/bash
set -Ceu

NAME=atomisu-ocg-123456/sample
ARCHITECTURE=linux/arm64 # linux or Macの場合
# ARCHITECTURE=linux/amd64 # Windowsの場合

cd ./code/cardApiCall/sample
docker buildx build --platform $ARCHITECTURE --provenance=false -t $NAME:latest .

# ローカルでrunする場合には、entrypointを指定する必要がある
docker run -d -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie $NAME:latest ./main