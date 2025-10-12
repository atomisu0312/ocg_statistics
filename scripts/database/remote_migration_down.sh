#!/bin/bash
set -Ceu

# 環境変数を読み込む
set -a
source .env
set +a

cd database/migrate
DATABASE_URL=${NEONDB_CONNECTION_STRING}
migrate -path db/migrations -database $DATABASE_URL down