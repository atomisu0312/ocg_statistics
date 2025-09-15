package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/util"
	"go.uber.org/zap"
)

// MonsterRepository defines the interface for monster card database operations.
type MonsterRepository interface {
	Repository
	GetMonsterByCardID(ctx context.Context, cardId int64) (cardrecord.MonsterCardSelectResult, error)
	GetMonsterByOcgApiID(ctx context.Context, ocgApiId int64) (cardrecord.MonsterCardSelectResult, error)
	GetMonsterByNeuronID(ctx context.Context, neuronId int64) (cardrecord.MonsterCardSelectResult, error)
	InsertMonster(ctx context.Context, cardId int64, raceId int32, attributeId int32, attack int32, defense int32, level int32, typeIds []int32) (sqlc_gen.Monster, error)
	GetMonsterTypeLineByCardID(ctx context.Context, cardId int64) (cardrecord.MonsterTypeLineSelectResult, error)
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

// GetMonsterByOcgApiID retrieves a monster card by its OCG API ID.
func (r *monsterRepositoryImpl) GetMonsterByOcgApiID(ctx context.Context, ocgApiId int64) (cardrecord.MonsterCardSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterByOcgApiID", start, zap.Int64("ocg_api_id", ocgApiId))

	monster, err := r.queries.SelectFullMonsterCardInfoByOcgApiID(ctx, util.ParseAsSqlNullInt64WithTreatZeroAsNull(ocgApiId))
	if err != nil {
		r.logDBError("GetMonsterByOcgApiID", err, zap.Int64("ocg_api_id", ocgApiId))
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
	r.logDBResult("GetMonsterByOcgApiID", result, zap.Int64("ocg_api_id", ocgApiId))
	return result, nil
}

// GetMonsterByNeuronID retrieves a monster card by its neuron ID.
func (r *monsterRepositoryImpl) GetMonsterByNeuronID(ctx context.Context, neuronId int64) (cardrecord.MonsterCardSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterByNeuronID", start, zap.Int64("neuron_id", neuronId))

	monster, err := r.queries.SelectFullMonsterCardInfoByNeuronID(ctx, util.ParseAsSqlNullInt64WithTreatZeroAsNull(neuronId))
	if err != nil {
		r.logDBError("GetMonsterByNeuronID", err, zap.Int64("neuron_id", neuronId))
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
	r.logDBResult("GetMonsterByNeuronID", result, zap.Int64("neuron_id", neuronId))
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

func (r *monsterRepositoryImpl) GetMonsterTypeLineByCardID(ctx context.Context, cardId int64) (cardrecord.MonsterTypeLineSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetMonsterTypeLineByCardID", start, zap.Int64("card_id", cardId))

	monsterTypeLine, err := r.queries.SelectMonsterTypeLineByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetMonsterTypeLineByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.MonsterTypeLineSelectResult{}, err
	}

	var result cardrecord.MonsterTypeLineSelectResult
	result = *result.FromSelectMonsterTypeLineByCardIDRow(cardrecord.SelectMonsterTypeLineByCardIDRow{
		ID:         monsterTypeLine.ID,
		OcgApiID:   monsterTypeLine.OcgApiID.Int64,
		NeuronID:   monsterTypeLine.NeuronID.Int64,
		IsNormal:   monsterTypeLine.IsNormal,
		IsEffect:   monsterTypeLine.IsEffect,
		IsToon:     monsterTypeLine.IsToon,
		IsSpirit:   monsterTypeLine.IsSpirit,
		IsUnion:    monsterTypeLine.IsUnion,
		IsDual:     monsterTypeLine.IsDual,
		IsTuner:    monsterTypeLine.IsTuner,
		IsReverse:  monsterTypeLine.IsReverse,
		IsRitual:   monsterTypeLine.IsRitual,
		IsXyz:      monsterTypeLine.IsXyz,
		IsSynchro:  monsterTypeLine.IsSynchro,
		IsFusion:   monsterTypeLine.IsFusion,
		IsLink:     monsterTypeLine.IsLink,
		IsPendulum: monsterTypeLine.IsPendulum,
	})
	r.logDBResult("GetMonsterTypeLineByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}
