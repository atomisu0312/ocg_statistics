package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestForTrap tests the TrapRepository
func TestForTrap(t *testing.T) {

	t.Run("正常系01 トラップカードの新規登録処理", func(t *testing.T) {
		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := do.New()
		do.Provide(injector, config.TestDbConnection)
		dbConn := do.MustInvoke[*config.DbConn](injector)
		defer dbConn.DB.Close()

		// Test data
		trapTypeID := int32(2) // Assuming trap type 2 exists

		inserData := sqlc_gen.InsertCardParams{
			NameJa:     sql.NullString{String: "テストカード", Valid: true},
			NameEn:     sql.NullString{String: "Test Card", Valid: true},
			CardTextJa: sql.NullString{String: "テストカードの説明", Valid: true},
			CardTextEn: sql.NullString{String: "Test Card Description", Valid: true},
			NeuronID:   sql.NullInt64{Int64: 1, Valid: true},
			OcgApiID:   sql.NullInt64{Int64: 1, Valid: true},
		}

		// Transactional insertion
		ctx := context.Background()
		tr := transaction.NewTx(dbConn.DB)

		var insertedTrap sqlc_gen.Trap
		err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
			cardRepo := repository.NewCardRepository(q)
			trapRepo := repository.NewTrapRepository(q)

			// First, create a card to associate the trap with
			card, err := cardRepo.InsertCard(ctx, inserData)
			if err != nil {
				return fmt.Errorf("error creating card: %w", err)
			}

			trap, err := trapRepo.InsertTrap(ctx, card.ID, trapTypeID)
			if err != nil {
				return fmt.Errorf("error inserting trap: %w", err)
			}
			insertedTrap = trap
			return nil
		})

		require.NoError(t, err, "Transaction should execute without error")

		// Verification
		assert.NotZero(t, insertedTrap.CardID, "Inserted trap should have a non-zero card ID")
		assert.Equal(t, trapTypeID, insertedTrap.TrapTypeID.Int32, "The trap's type ID should match the input")

		// Verify that the data is actually in the database
		q := sqlc_gen.New(dbConn.DB)
		repo := repository.NewTrapRepository(q)
		fetchedTrap, err := repo.GetTrapByCardID(ctx, insertedTrap.CardID)
		require.NoError(t, err, "Should be able to fetch the newly inserted trap")
		assert.Equal(t, insertedTrap.CardID, fetchedTrap.CardID, "Fetched trap card ID should match the inserted one")
		assert.Equal(t, trapTypeID, fetchedTrap.TrapTypeID.Int32, "Fetched trap type ID should match the inserted one")
	})
}
