echo "### initialize start ####"
set PGCLIENTENCODING=utf-8
chcp 65001
set PGPASSWORD=postgres
set PGUSER=postgres
set TEST_DB_NAME=test_ocg_statics

psql -U postgres --command="CREATE DATABASE test_ocg_statics LC_COLLATE 'ja_JP.UTF-8' LC_CTYPE 'ja_JP.UTF-8' ENCODING 'UTF8' TEMPLATE template0" > /dev/null
psql -U postgres -d test_ocg_statics -f /docker-entrypoint-initdb.d/sql/02-init.sql > /dev/null
psql -U postgres -d test_ocg_statics -f /docker-entrypoint-initdb.d/sql/03-dml.sql > /dev/null

# SQLファイルの実行
psql -U postgres -d postgres -f /docker-entrypoint-initdb.d/sql/02-init.sql
psql -U postgres -d postgres -f /docker-entrypoint-initdb.d/sql/03-dml.sql