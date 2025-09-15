package repository

import (
	"context"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/util"
	"go.uber.org/zap"
)

// LinkMonsterRepository defines the interface for monster card database operations.
type LinkMonsterRepository interface {
	Repository
	GetLinkMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.LinkMonsterSelectResult, error)
	InsertLinkMonster(ctx context.Context, cardId int64, linkMarker int32) (sqlc_gen.LinkMonster, error)
}

type linkMonsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewLinkMonsterRepository	creates a new instance of LinkMonsterRepository.
func NewLinkMonsterRepository(q *sqlc_gen.Queries) LinkMonsterRepository {
	return NewRepository(func(r *repository) LinkMonsterRepository {
		return &linkMonsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetLinkMonsterByCardID retrieves a monster card by its card ID.
func (r *linkMonsterRepositoryImpl) GetLinkMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.LinkMonsterSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetLinkMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullLinkMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetLinkMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.LinkMonsterSelectResult{}, err
	}

	var result cardrecord.LinkMonsterSelectResult
	result = *result.FromSelectFullLinkMonsterCardInfoRow(cardrecord.SelectFullLinkMonsterCardInfoRow{
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
		LinkMarker: monster.LinkMarker,
	})
	r.logDBResult("GetLinkMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertLinkMonster inserts a new monster card into the database.
func (r *linkMonsterRepositoryImpl) InsertLinkMonster(ctx context.Context, cardId int64, linkMarker int32) (sqlc_gen.LinkMonster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertLinkMonster", start, zap.Int64("card_id", cardId))

	arg := sqlc_gen.InsertLinkMonsterParams{
		CardID:     cardId,
		LinkMarker: util.ParseAsSqlNullInt32WithTreatZeroAsNull(linkMarker),
	}
	monster, err := r.queries.InsertLinkMonster(ctx, arg)
	if err != nil {
		r.logDBError("InsertLinkMonster", err, zap.Int64("card_id", cardId))
		return sqlc_gen.LinkMonster{}, err
	}

	r.logDBResult("InsertLinkMonster", monster, zap.Int64("card_id", cardId))
	return monster, nil
}
