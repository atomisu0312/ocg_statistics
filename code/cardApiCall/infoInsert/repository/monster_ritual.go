package repository

import (
	"context"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// RitualMonsterRepository defines the interface for monster card database operations.
type RitualMonsterRepository interface {
	Repository
	GetRitualMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.RitualMonsterSelectResult, error)
	InsertRitualMonster(ctx context.Context, cardId int64) (sqlc_gen.RitualMonster, error)
}

type ritualMonsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewRitualMonsterRepository	creates a new instance of RitualMonsterRepository.
func NewRitualMonsterRepository(q *sqlc_gen.Queries) RitualMonsterRepository {
	return NewRepository(func(r *repository) RitualMonsterRepository {
		return &ritualMonsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetRitualMonsterByCardID retrieves a monster card by its card ID.
func (r *ritualMonsterRepositoryImpl) GetRitualMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.RitualMonsterSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetRitualMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullRitualMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetRitualMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.RitualMonsterSelectResult{}, err
	}

	var result cardrecord.RitualMonsterSelectResult
	result = *result.FromSelectFullRitualMonsterCardInfoRow(cardrecord.SelectFullRitualMonsterCardInfoRow{
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
	r.logDBResult("GetRitualMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertRitualMonster inserts a new monster card into the database.
func (r *ritualMonsterRepositoryImpl) InsertRitualMonster(ctx context.Context, cardId int64) (sqlc_gen.RitualMonster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertRitualMonster", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.InsertRitualMonster(ctx, cardId)
	if err != nil {
		r.logDBError("InsertRitualMonster", err, zap.Int64("card_id", cardId))
		return sqlc_gen.RitualMonster{}, err
	}

	r.logDBResult("InsertRitualMonster", monster, zap.Int64("card_id", cardId))
	return monster, nil
}
