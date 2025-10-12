#!/bin/bash
set -Ceu

NAME=atomisu-ocg-123456/sample

docker ps -a | grep $NAME | awk '{print $1}' | xargs -r docker stop
docker ps -a | grep $NAME | awk '{print $1}' | xargs -r docker rm