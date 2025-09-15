package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/kind"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// RaceRepository defines the interface for race database operations.
type RaceRepository interface {
	Repository
	GetRaceByNameJa(ctx context.Context, nameJa string) (kind.RaceKind, error)
	GetRaceByNameEn(ctx context.Context, nameEn string) (kind.RaceKind, error)
	GetRaceById(ctx context.Context, id int32) (kind.RaceKind, error)
}

type raceRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewRaceRepository creates a new instance of RaceRepository.
func NewRaceRepository(q *sqlc_gen.Queries) RaceRepository {
	return NewRepository(func(r *repository) RaceRepository {
		return &raceRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetRaceByNameJa は和名で種族を取得する
func (r *raceRepositoryImpl) GetRaceByNameJa(ctx context.Context, nameJa string) (kind.RaceKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetRaceByNameJa", start, zap.String("name_ja", nameJa))

	k := kind.RaceKind{}
	race, err := r.queries.SelectRacesByNameJa(ctx, sql.NullString{String: nameJa, Valid: true})
	if err != nil {
		r.logDBError("GetRaceByNameJa", err, zap.String("name_ja", nameJa))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     race.ID,
		NameJa: race.NameJa,
		NameEn: race.NameEn,
	}

	return kind.RaceKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}

// GetRaceByNameEn は英名で種族を取得する
func (r *raceRepositoryImpl) GetRaceByNameEn(ctx context.Context, nameEn string) (kind.RaceKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetRaceByNameEn", start, zap.String("name_en", nameEn))

	k := kind.RaceKind{}
	race, err := r.queries.SelectRacesByNameEn(ctx, sql.NullString{String: nameEn, Valid: true})
	if err != nil {
		r.logDBError("GetRaceByNameEn", err, zap.String("name_en", nameEn))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     race.ID,
		NameJa: race.NameJa,
		NameEn: race.NameEn,
	}

	return kind.RaceKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}

// GetRaceById はIDで種族を取得する
func (r *raceRepositoryImpl) GetRaceById(ctx context.Context, id int32) (kind.RaceKind, error) {
	start := time.Now()
	defer r.logDBOperation("GetRaceById", start, zap.Int32("id", id))

	k := kind.RaceKind{}
	race, err := r.queries.SelectRacesById(ctx, id)
	if err != nil {
		r.logDBError("GetRaceById", err, zap.Int32("id", id))
		return k, err
	}

	row := kind.SelectFullKindInfoRow{
		ID:     race.ID,
		NameJa: race.NameJa,
		NameEn: race.NameEn,
	}

	return kind.RaceKind{Kind: k.FromSelectFullKindInfoRow(row)}, nil
}
