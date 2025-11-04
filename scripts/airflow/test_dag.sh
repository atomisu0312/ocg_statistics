#!/bin/bash
set -Ceu
docker exec aws-mwaa-local-runner-2_10_3-local-runner-1 pytest /usr/local/airflow/dags/test/test_db_insert.py
