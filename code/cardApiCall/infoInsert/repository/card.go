package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

type CardRepository interface {
	Repository
	GetCardByID(ctx context.Context, cardId int64) (sqlc_gen.Card, error)
	InsertCard(ctx context.Context, arg sqlc_gen.InsertCardParams) (sqlc_gen.Card, error)
	GetCardByNameEn(ctx context.Context, nameEn string) (sqlc_gen.Card, error)
	GetCardByNameJa(ctx context.Context, nameJa string) (sqlc_gen.Card, error)
	GetCardPatternByCardID(ctx context.Context, cardId int64) (sqlc_gen.SelectCardPatternByCardIDRow, error)
}

type cardRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewCardRepository creates a new instance of CardRepository.
func NewCardRepository(q *sqlc_gen.Queries) CardRepository {
	return NewRepository(func(r *repository) CardRepository {
		return &cardRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

func (r *cardRepositoryImpl) GetCardByID(ctx context.Context, cardId int64) (sqlc_gen.Card, error) {
	start := time.Now()
	defer r.logDBOperation("GetCardByID", start, zap.Int64("card_id", cardId))

	card, err := r.queries.GetCard(ctx, cardId)
	if err != nil {
		r.logDBError("GetCardByID", err, zap.Int64("card_id", cardId))
		return sqlc_gen.Card{}, err
	}

	r.logDBResult("GetCardByID", card, zap.Int64("card_id", cardId))
	return card, nil
}

func (r *cardRepositoryImpl) InsertCard(ctx context.Context, arg sqlc_gen.InsertCardParams) (sqlc_gen.Card, error) {
	start := time.Now()
	defer r.logDBOperation("InsertCard", start, zap.Int64("card_id", arg.OcgApiID.Int64))

	card, err := r.queries.InsertCard(ctx, arg)
	if err != nil {
		r.logDBError("InsertCard", err, zap.Int64("card_id", arg.OcgApiID.Int64))
		return sqlc_gen.Card{}, err
	}

	r.logDBResult("InsertCard", card, zap.Int64("card_id", arg.OcgApiID.Int64))
	return card, nil
}

func (r *cardRepositoryImpl) GetCardByNameEn(ctx context.Context, nameEn string) (sqlc_gen.Card, error) {
	start := time.Now()
	defer r.logDBOperation("GetCardByNameEn", start, zap.String("name_en", nameEn))

	card, err := r.queries.SelectByCardNameEn(ctx, sql.NullString{String: nameEn, Valid: true})
	if err != nil {
		r.logDBError("GetCardByNameEn", err, zap.String("name_en", nameEn))
		return sqlc_gen.Card{}, err
	}

	r.logDBResult("GetCardByNameEn", card, zap.String("name_en", nameEn))
	return card, nil
}

func (r *cardRepositoryImpl) GetCardByNameJa(ctx context.Context, nameJa string) (sqlc_gen.Card, error) {
	start := time.Now()
	defer r.logDBOperation("GetCardByNameJa", start, zap.String("name_ja", nameJa))

	card, err := r.queries.SelectByCardNameJa(ctx, sql.NullString{String: nameJa, Valid: true})
	if err != nil {
		r.logDBError("GetCardByNameJa", err, zap.String("name_ja", nameJa))
		return sqlc_gen.Card{}, err
	}

	r.logDBResult("GetCardByNameJa", card, zap.String("name_ja", nameJa))
	return card, nil
}

func (r *cardRepositoryImpl) GetCardPatternByCardID(ctx context.Context, cardId int64) (sqlc_gen.SelectCardPatternByCardIDRow, error) {
	start := time.Now()
	defer r.logDBOperation("GetCardPatternByCardID", start, zap.Int64("card_id", cardId))

	cardPattern, err := r.queries.SelectCardPatternByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetCardPatternByCardID", err, zap.Int64("card_id", cardId))
		return sqlc_gen.SelectCardPatternByCardIDRow{}, err
	}

	r.logDBResult("GetCardPatternByCardID", cardPattern, zap.Int64("card_id", cardId))
	return cardPattern, nil
}
