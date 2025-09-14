#!/bin/bash

echo "Initializing local Airflow"

cd ./pipeline

./mwaa-local-env reset-db
./mwaa-local-env start

echo "Local Airflow initialized"
