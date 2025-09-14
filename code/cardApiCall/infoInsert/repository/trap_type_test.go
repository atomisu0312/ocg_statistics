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
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// テスト用の共通処理
func testTrapTypeRetrieval(t *testing.T, dbConn *config.DbConn, target kind.TrapKind, getFunc func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error)) {
	ctx := context.Background()
	tr := transaction.NewTx(dbConn.DB)

	var result kind.TrapKind
	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		trapTypeRepo := repository.NewTrapTypeRepository(q)
		trapType, err := getFunc(ctx, trapTypeRepo)
		if err != nil {
			return fmt.Errorf("error getting trap type: %w", err)
		}
		result = trapType
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, result.ID, target.ID, "IDが一致しません")
	assert.Equal(t, target.NameJa, result.NameJa, "名前（和名）が一致しません")
	assert.Equal(t, target.NameEn, result.NameEn, "名前（英名）が一致しません")
}

// TestForTrap tests the TrapRepository
func TestForTrapType(t *testing.T) {

	t.Run("カウンター罠Enumの取得", func(t *testing.T) {
		target := kind.TrapTypeCounter

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeById(ctx, target.ID)
		})
	})

	t.Run("永続罠Enumの取得", func(t *testing.T) {
		target := kind.TrapTypeContinuous

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeById(ctx, target.ID)
		})
	})

	t.Run("通常罠Enumの取得", func(t *testing.T) {
		target := kind.TrapTypeNormal

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testTrapTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.TrapTypeRepository) (kind.TrapKind, error) {
			return repo.GetTrapTypeById(ctx, target.ID)
		})
	})
}
