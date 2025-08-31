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
