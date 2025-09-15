package repository

import (
	"context"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/util"
	"go.uber.org/zap"
)

// PendulumMonsterRepository defines the interface for monster card database operations.
type PendulumMonsterRepository interface {
	Repository
	GetPendulumMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.PendulumMonsterSelectResult, error)
	InsertPendulumMonster(ctx context.Context, cardId int64, scale int32, pendulumTextJa string, pendulumTextEn string) (sqlc_gen.PendulumMonster, error)
}

type pendulumMonsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewPendulumMonsterRepository	creates a new instance of PendulumMonsterRepository.
func NewPendulumMonsterRepository(q *sqlc_gen.Queries) PendulumMonsterRepository {
	return NewRepository(func(r *repository) PendulumMonsterRepository {
		return &pendulumMonsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetPendulumMonsterByCardID retrieves a monster card by its card ID.
func (r *pendulumMonsterRepositoryImpl) GetPendulumMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.PendulumMonsterSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetPendulumMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullPendulumMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetPendulumMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.PendulumMonsterSelectResult{}, err
	}

	var result cardrecord.PendulumMonsterSelectResult
	result = *result.FromSelectFullPendulumMonsterCardInfoRow(cardrecord.SelectFullPendulumMonsterCardInfoRow{
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
		Scale:          monster.Scale,
		PendulumTextJa: monster.PendulumTextJa,
		PendulumTextEn: monster.PendulumTextEn,
	})
	r.logDBResult("GetPendulumMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertPendulumMonster inserts a new monster card into the database.
func (r *pendulumMonsterRepositoryImpl) InsertPendulumMonster(ctx context.Context, cardId int64, scale int32, pendulumTextJa string, pendulumTextEn string) (sqlc_gen.PendulumMonster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertPendulumMonster", start, zap.Int64("card_id", cardId))

	params := sqlc_gen.InsertPendulumMonsterParams{
		CardID:         cardId,
		Scale:          util.ParseAsSqlNullInt32WithTreatZeroAsNull(scale),
		PendulumTextJa: util.ParseAsSqlNullString(pendulumTextJa),
		PendulumTextEn: util.ParseAsSqlNullString(pendulumTextEn),
	}
	monster, err := r.queries.InsertPendulumMonster(ctx, params)
	if err != nil {
		r.logDBError("InsertPendulumMonster", err, zap.Int64("card_id", cardId))
		return sqlc_gen.PendulumMonster{}, err
	}

	r.logDBResult("InsertPendulumMonster", monster, zap.Int64("card_id", cardId))
	return monster, nil
}
