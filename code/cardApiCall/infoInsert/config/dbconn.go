package config

import "database/sql"

type DbConn struct {
	*sql.DB
}
