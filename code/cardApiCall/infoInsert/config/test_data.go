package config

import (
	"log"

	"github.com/samber/do"
)

const truncateSQL = `
	TRUNCATE TABLE cards, traps, spells RESTART IDENTITY CASCADE;
`

func BeforeEachForUnitTest() {

	injector := do.New()

	do.Provide(injector, TestDbConnection)

	dbConn := do.MustInvoke[*DbConn](injector)

	// テーブルのトランケート
	_, err := dbConn.Exec(truncateSQL)
	if err != nil {
		log.Fatalf("Failed to truncate tables: %v", err)
	}

	// DBコネクションのクローズ
	dbConn.Close()

}

func AfterEachForUnitTest() {
	injector := do.New()

	do.Provide(injector, TestDbConnection)

	dbConn := do.MustInvoke[*DbConn](injector)

	var err error
	// テーブルのトランケート
	_, err = dbConn.Exec(truncateSQL)

	if err != nil {
		log.Fatalf("Failed to delete test data: %v", err)
	}

	// DBコネションのクローズ
	dbConn.Close()

}