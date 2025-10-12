#!/bin/bash
set -Ceu

cd database/migrate
DATABASE_URL="postgres://postgres:postgres@localhost:5555/postgres?sslmode=disable"
migrate -path db/migrations -database $DATABASE_URL up