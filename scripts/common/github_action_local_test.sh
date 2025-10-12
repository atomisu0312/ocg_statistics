#!/bin/bash
set -Ceu

ARCHITECTURE=linux/amd64
SECRET_FILE=act_secretfile
WORKFLOW_FILE=.github/workflows/lamda-sample-deploy.yml

# actがインストールされているかチェックする
if ! command -v act &> /dev/null; then
    echo "act could not be found. Please install it first."
    exit 1
fi

act --secret-file $SECRET_FILE --container-architecture $ARCHITECTURE --workflows $WORKFLOW_FILE