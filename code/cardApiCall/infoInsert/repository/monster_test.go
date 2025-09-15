package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestMonsterInsert(t *testing.T) {
	t.Run("正常系01: モンスターの挿入", func(t *testing.T) {
		// セットアップ
		dbConn, card, cleanup := setupTest(t)
		defer cleanup()

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repo := repository.NewMonsterRepository(q)

		testCardID := card.ID       // Unique card ID for testing
		testRaceID := int32(1)      // Example race ID
		testAttributeID := int32(1) // Example attribute ID
		testAttack := int32(1000)
		testDefense := int32(500)
		testLevel := int32(4)
		testTypeIDs := []int32{1, 2} // Example type IDs

		monster, err := repo.InsertMonster(ctx, testCardID, testRaceID, testAttributeID, testAttack, testDefense, testLevel, testTypeIDs)
		assert.NoError(t, err)
		assert.NotEqual(t, sqlc_gen.Monster{}, monster)
		assert.Equal(t, testCardID, monster.CardID)
		assert.Equal(t, testRaceID, monster.RaceID.Int32)
		assert.Equal(t, testAttributeID, monster.AttributeID.Int32)
		assert.Equal(t, testAttack, monster.Attack.Int32)
		assert.Equal(t, testDefense, monster.Defense.Int32)
		assert.Equal(t, testLevel, monster.Level.Int32)
		assert.Equal(t, testTypeIDs, monster.TypeIds)
		assert.True(t, monster.RegistDate.Valid)
		assert.True(t, monster.EnableStartDate.Valid)
		assert.True(t, monster.Version.Valid)
		assert.Equal(t, int64(1), monster.Version.Int64)
	})
}

func TestGetMonsterByCardID(t *testing.T) {
	t.Run("正常系01: 既存のモンスターを取得", func(t *testing.T) {
		dbConn, card, cleanup := setupTest(t)
		defer cleanup()

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repo := repository.NewMonsterRepository(q)

		testCardID := card.ID // Unique card ID for testing
		testRaceID := int32(2)
		testAttributeID := int32(2)
		testAttack := int32(2000)
		testDefense := int32(1000)
		testLevel := int32(6)
		testTypeIDs := []int32{3, 4}

		// Insert a monster first
		insertedMonster, err := repo.InsertMonster(ctx, testCardID, testRaceID, testAttributeID, testAttack, testDefense, testLevel, testTypeIDs)
		assert.NoError(t, err)
		assert.NotEqual(t, sqlc_gen.Monster{}, insertedMonster)
		fmt.Println(insertedMonster)

		// Now retrieve it
		retrievedMonsterResult, err := repo.GetMonsterByCardID(ctx, testCardID)
		assert.NoError(t, err)
		assert.NotEqual(t, cardrecord.MonsterCardSelectResult{}, retrievedMonsterResult)
		assert.Equal(t, testCardID, retrievedMonsterResult.ID)
		assert.Equal(t, "ドラゴン族", retrievedMonsterResult.RaceNameJa)
		assert.Equal(t, "Dragon", retrievedMonsterResult.RaceNameEn)
		assert.Equal(t, "闇", retrievedMonsterResult.AttributeNameJa)
		assert.Equal(t, "DARK", retrievedMonsterResult.AttributeNameEn)
		assert.Equal(t, insertedMonster.Attack.Int32, retrievedMonsterResult.Attack)
		assert.Equal(t, insertedMonster.Defense.Int32, retrievedMonsterResult.Defense)
		assert.Equal(t, insertedMonster.Level.Int32, retrievedMonsterResult.Level)
		assert.Equal(t, []string{"トゥーン", "スピリット"}, retrievedMonsterResult.TypeNamesJa)
		assert.Equal(t, []string{"Toon", "Spirit"}, retrievedMonsterResult.TypeNamesEn)
	})

	t.Run("異常系01: 存在しないモンスターを取得", func(t *testing.T) {
		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := do.New()
		do.Provide(injector, config.TestDbConnection)
		dbConn := do.MustInvoke[*config.DbConn](injector)

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repo := repository.NewMonsterRepository(q)

		nonExistentCardID := int64(99999998) // A card ID that should not exist

		_, err := repo.GetMonsterByCardID(ctx, nonExistentCardID)
		assert.Error(t, err)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}
