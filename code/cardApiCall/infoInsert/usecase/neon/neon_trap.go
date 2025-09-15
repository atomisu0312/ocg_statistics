package neon

import (
	"context"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

func (n *neonUseCaseImpl) InsertTrapCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error) {

	tr := transaction.NewTx(n.ProduceConnDB())

	result := int64(0)

	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		// リポジトリの準備
		cardRepo := repository.NewCardRepository(q)
		trapRepo := repository.NewTrapRepository(q)
		trapTypeRepo := repository.NewTrapTypeRepository(q)

		// カードの挿入
		card, err := cardRepo.InsertCard(ctx, cardInfo.ToInsertCardParamsExceptMonster())
		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		// トラップ種別の取得
		trapType, err := trapTypeRepo.GetTrapTypeByNameEn(ctx, cardInfo.Race)
		if err != nil {
			return fmt.Errorf("error get trap type %w", err)
		}

		// Trapテーブルへの挿入
		_, err = trapRepo.InsertTrap(ctx, card.ID, trapType.ID)

		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		result = card.ID
		return nil
	})

	return result, err
}

func (n *neonUseCaseImpl) GetTrapCardByID(ctx context.Context, cardID int64) (cardrecord.TrapCardSelectResult, error) {
	trapRepo := repository.NewTrapRepository(sqlc_gen.New(n.ProduceConnDB()))
	return trapRepo.GetTrapByCardID(ctx, cardID)
}
