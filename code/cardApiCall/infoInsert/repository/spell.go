package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// SpellRepository defines the interface for spell card database operations.
type SpellRepository interface {
	Repository
	GetSpellByCardID(ctx context.Context, cardId int64) (cardrecord.SpellCardSelectResult, error)
	InsertSpell(ctx context.Context, cardId int64, spellTypeId int32) (sqlc_gen.Spell, error)
}

type spellRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewSpellRepository creates a new instance of SpellRepository.
func NewSpellRepository(q *sqlc_gen.Queries) SpellRepository {
	return NewRepository(func(r *repository) SpellRepository {
		return &spellRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetSpellByCardID retrieves a spell card by its card ID.
func (r *spellRepositoryImpl) GetSpellByCardID(ctx context.Context, cardId int64) (cardrecord.SpellCardSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetSpellByCardID", start, zap.Int64("card_id", cardId))

	spell, err := r.queries.SelectFullSpellCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetSpellByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.SpellCardSelectResult{}, err
	}
	var result cardrecord.SpellCardSelectResult
	result = *result.FromSelectFullSpellCardInfoRow(cardrecord.SelectFullSpellCardInfoRow(spell))
	r.logDBResult("GetSpellByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertSpell inserts a new spell card into the database.
func (r *spellRepositoryImpl) InsertSpell(ctx context.Context, cardId int64, spellTypeId int32) (sqlc_gen.Spell, error) {
	start := time.Now()
	defer r.logDBOperation("InsertSpell", start, zap.Int64("card_id", cardId), zap.Int32("spell_type_id", spellTypeId))

	spell, err := r.queries.InsertSpell(ctx, sqlc_gen.InsertSpellParams{
		CardID:      cardId,
		SpellTypeID: sql.NullInt32{Int32: spellTypeId, Valid: true},
	})
	if err != nil {
		r.logDBError("InsertSpell", err, zap.Int64("card_id", cardId), zap.Int32("spell_type_id", spellTypeId))
		return sqlc_gen.Spell{}, err
	}

	r.logDBResult("InsertSpell", spell, zap.Int64("card_id", cardId), zap.Int32("spell_type_id", spellTypeId))
	return spell, nil
}
