#!/bin/bash
set -Ceu
docker exec -it aws-mwaa-local-runner-2_10_3-local-runner-1 airflow dags list