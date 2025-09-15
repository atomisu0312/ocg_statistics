package neon

import (
	"context"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

func (n *neonUseCaseImpl) InsertSpellCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error) {

	tr := transaction.NewTx(n.ProduceConnDB())

	result := int64(0)

	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		// リポジトリの準備
		cardRepo := repository.NewCardRepository(q)
		spellRepo := repository.NewSpellRepository(q)
		spellTypeRepo := repository.NewSpellTypeRepository(q)

		// カードの挿入
		card, err := cardRepo.InsertCard(ctx, cardInfo.ToInsertCardParamsExceptMonster())
		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		// 魔法種別の取得
		spellType, err := spellTypeRepo.GetSpellTypeByNameEn(ctx, cardInfo.Race)
		if err != nil {
			return fmt.Errorf("error get spell type %w", err)
		}

		// Spellテーブルへの挿入
		_, err = spellRepo.InsertSpell(ctx, card.ID, spellType.ID)

		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		result = card.ID
		return nil
	})

	return result, err
}

func (n *neonUseCaseImpl) GetSpellCardByID(ctx context.Context, cardID int64) (cardrecord.SpellCardSelectResult, error) {
	spellRepo := repository.NewSpellRepository(sqlc_gen.New(n.ProduceConnDB()))
	return spellRepo.GetSpellByCardID(ctx, cardID)
}
