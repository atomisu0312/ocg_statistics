package repository_test

import (
	"context"
	"fmt"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/dto/carddto"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestForTrap tests the TrapRepository
func TestForTrap(t *testing.T) {

	t.Run("正常系01 トラップカードの新規登録処理", func(t *testing.T) {
		// セットアップ
		dbConn, card, cleanup := setupTest(t)
		defer cleanup()

		// Test data
		trapTypeID := int32(3) // カウンター罠のIDを便宜上セット

		// トランザクションの整備
		ctx := context.Background()
		tr := transaction.NewTx(dbConn.DB)

		// トランザクション境界の中で実行(useCaseではこの中にbaseCard挿入処理を入れる)
		var insertedTrap sqlc_gen.Trap
		err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
			trapRepo := repository.NewTrapRepository(q)

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

		// データの取得
		ctx2 := context.Background()
		var fetchedTrap carddto.TrapCardSelectResult

		err = tr.ExecTx(ctx2, func(q *sqlc_gen.Queries) error {
			trapRepo := repository.NewTrapRepository(q)

			trap, err := trapRepo.GetTrapByCardID(ctx2, insertedTrap.CardID)
			if err != nil {
				return fmt.Errorf("error inserting trap: %w", err)
			}
			fetchedTrap = trap
			return nil
		})

		assert.Equal(t, insertedTrap.CardID, fetchedTrap.ID, "Fetched trap card ID should match the inserted one")
		assert.Equal(t, "カウンター", fetchedTrap.TrapTypeNameJa, "Fetched trap type ID should match the inserted one")
	})
}
