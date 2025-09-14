package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/kind"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// TrapTypeRepository defines the interface for trap type database operations.
type TrapTypeRepository interface {
	Repository
	GetTrapTypeByNameJa(ctx context.Context, nameJa string) (kind.TrapKind, error)
	GetTrapTypeByNameEn(ctx context.Context, nameEn string) (kind.TrapKind, error)
	GetTrapTypeById(ctx context.Context, id int32) (kind.TrapKind, error)
}

type trapTypeRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewTrapTypeRepository creates a new instance of TrapTypeRepository.
func NewTrapTypeRepository(q *sqlc_gen.Queries) TrapTypeRepository {
	return NewRepository(func(r *repository) TrapTypeRepository {
		return &trapTypeRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetTrapTypeByNameJa は和名で罠タイプを取得する
func (r *trapTypeRepositoryImpl) GetTrapTypeByNameJa(ctx context.Context, nameJa string) (kind.TrapKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetTrapTypeByNameJa", start, zap.String("name_ja", nameJa))

	k := kind.TrapKind{}
	trapType, err := r.queries.SelectTrapTypesByNameJa(ctx, sql.NullString{String: nameJa, Valid: true})
	if err != nil {
		r.logDBError("GetTrapTypeByNameJa", err, zap.String("name_ja", nameJa))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     trapType.ID,
		NameJa: trapType.NameJa,
		NameEn: trapType.NameEn,
	}

	return k.FromSelectFullKindInfoRow(row), nil
}

// GetTrapTypeByNameEn は英名で罠タイプを取得する
func (r *trapTypeRepositoryImpl) GetTrapTypeByNameEn(ctx context.Context, nameEn string) (kind.TrapKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetTrapTypeByNameEn", start, zap.String("name_en", nameEn))

	k := kind.TrapKind{}
	trapType, err := r.queries.SelectTrapTypesByNameEn(ctx, sql.NullString{String: nameEn, Valid: true})
	if err != nil {
		r.logDBError("GetTrapTypeByNameEn", err, zap.String("name_en", nameEn))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     trapType.ID,
		NameJa: trapType.NameJa,
		NameEn: trapType.NameEn,
	}

	return k.FromSelectFullKindInfoRow(row), nil
}

// GetTrapTypeById はIDで罠タイプを取得する
func (r *trapTypeRepositoryImpl) GetTrapTypeById(ctx context.Context, id int32) (kind.TrapKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetTrapTypeById", start, zap.Int32("id", id))

	k := kind.TrapKind{}
	trapType, err := r.queries.SelectTrapTypesById(ctx, id)
	if err != nil {
		r.logDBError("GetTrapTypeById", err, zap.Int32("id", id))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     trapType.ID,
		NameJa: trapType.NameJa,
		NameEn: trapType.NameEn,
	}

	return k.FromSelectFullKindInfoRow(row), nil
}
