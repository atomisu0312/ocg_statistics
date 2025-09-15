package repository

import (
	"context"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// SynchroMonsterRepository defines the interface for monster card database operations.
type SynchroMonsterRepository interface {
	Repository
	GetSynchroMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.SynchroMonsterSelectResult, error)
	InsertSynchroMonster(ctx context.Context, cardId int64) (sqlc_gen.SynchroMonster, error)
}

type synchroMonsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewSynchroMonsterRepository	creates a new instance of SynchroMonsterRepository.
func NewSynchroMonsterRepository(q *sqlc_gen.Queries) SynchroMonsterRepository {
	return NewRepository(func(r *repository) SynchroMonsterRepository {
		return &synchroMonsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetSynchroMonsterByCardID retrieves a monster card by its card ID.
func (r *synchroMonsterRepositoryImpl) GetSynchroMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.SynchroMonsterSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetSynchroMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullSynchroMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetSynchroMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.SynchroMonsterSelectResult{}, err
	}

	var result cardrecord.SynchroMonsterSelectResult
	result = *result.FromSelectFullSynchroMonsterCardInfoRow(cardrecord.SelectFullSynchroMonsterCardInfoRow{
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
	r.logDBResult("GetSynchroMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertSynchroMonster inserts a new monster card into the database.
func (r *synchroMonsterRepositoryImpl) InsertSynchroMonster(ctx context.Context, cardId int64) (sqlc_gen.SynchroMonster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertSynchroMonster", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.InsertSynchroMonster(ctx, cardId)
	if err != nil {
		r.logDBError("InsertSynchroMonster", err, zap.Int64("card_id", cardId))
		return sqlc_gen.SynchroMonster{}, err
	}

	r.logDBResult("InsertSynchroMonster", monster, zap.Int64("card_id", cardId))
	return monster, nil
}
