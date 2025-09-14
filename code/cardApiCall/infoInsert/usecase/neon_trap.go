package usecase

import (
	"context"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/dto/carddto"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

func (n *neonUseCaseImpl) InsertTrapCardInfo(ctx context.Context, cardInfo carddto.StandardCard) (int64, error) {

	tr := transaction.NewTx(n.dbConn.DB)

	result := int64(0)

	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		// リポジトリの準備
		cardRepo := repository.NewCardRepository(q)
		trapRepo := repository.NewTrapRepository(q)

		// カードの挿入
		card, err := cardRepo.InsertCard(ctx, cardInfo.ToInsertCardParamsExceptMonster())
		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		// トラップの挿入
		// TODO: 適切にトラップ種別IDを判別する
		_, err = trapRepo.InsertTrap(ctx, card.ID, 3)

		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		result = card.ID
		return nil
	})

	return result, err
}

func (n *neonUseCaseImpl) GetTrapCardByID(ctx context.Context, cardID int64) (carddto.TrapCardSelectResult, error) {
	trapRepo := repository.NewTrapRepository(sqlc_gen.New(n.dbConn.DB))
	return trapRepo.GetTrapByCardID(ctx, cardID)
}
