package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// MonsterRepository defines the interface for monster card database operations.
type MonsterRepository interface {
	Repository
	GetMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.MonsterCardSelectResult, error)
	InsertMonster(ctx context.Context, cardId int64, raceId int32, attributeId int32, attack int32, defense int32, level int32, typeIds []int32) (sqlc_gen.Monster, error)
}

type monsterRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewMonsterRepository creates a new instance of MonsterRepository.
func NewMonsterRepository(q *sqlc_gen.Queries) MonsterRepository {
	return NewRepository(func(r *repository) MonsterRepository {
		return &monsterRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetMonsterByCardID retrieves a monster card by its card ID.
func (r *monsterRepositoryImpl) GetMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.MonsterCardSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterByCardID", start, zap.Int64("card_id", cardId))

	monster, err := r.queries.SelectFullMonsterCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetMonsterByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.MonsterCardSelectResult{}, err
	}

	var result cardrecord.MonsterCardSelectResult
	result = *result.FromSelectFullMonsterCardInfoRow(cardrecord.SelectFullMonsterCardInfoRow{
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
	})
	r.logDBResult("GetMonsterByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertMonster inserts a new monster card into the database.
func (r *monsterRepositoryImpl) InsertMonster(ctx context.Context, cardId int64, raceId int32, attributeId int32, attack int32, defense int32, level int32, typeIds []int32) (sqlc_gen.Monster, error) {
	start := time.Now()
	defer r.logDBOperation("InsertMonster", start, zap.Int64("card_id", cardId), zap.Int32("race_id", raceId), zap.Int32("attribute_id", attributeId), zap.Int32("attack", attack), zap.Int32("defense", defense), zap.Int32("level", level), zap.Any("type_ids", typeIds))

	monster, err := r.queries.InsertMonster(ctx, sqlc_gen.InsertMonsterParams{
		CardID:      cardId,
		RaceID:      sql.NullInt32{Int32: raceId, Valid: true},
		AttributeID: sql.NullInt32{Int32: attributeId, Valid: true},
		Attack:      sql.NullInt32{Int32: attack, Valid: true},
		Defense:     sql.NullInt32{Int32: defense, Valid: true},
		Level:       sql.NullInt32{Int32: level, Valid: true},
		TypeIds:     typeIds,
	})
	if err != nil {
		r.logDBError("InsertMonster", err, zap.Int64("card_id", cardId), zap.Int32("race_id", raceId), zap.Int32("attribute_id", attributeId), zap.Int32("attack", attack), zap.Int32("defense", defense), zap.Int32("level", level), zap.Any("type_ids", typeIds))
		return sqlc_gen.Monster{}, err
	}

	r.logDBResult("InsertMonster", monster, zap.Int64("card_id", cardId), zap.Int32("race_id", raceId), zap.Int32("attribute_id", attributeId), zap.Int32("attack", attack), zap.Int32("defense", defense), zap.Int32("level", level), zap.Any("type_ids", typeIds))
	return monster, nil
}
