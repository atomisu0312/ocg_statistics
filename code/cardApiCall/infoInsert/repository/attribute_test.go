package repository_test

import (
	"context"
	"fmt"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/dto/kind"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
	"github.com/stretchr/testify/assert"
)

// テスト用の共通処理
func testAttributeRetrieval(t *testing.T, dbConn *config.DbConn, target kind.AttributeKind, getFunc func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error)) {
	ctx := context.Background()
	tr := transaction.NewTx(dbConn.DB)

	var result kind.AttributeKind
	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		attributeRepo := repository.NewAttributeRepository(q)
		attribute, err := getFunc(ctx, attributeRepo)
		if err != nil {
			return fmt.Errorf("error getting attribute: %w", err)
		}
		result = attribute
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, target.ID, result.ID, "IDが一致しません")
	assert.Equal(t, target.NameJa, result.NameJa, "名前（和名）が一致しません")
	assert.Equal(t, target.NameEn, result.NameEn, "名前（英名）が一致しません")
}

// TestForAttribute tests the AttributeRepository
func TestForAttribute(t *testing.T) {

	t.Run("光属性Enumの取得", func(t *testing.T) {
		target := kind.AttributeLight

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testAttributeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error) {
			return repo.GetAttributeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testAttributeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error) {
			return repo.GetAttributeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testAttributeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error) {
			return repo.GetAttributeById(ctx, target.ID)
		})
	})

	t.Run("闇属性Enumの取得", func(t *testing.T) {
		target := kind.AttributeDark

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testAttributeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error) {
			return repo.GetAttributeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testAttributeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error) {
			return repo.GetAttributeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testAttributeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.AttributeRepository) (kind.AttributeKind, error) {
			return repo.GetAttributeById(ctx, target.ID)
		})
	})
}
