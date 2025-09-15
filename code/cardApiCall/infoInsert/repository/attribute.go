package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/kind"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// AttributeRepository defines the interface for attribute database operations.
type AttributeRepository interface {
	Repository
	GetAttributeByNameJa(ctx context.Context, nameJa string) (kind.AttributeKind, error)
	GetAttributeByNameEn(ctx context.Context, nameEn string) (kind.AttributeKind, error)
	GetAttributeById(ctx context.Context, id int32) (kind.AttributeKind, error)
}

type attributeRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewAttributeRepository creates a new instance of AttributeRepository.
func NewAttributeRepository(q *sqlc_gen.Queries) AttributeRepository {
	return NewRepository(func(r *repository) AttributeRepository {
		return &attributeRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetAttributeByNameJa は和名で属性を取得する
func (r *attributeRepositoryImpl) GetAttributeByNameJa(ctx context.Context, nameJa string) (kind.AttributeKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetAttributeByNameJa", start, zap.String("name_ja", nameJa))

	k := kind.AttributeKind{}
	attribute, err := r.queries.SelectAttributesByNameJa(ctx, sql.NullString{String: nameJa, Valid: true})
	if err != nil {
		r.logDBError("GetAttributeByNameJa", err, zap.String("name_ja", nameJa))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     attribute.ID,
		NameJa: attribute.NameJa,
		NameEn: attribute.NameEn,
	}

	return kind.AttributeKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}

// GetAttributeByNameEn は英名で属性を取得する
func (r *attributeRepositoryImpl) GetAttributeByNameEn(ctx context.Context, nameEn string) (kind.AttributeKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetAttributeByNameEn", start, zap.String("name_en", nameEn))

	k := kind.AttributeKind{}
	attribute, err := r.queries.SelectAttributesByNameEn(ctx, sql.NullString{String: nameEn, Valid: true})
	if err != nil {
		r.logDBError("GetAttributeByNameEn", err, zap.String("name_en", nameEn))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     attribute.ID,
		NameJa: attribute.NameJa,
		NameEn: attribute.NameEn,
	}

	return kind.AttributeKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}

// GetAttributeById はIDで属性を取得する
func (r *attributeRepositoryImpl) GetAttributeById(ctx context.Context, id int32) (kind.AttributeKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetAttributeById", start, zap.Int32("id", id))

	k := kind.AttributeKind{}
	attribute, err := r.queries.SelectAttributesById(ctx, id)
	if err != nil {
		r.logDBError("GetAttributeById", err, zap.Int32("id", id))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     attribute.ID,
		NameJa: attribute.NameJa,
		NameEn: attribute.NameEn,
	}

	return kind.AttributeKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}
