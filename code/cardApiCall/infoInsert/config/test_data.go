package config

import (
	"log"

	"github.com/samber/do"
)

const truncateSQL = `
	TRUNCATE TABLE cards RESTART IDENTITY CASCADE;
	TRUNCATE TABLE traps RESTART IDENTITY CASCADE;
`

func BeforeEachForUnitTest() {

	injector := do.New()

	do.Provide(injector, TestDbConnection)

	dbConn := do.MustInvoke[*DbConn](injector)

	// Insert initial data
	_, err := dbConn.Exec(truncateSQL)
	if err != nil {
		log.Fatalf("Failed to insert initial data: %v", err)
	}

	// Clean up
	dbConn.Close()

}

func AfterEachForUnitTest() {
	injector := do.New()

	do.Provide(injector, TestDbConnection)

	dbConn := do.MustInvoke[*DbConn](injector)

	var err error
	// Clean up: Delete the inserted data
	_, err = dbConn.Exec(truncateSQL)

	if err != nil {
		log.Fatalf("Failed to delete test data: %v", err)
	}

	// Clean up
	dbConn.Close()

}
