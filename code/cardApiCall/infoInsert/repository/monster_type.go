package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/kind"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// MonsterTypeRepository defines the interface for monster type database operations.
type MonsterTypeRepository interface {
	Repository
	GetMonsterTypeByNameJa(ctx context.Context, nameJa string) (kind.MonsterKind, error)
	GetMonsterTypeByNameEn(ctx context.Context, nameEn string) (kind.MonsterKind, error)
	GetMonsterTypeById(ctx context.Context, id int32) (kind.MonsterKind, error)
}

type monsterTypeRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewMonsterTypeRepository creates a new instance of MonsterTypeRepository.
func NewMonsterTypeRepository(q *sqlc_gen.Queries) MonsterTypeRepository {
	return NewRepository(func(r *repository) MonsterTypeRepository {
		return &monsterTypeRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetMonsterTypeByNameJa は和名でモンスタータイプを取得する
func (r *monsterTypeRepositoryImpl) GetMonsterTypeByNameJa(ctx context.Context, nameJa string) (kind.MonsterKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterTypeByNameJa", start, zap.String("name_ja", nameJa))

	k := kind.MonsterKind{}
	monsterType, err := r.queries.SelectMonsterTypesByNameJa(ctx, sql.NullString{String: nameJa, Valid: true})
	if err != nil {
		r.logDBError("GetMonsterTypeByNameJa", err, zap.String("name_ja", nameJa))
		return k, err
	}

	// 0件の場合は、空を返す
	if len(monsterType) == 0 {
		return k, nil
	}

	// 複数件は想定外なのでエラーにする（必要なら先頭採用に変更）
	if len(monsterType) > 1 {
		return k, fmt.Errorf("multiple monster_types found (name_ja=%v): count=%d", nameJa, len(monsterType))
	}

	// 1つ目の要素のみ取得
	row := kind.SelectFullKindInfoRow{
		ID:     monsterType[0].ID,
		NameJa: monsterType[0].NameJa,
		NameEn: monsterType[0].NameEn,
	}

	return kind.MonsterKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}

// GetMonsterTypeByNameEn は英名でモンスタータイプを取得する
func (r *monsterTypeRepositoryImpl) GetMonsterTypeByNameEn(ctx context.Context, nameEn string) (kind.MonsterKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterTypeByNameEn", start, zap.String("name_en", nameEn))

	k := kind.MonsterKind{}
	monsterType, err := r.queries.SelectMonsterTypesByNameEn(ctx, sql.NullString{String: nameEn, Valid: true})
	if err != nil {
		r.logDBError("GetMonsterTypeByNameEn", err, zap.String("name_en", nameEn))
		return k, err
	}

	// 0件の場合は、空を返す
	if len(monsterType) == 0 {
		return k, nil
	}

	// 複数件は想定外なのでエラーにする（必要なら先頭採用に変更）
	if len(monsterType) > 1 {
		return k, fmt.Errorf("multiple monster_types found (name_en=%v): count=%d", nameEn, len(monsterType))
	}

	// 1つ目の要素のみ取得
	row := kind.SelectFullKindInfoRow{
		ID:     monsterType[0].ID,
		NameJa: monsterType[0].NameJa,
		NameEn: monsterType[0].NameEn,
	}

	return kind.MonsterKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}

// GetMonsterTypeById はIDでモンスタータイプを取得する
func (r *monsterTypeRepositoryImpl) GetMonsterTypeById(ctx context.Context, id int32) (kind.MonsterKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterTypeById", start, zap.Int32("id", id))

	k := kind.MonsterKind{}
	monsterType, err := r.queries.SelectMonsterTypesById(ctx, id)
	if err != nil {
		r.logDBError("GetMonsterTypeById", err, zap.Int32("id", id))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     monsterType.ID,
		NameJa: monsterType.NameJa,
		NameEn: monsterType.NameEn,
	}

	return kind.MonsterKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}
