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
func testMonsterTypeRetrieval(t *testing.T, dbConn *config.DbConn, target kind.MonsterKind, getFunc func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error)) {
	ctx := context.Background()
	tr := transaction.NewTx(dbConn.DB)

	var result kind.MonsterKind
	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		monsterTypeRepo := repository.NewMonsterTypeRepository(q)
		monsterType, err := getFunc(ctx, monsterTypeRepo)
		if err != nil {
			return fmt.Errorf("error getting monster type: %w", err)
		}
		result = monsterType
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, target.ID, result.ID, "IDが一致しません")
	assert.Equal(t, target.NameJa, result.NameJa, "名前（和名）が一致しません")
	assert.Equal(t, target.NameEn, result.NameEn, "名前（英名）が一致しません")
}

// TestForMonsterType tests the MonsterTypeRepository
func TestForMonsterType(t *testing.T) {

	t.Run("通常モンスターEnumの取得", func(t *testing.T) {
		target := kind.MonsterTypeNormal

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testMonsterTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error) {
			return repo.GetMonsterTypeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testMonsterTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error) {
			return repo.GetMonsterTypeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testMonsterTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error) {
			return repo.GetMonsterTypeById(ctx, target.ID)
		})
	})

	t.Run("効果モンスターEnumの取得", func(t *testing.T) {
		target := kind.MonsterTypeEffect

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testMonsterTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error) {
			return repo.GetMonsterTypeByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testMonsterTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error) {
			return repo.GetMonsterTypeByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testMonsterTypeRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.MonsterTypeRepository) (kind.MonsterKind, error) {
			return repo.GetMonsterTypeById(ctx, target.ID)
		})
	})
}
