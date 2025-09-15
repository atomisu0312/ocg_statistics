package repository

import (
	"context"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// XyzMonsterRepository defines the interface for monster card database operations.
type XyzMonsterRepository interface {
	Repository
	GetXyzMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.XyzMonsterSelectResult, error)
	InsertXyzMonster(ctx context.Context, cardId int64) (sqlc_gen.XyzMonster, error)
}

type xyzMonsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewXyzMonsterRepository	creates a new instance of XyzMonsterRepository.
func NewXyzMonsterRepository(q *sqlc_gen.Queries) XyzMonsterRepository {
	return NewRepository(func(r *repository) XyzMonsterRepository {
		return &xyzMonsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetXyzMonsterByCardID retrieves a monster card by its card ID.
func (r *xyzMonsterRepositoryImpl) GetXyzMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.XyzMonsterSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetXyzMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullXyzMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetXyzMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.XyzMonsterSelectResult{}, err
	}

	var result cardrecord.XyzMonsterSelectResult
	result = *result.FromSelectFullXyzMonsterCardInfoRow(cardrecord.SelectFullXyzMonsterCardInfoRow{
		SelectFullMonsterCardInfoRow: cardrecord.SelectFullMonsterCardInfoRow{
			ID:              monster.ID,
			NeuronID:        monster.NeuronID,
			OcgApiID:        monster.OcgApiID,
			NameJa:          monster.NameJa,
			NameEn:          monster.NameEn,
			CardTextJa:      monster.CardTextJa,
			CardTextEn:      monster.CardTextEn,
			Dataowner:       monster.Dataowner,
			RegistDate:      monster.RegistDate,
			EnableStartDate: monster.EnableStartDate,
			EnableEndDate:   monster.EnableEndDate,
			Version:         monster.Version,
			Attack:          monster.Attack,
			Defense:         monster.Defense,
			Level:           monster.Level,
			TypeNamesJa:     monster.TypeNamesJa,
			TypeNamesEn:     monster.TypeNamesEn,
			RaceNameJa:      monster.RaceNameJa,
			RaceNameEn:      monster.RaceNameEn,
			AttributeNameJa: monster.AttributeNameJa,
			AttributeNameEn: monster.AttributeNameEn,
		},
	})
	r.logDBResult("GetXyzMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertXyzMonster inserts a new monster card into the database.
func (r *xyzMonsterRepositoryImpl) InsertXyzMonster(ctx context.Context, cardId int64) (sqlc_gen.XyzMonster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertXyzMonster", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.InsertXyzMonster(ctx, cardId)
	if err != nil {
		r.logDBError("InsertXyzMonster", err, zap.Int64("card_id", cardId))
		return sqlc_gen.XyzMonster{}, err
	}

	r.logDBResult("InsertXyzMonster", monster, zap.Int64("card_id", cardId))
	return monster, nil
}
