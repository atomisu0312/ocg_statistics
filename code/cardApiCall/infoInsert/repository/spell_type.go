package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/kind"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// SpellTypeRepository defines the interface for spell type database operations.
type SpellTypeRepository interface {
	Repository
	GetSpellTypeByNameJa(ctx context.Context, nameJa string) (kind.SpellKind, error)
	GetSpellTypeByNameEn(ctx context.Context, nameEn string) (kind.SpellKind, error)
	GetSpellTypeById(ctx context.Context, id int32) (kind.SpellKind, error)
}

type spellTypeRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewSpellTypeRepository creates a new instance of SpellTypeRepository.
func NewSpellTypeRepository(q *sqlc_gen.Queries) SpellTypeRepository {
	return NewRepository(func(r *repository) SpellTypeRepository {
		return &spellTypeRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetSpellTypeByNameJa は和名で魔法タイプを取得する
func (r *spellTypeRepositoryImpl) GetSpellTypeByNameJa(ctx context.Context, nameJa string) (kind.SpellKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetSpellTypeByNameJa", start, zap.String("name_ja", nameJa))

	k := kind.SpellKind{}
	spellType, err := r.queries.SelectSpellTypesByNameJa(ctx, sql.NullString{String: nameJa, Valid: true})
	if err != nil {
		r.logDBError("GetSpellTypeByNameJa", err, zap.String("name_ja", nameJa))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     spellType.ID,
		NameJa: spellType.NameJa,
		NameEn: spellType.NameEn,
	}

	return k.FromSelectFullKindInfoRow(row), nil
}

// GetSpellTypeByNameEn は英名で魔法タイプを取得する
func (r *spellTypeRepositoryImpl) GetSpellTypeByNameEn(ctx context.Context, nameEn string) (kind.SpellKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetSpellTypeByNameEn", start, zap.String("name_en", nameEn))

	k := kind.SpellKind{}
	spellType, err := r.queries.SelectSpellTypesByNameEn(ctx, sql.NullString{String: nameEn, Valid: true})
	if err != nil {
		r.logDBError("GetSpellTypeByNameEn", err, zap.String("name_en", nameEn))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     spellType.ID,
		NameJa: spellType.NameJa,
		NameEn: spellType.NameEn,
	}

	return k.FromSelectFullKindInfoRow(row), nil
}

// GetSpellTypeById はIDで魔法タイプを取得する
func (r *spellTypeRepositoryImpl) GetSpellTypeById(ctx context.Context, id int32) (kind.SpellKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetSpellTypeById", start, zap.Int32("id", id))

	k := kind.SpellKind{}
	spellType, err := r.queries.SelectSpellTypesById(ctx, id)
	if err != nil {
		r.logDBError("GetSpellTypeById", err, zap.Int32("id", id))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     spellType.ID,
		NameJa: spellType.NameJa,
		NameEn: spellType.NameEn,
	}

	return k.FromSelectFullKindInfoRow(row), nil
}
