package usecase

import (
	"context"
	"fmt"
	"strings"

	"atomisu.com/ocg-statics/infoInsert/dto/carddto"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

func (n *neonUseCaseImpl) InsertTrapCardInfo(ctx context.Context, cardInfo carddto.StandardCard) (int64, error) {

	tr := transaction.NewTx(n.dbConn.DB)

	result := int64(0)

	// 一応のクレンジング処理
	race := strings.ToLower(strings.TrimSpace(cardInfo.Race))

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
		trapType, err := trapTypeRepo.GetTrapTypeByNameEn(ctx, race)
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

func (n *neonUseCaseImpl) GetTrapCardByID(ctx context.Context, cardID int64) (carddto.TrapCardSelectResult, error) {
	trapRepo := repository.NewTrapRepository(sqlc_gen.New(n.dbConn.DB))
	return trapRepo.GetTrapByCardID(ctx, cardID)
}
