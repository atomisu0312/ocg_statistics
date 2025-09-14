package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"github.com/samber/do"
)

// UTの際に用いるベースとなるカード
var baseCard = sqlc_gen.InsertCardParams{
	NameJa:     sql.NullString{String: "テストカード", Valid: true},
	NameEn:     sql.NullString{String: "Test Card", Valid: true},
	CardTextJa: sql.NullString{String: "テストカードの説明", Valid: true},
	CardTextEn: sql.NullString{String: "Test Card Description", Valid: true},
	NeuronID:   sql.NullInt64{Int64: 4000, Valid: true},
	OcgApiID:   sql.NullInt64{Int64: 14000, Valid: true},
}

// ベースとなるカードをとりあえず挿入
func insertBaseCard(db *sql.DB) (sqlc_gen.Card, error) {
	ctx := context.Background()
	cardRepo := repository.NewCardRepository(sqlc_gen.New(db))
	card, err := cardRepo.InsertCard(ctx, baseCard)
	if err != nil {
		return sqlc_gen.Card{}, fmt.Errorf("error creating card: %w", err)
	}
	return card, nil
}

// テストの共通セットアップ処理
// 1. DBの疎通確認
// 2. ベースカードを挿入
// 3. クリーンアップ関数の返却
func setupTest(t *testing.T) (*config.DbConn, sqlc_gen.Card, func()) {
	// テスト前処理
	config.BeforeEachForUnitTest()

	// DIコンテナ内の依存関係を設定
	injector := do.New()
	do.Provide(injector, config.TestDbConnection)
	dbConn := do.MustInvoke[*config.DbConn](injector)

	// ベースカードを挿入
	card, err := insertBaseCard(dbConn.DB)

	// エラーが発生した場合はクリーンアップを行う
	if err != nil {
		dbConn.Close()
		config.AfterEachForUnitTest()
		t.Fatalf("Failed to insert base card: %v", err)
	}

	// クリーンアップ関数を返す
	cleanup := func() {
		dbConn.Close()
		config.AfterEachForUnitTest()
	}

	return dbConn, card, cleanup
}
