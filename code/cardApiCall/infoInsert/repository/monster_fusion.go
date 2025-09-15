package repository

import (
	"context"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// MonsterRepository defines the interface for monster card database operations.
type FusionMonsterRepository interface {
	Repository
	GetFusionMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.FusionMonsterSelectResult, error)
	InsertFusionMonster(ctx context.Context, cardId int64) (sqlc_gen.FusionMonster, error)
}

type fusionMonsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewFusionMonsterRepository	creates a new instance of FusionMonsterRepository.
func NewFusionMonsterRepository(q *sqlc_gen.Queries) FusionMonsterRepository {
	return NewRepository(func(r *repository) FusionMonsterRepository {
		return &fusionMonsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetMonsterByCardID retrieves a monster card by its card ID.
func (r *fusionMonsterRepositoryImpl) GetFusionMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.FusionMonsterSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetFusionMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetFusionMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.FusionMonsterSelectResult{}, err
	}

	var result cardrecord.FusionMonsterSelectResult
	result = *result.FromSelectFullFusionMonsterCardInfoRow(cardrecord.SelectFullFusionMonsterCardInfoRow{
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
	r.logDBResult("GetFusionMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertMonster inserts a new monster card into the database.
func (r *fusionMonsterRepositoryImpl) InsertFusionMonster(ctx context.Context, cardId int64) (sqlc_gen.FusionMonster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertFusionMonster", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.InsertFusionMonster(ctx, cardId)
	if err != nil {
		r.logDBError("InsertFusionMonster", err, zap.Int64("card_id", cardId))
		return sqlc_gen.FusionMonster{}, err
	}

	r.logDBResult("InsertFusionMonster", monster, zap.Int64("card_id", cardId))
	return monster, nil
}
