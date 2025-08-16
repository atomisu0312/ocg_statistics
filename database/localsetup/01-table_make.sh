echo "### initialize start ####"
set PGCLIENTENCODING=utf-8
chcp 65001
set PGPASSWORD=postgres
set PGUSER=postgres

# SQLファイルの実行
#psql -U postgres -d postgres -f /docker-entrypoint-initdb.d/sql/02-init.sql