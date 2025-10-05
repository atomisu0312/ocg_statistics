package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

// 正常系テスト
func TestForCardInit(t *testing.T) {
	t.Run("正常系01 ID:1", func(t *testing.T) {
		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理

		// DIコンテナ内の依存関係を設定
		injector := do.New()
		do.Provide(injector, config.TestDbConnection)
		dbConn := do.MustInvoke[*config.DbConn](injector)

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repo := repository.NewCardRepository(q)
		card, err := repo.GetCardByID(ctx, 1)
		assert.Equal(t, sqlc_gen.Card{}, card)
		assert.Error(t, err)
	})
}

func TestForCardInsert(t *testing.T) {
	t.Run("正常系01 ID:1", func(t *testing.T) {
		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理

		inserData := sqlc_gen.InsertCardParams{
			NameEn:     sql.NullString{String: "Test Card", Valid: true},
			NameJa:     sql.NullString{String: "テストカード", Valid: true},
			CardTextJa: sql.NullString{String: "テストカードの説明", Valid: true},
			CardTextEn: sql.NullString{String: "Test Card Description", Valid: true},
			NeuronID:   sql.NullInt64{Int64: 1, Valid: true},
			OcgApiID:   sql.NullInt64{Int64: 1, Valid: true},
		}

		injector := do.New()
		do.Provide(injector, config.TestDbConnection)
		dbConn := do.MustInvoke[*config.DbConn](injector)

		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repo := repository.NewCardRepository(q)
		card, err := repo.InsertCard(ctx, inserData)
		assert.NoError(t, err)
		assert.NotEqual(t, sqlc_gen.Card{}, card)
		assert.Equal(t, inserData.NameJa.String, card.NameJa.String)
		assert.Equal(t, inserData.NameEn.String, card.NameEn.String)
		assert.Equal(t, inserData.CardTextJa.String, card.CardTextJa.String)
		assert.Equal(t, inserData.CardTextEn.String, card.CardTextEn.String)
		assert.Equal(t, inserData.NeuronID.Int64, card.NeuronID.Int64)
		assert.Equal(t, inserData.OcgApiID.Int64, card.OcgApiID.Int64)
	})
}

