package repository

import (
	"context"
	"database/sql"
	"time"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"go.uber.org/zap"
)

// TrapRepository defines the interface for trap card database operations.
type TrapRepository interface {
	Repository
	GetTrapByCardID(ctx context.Context, cardId int64) (cardrecord.TrapCardSelectResult, error)
	InsertTrap(ctx context.Context, cardId int64, trapTypeId int32) (sqlc_gen.Trap, error)
}

type trapRepositoryImpl struct {
	*repository
	queries *sqlc_gen.Queries
}

// NewTrapRepository creates a new instance of TrapRepository.
func NewTrapRepository(q *sqlc_gen.Queries) TrapRepository {
	return NewRepository(func(r *repository) TrapRepository {
		return &trapRepositoryImpl{
			repository: r,
			queries:    q,
		}
	})
}

// GetTrapByCardID retrieves a trap card by its card ID.
func (r *trapRepositoryImpl) GetTrapByCardID(ctx context.Context, cardId int64) (cardrecord.TrapCardSelectResult, error) {
	start := time.Now()
	defer r.logDBOperation("GetTrapByCardID", start, zap.Int64("card_id", cardId))

	trap, err := r.queries.SelectFullTrapCardInfoByCardID(ctx, cardId)
	if err != nil {
		r.logDBError("GetTrapByCardID", err, zap.Int64("card_id", cardId))
		return cardrecord.TrapCardSelectResult{}, err
	}
	var result cardrecord.TrapCardSelectResult
	result = *result.FromSelectFullTrapCardInfoRow(cardrecord.SelectFullTrapCardInfoRow(trap))
	r.logDBResult("GetTrapByCardID", result, zap.Int64("card_id", cardId))
	return result, nil
}

// InsertTrap inserts a new trap card into the database.
func (r *trapRepositoryImpl) InsertTrap(ctx context.Context, cardId int64, trapTypeId int32) (sqlc_gen.Trap, error) {
	start := time.Now()
	defer r.logDBOperation("InsertTrap", start, zap.Int64("card_id", cardId), zap.Int32("trap_type_id", trapTypeId))

	trap, err := r.queries.InsertTrap(ctx, sqlc_gen.InsertTrapParams{
		CardID:     cardId,
		TrapTypeID: sql.NullInt32{Int32: trapTypeId, Valid: true},
	})
	if err != nil {
		r.logDBError("InsertTrap", err, zap.Int64("card_id", cardId), zap.Int32("trap_type_id", trapTypeId))
		return sqlc_gen.Trap{}, err
	}

	r.logDBResult("InsertTrap", trap, zap.Int64("card_id", cardId), zap.Int32("trap_type_id", trapTypeId))
	return trap, nil
}
