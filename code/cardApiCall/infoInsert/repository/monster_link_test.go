package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
	"github.com/stretchr/testify/assert"
)

func TestLinkMonsterInsert(t *testing.T) {
	t.Run("正常系01: モンスターの挿入", func(t *testing.T) {
		// セットアップ
		dbConn, card, cleanup := setupTest(t)
		defer cleanup()

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repoMonster := repository.NewMonsterRepository(q)
		repoLink := repository.NewLinkMonsterRepository(q)

		testCardID := card.ID       // Unique card ID for testing
		testRaceID := int32(1)      // Example race ID
		testAttributeID := int32(1) // Example attribute ID
		testAttack := int32(1000)
		testDefense := int32(500)
		testLevel := int32(4)
		testTypeIDs := []int32{1, 2} // Example type IDs
		testLinkMarker := int32(4)

		tr := transaction.NewTx(dbConn.DB)
		var insertedMonster sqlc_gen.LinkMonster
		err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
			monster, err := repoMonster.InsertMonster(ctx, testCardID, testRaceID, testAttributeID, testAttack, testDefense, testLevel, testTypeIDs)
			if err != nil {
				return fmt.Errorf("error inserting monster: %w", err)
			}
			target, err := repoLink.InsertLinkMonster(ctx, monster.CardID, testLinkMarker)
			if err != nil {
				return fmt.Errorf("error inserting link monster: %w", err)
			}
			insertedMonster = target
			return nil
		})

		assert.NoError(t, err)
		assert.NotEqual(t, sqlc_gen.LinkMonster{}, insertedMonster)
		assert.Equal(t, testCardID, insertedMonster.CardID)
		assert.Equal(t, testLinkMarker, insertedMonster.LinkMarker.Int32)
	})
}

func TestGetLinkMonsterByCardID(t *testing.T) {
	t.Run("正常系01: 既存のモンスターを取得", func(t *testing.T) {
		dbConn, card, cleanup := setupTest(t)
		defer cleanup()

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repoMonster := repository.NewMonsterRepository(q)
		repoLink := repository.NewLinkMonsterRepository(q)

		testCardID := card.ID // Unique card ID for testing
		testRaceID := int32(2)
		testAttributeID := int32(2)
		testAttack := int32(2000)
		testDefense := int32(1000)
		testLevel := int32(6)
		testTypeIDs := []int32{2, 4}
		testLinkMarker := int32(4)

		tr := transaction.NewTx(dbConn.DB)
		err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
			monster, err := repoMonster.InsertMonster(ctx, testCardID, testRaceID, testAttributeID, testAttack, testDefense, testLevel, testTypeIDs)
			if err != nil {
				return fmt.Errorf("error inserting monster: %w", err)
			}
			_, err = repoLink.InsertLinkMonster(ctx, monster.CardID, testLinkMarker)
			if err != nil {
				return fmt.Errorf("error inserting link monster: %w", err)
			}
			return nil
		})

		// Now retrieve it
		retrievedMonsterResult, err := repoLink.GetLinkMonsterByCardID(ctx, testCardID)
		assert.NoError(t, err)
		assert.NotEqual(t, cardrecord.LinkMonsterSelectResult{}, retrievedMonsterResult)
		assert.Equal(t, testCardID, retrievedMonsterResult.ID)
		assert.Equal(t, "ドラゴン族", retrievedMonsterResult.RaceNameJa)
		assert.Equal(t, "Dragon", retrievedMonsterResult.RaceNameEn)
		assert.Equal(t, "闇", retrievedMonsterResult.AttributeNameJa)
		assert.Equal(t, "DARK", retrievedMonsterResult.AttributeNameEn)
		assert.Equal(t, testAttack, retrievedMonsterResult.Attack)
		assert.Equal(t, testDefense, retrievedMonsterResult.Defense)
		assert.Equal(t, testLevel, retrievedMonsterResult.Level)
		assert.Equal(t, []string{"効果", "スピリット"}, retrievedMonsterResult.TypeNamesJa)
		assert.Equal(t, []string{"Effect", "Spirit"}, retrievedMonsterResult.TypeNamesEn)
		assert.Equal(t, testLinkMarker, retrievedMonsterResult.LinkMarker.Int32)
	})

	t.Run("異常系01: 存在しないモンスターを取得", func(t *testing.T) {
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repo := repository.NewLinkMonsterRepository(q)

		nonExistentCardID := int64(99999998) // A card ID that should not exist

		_, err := repo.GetLinkMonsterByCardID(ctx, nonExistentCardID)
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}