func TestForCardPatternByCardID(t *testing.T) {
	t.Run("正常系01 モンスターカード", func(t *testing.T) {
		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理
		injector := do.New()
		do.Provide(injector, config.TestDbConnection)

		insertCardData := sqlc_gen.InsertCardParams{
			NameEn:     sql.NullString{String: "Test Card", Valid: true},
			NameJa:     sql.NullString{String: "テストカード", Valid: true},
			CardTextJa: sql.NullString{String: "テストカードの説明", Valid: true},
			CardTextEn: sql.NullString{String: "Test Card Description", Valid: true},
			NeuronID:   sql.NullInt64{Int64: 1, Valid: true},
			OcgApiID:   sql.NullInt64{Int64: 1, Valid: true},
		}

		dbConn := do.MustInvoke[*config.DbConn](injector)
		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repoCard := repository.NewCardRepository(q)
		card, err := repoCard.InsertCard(ctx, insertCardData)
		assert.NoError(t, err)

		repoMonster := repository.NewMonsterRepository(q)
		testCardID := card.ID
		testRaceID := int32(1)
		testAttributeID := int32(1)
		testAttack := int32(1000)
		testDefense := int32(500)
		testLevel := int32(4)
		testTypeIDs := []int32{1, 2}

		monster, err := repoMonster.InsertMonster(ctx, testCardID, testRaceID, testAttributeID, testAttack, testDefense, testLevel, testTypeIDs)
		assert.NoError(t, err)

		cardPattern, err := repoCard.GetCardPatternByCardID(ctx, monster.CardID)
		assert.NoError(t, err)
		assert.Equal(t, testCardID, cardPattern.CardID)
		assert.Equal(t, insertCardData.NeuronID.Int64, cardPattern.NeuronID.Int64)
		assert.Equal(t, insertCardData.OcgApiID.Int64, cardPattern.OcgApiID.Int64)
		assert.True(t, cardPattern.IsMonster)
		assert.False(t, cardPattern.IsSpell)
		assert.False(t, cardPattern.IsTrap)
	})

	t.Run("正常系02 魔法カード", func(t *testing.T) {
		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理
		injector := do.New()
		do.Provide(injector, config.TestDbConnection)

		insertCardData := sqlc_gen.InsertCardParams{
			NameEn:     sql.NullString{String: "Test Spell Card", Valid: true},
			NameJa:     sql.NullString{String: "テスト魔法カード", Valid: true},
			CardTextJa: sql.NullString{String: "テスト魔法カードの説明", Valid: true},
			CardTextEn: sql.NullString{String: "Test Spell Card Description", Valid: true},
			NeuronID:   sql.NullInt64{Int64: 2, Valid: true},
			OcgApiID:   sql.NullInt64{Int64: 2, Valid: true},
		}

		dbConn := do.MustInvoke[*config.DbConn](injector)
		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repoCard := repository.NewCardRepository(q)
		card, err := repoCard.InsertCard(ctx, insertCardData)
		assert.NoError(t, err)

		repoSpell := repository.NewSpellRepository(q)
		testCardID := card.ID
		testSpellTypeID := int32(1)

		spell, err := repoSpell.InsertSpell(ctx, testCardID, testSpellTypeID)
		assert.NoError(t, err)

		cardPattern, err := repoCard.GetCardPatternByCardID(ctx, spell.CardID)
		assert.NoError(t, err)
		assert.Equal(t, testCardID, cardPattern.CardID)
		assert.Equal(t, insertCardData.NeuronID.Int64, cardPattern.NeuronID.Int64)
		assert.Equal(t, insertCardData.OcgApiID.Int64, cardPattern.OcgApiID.Int64)
		assert.False(t, cardPattern.IsMonster)
		assert.True(t, cardPattern.IsSpell)
		assert.False(t, cardPattern.IsTrap)
	})

	t.Run("正常系03 罠カード", func(t *testing.T) {
		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理
		injector := do.New()
		do.Provide(injector, config.TestDbConnection)

		insertCardData := sqlc_gen.InsertCardParams{
			NameEn:     sql.NullString{String: "Test Trap Card", Valid: true},
			NameJa:     sql.NullString{String: "テスト罠カード", Valid: true},
			CardTextJa: sql.NullString{String: "テスト罠カードの説明", Valid: true},
			CardTextEn: sql.NullString{String: "Test Trap Card Description", Valid: true},
			NeuronID:   sql.NullInt64{Int64: 3, Valid: true},
			OcgApiID:   sql.NullInt64{Int64: 3, Valid: true},
		}

		dbConn := do.MustInvoke[*config.DbConn](injector)
		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repoCard := repository.NewCardRepository(q)
		card, err := repoCard.InsertCard(ctx, insertCardData)
		assert.NoError(t, err)

		repoTrap := repository.NewTrapRepository(q)
		testCardID := card.ID
		testTrapTypeID := int32(1)

		trap, err := repoTrap.InsertTrap(ctx, testCardID, testTrapTypeID)
		assert.NoError(t, err)

		cardPattern, err := repoCard.GetCardPatternByCardID(ctx, trap.CardID)
		assert.NoError(t, err)
		assert.Equal(t, testCardID, cardPattern.CardID)
		assert.Equal(t, insertCardData.NeuronID.Int64, cardPattern.NeuronID.Int64)
		assert.Equal(t, insertCardData.OcgApiID.Int64, cardPattern.OcgApiID.Int64)
		assert.False(t, cardPattern.IsMonster)
		assert.False(t, cardPattern.IsSpell)
		assert.True(t, cardPattern.IsTrap)
	})

	t.Run("正常系04 存在しないカードID", func(t *testing.T) {
		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理
		injector := do.New()
		do.Provide(injector, config.TestDbConnection)

		dbConn := do.MustInvoke[*config.DbConn](injector)
		ctx := context.Background()
		q := sqlc_gen.New(dbConn)
		repoCard := repository.NewCardRepository(q)

		testID := int64(999)
		cardPattern, err := repoCard.GetCardPatternByCardID(ctx, testID)
		assert.NoError(t, err)
		assert.Equal(t, testID, cardPattern.CardID)
		assert.False(t, cardPattern.NeuronID.Valid)
		assert.False(t, cardPattern.OcgApiID.Valid)
		assert.False(t, cardPattern.IsMonster)
		assert.False(t, cardPattern.IsSpell)
		assert.False(t, cardPattern.IsTrap)
	})
}
