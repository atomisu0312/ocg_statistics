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
func testRaceRetrieval(t *testing.T, dbConn *config.DbConn, target kind.RaceKind, getFunc func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error)) {
	ctx := context.Background()
	tr := transaction.NewTx(dbConn.DB)

	var result kind.RaceKind
	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		raceRepo := repository.NewRaceRepository(q)
		race, err := getFunc(ctx, raceRepo)
		if err != nil {
			return fmt.Errorf("error getting race: %w", err)
		}
		result = race
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, target.ID, result.ID, "IDが一致しません")
	assert.Equal(t, target.NameJa, result.NameJa, "名前（和名）が一致しません")
	assert.Equal(t, target.NameEn, result.NameEn, "名前（英名）が一致しません")
}

// TestForRace tests the RaceRepository
func TestForRace(t *testing.T) {

	t.Run("魔法使い族Enumの取得", func(t *testing.T) {
		target := kind.RaceSpellcaster

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testRaceRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error) {
			return repo.GetRaceByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testRaceRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error) {
			return repo.GetRaceByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testRaceRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error) {
			return repo.GetRaceById(ctx, target.ID)
		})
	})

	t.Run("ドラゴン族Enumの取得", func(t *testing.T) {
		target := kind.RaceDragon

		// セットアップ
		dbConn, _, cleanup := setupTest(t)
		defer cleanup()

		// 和名で取得テスト
		testRaceRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error) {
			return repo.GetRaceByNameJa(ctx, target.NameJa)
		})

		// 英名で取得テスト
		testRaceRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error) {
			return repo.GetRaceByNameEn(ctx, target.NameEn)
		})

		// IDで取得テスト
		testRaceRetrieval(t, dbConn, target, func(ctx context.Context, repo repository.RaceRepository) (kind.RaceKind, error) {
			return repo.GetRaceById(ctx, target.ID)
		})
	})
}
