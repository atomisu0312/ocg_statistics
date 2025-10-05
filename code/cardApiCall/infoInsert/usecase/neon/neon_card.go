package neon

import (
	"context"
	"fmt"
	"strings"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

func isMonsterCardWithStandardCard(cardInfo cardrecord.StandardCard) bool {
	return strings.Contains(cardInfo.Type, "Monster")
}

func isTrapCardWithStandardCard(cardInfo cardrecord.StandardCard) bool {
	return strings.Contains(cardInfo.Type, "Trap")
}

func isSpellCardWithStandardCard(cardInfo cardrecord.StandardCard) bool {
	return strings.Contains(cardInfo.Type, "Spell")
}

// StandardCardを引数として、魔法・罠・モンスターを判定して適切なテーブルに挿入する
func (n *neonUseCaseImpl) InsertCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error) {
	if isMonsterCardWithStandardCard(cardInfo) {
		return n.InsertMonsterCardInfo(ctx, cardInfo)
	}
	if isTrapCardWithStandardCard(cardInfo) {
		return n.InsertTrapCardInfo(ctx, cardInfo)
	}
	if isSpellCardWithStandardCard(cardInfo) {
		return n.InsertSpellCardInfo(ctx, cardInfo)
	}
	return 0, fmt.Errorf("invalid card type")
}

func (n *neonUseCaseImpl) GetCardPatternByCardID(ctx context.Context, cardID int64) (cardrecord.CardPatternSelectResult, error) {
	tr := transaction.NewTx(n.ProduceConnDB())

	result := cardrecord.CardPatternSelectResult{}

	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		cardRepo := repository.NewCardRepository(q)
		cardPattern, err := cardRepo.GetCardPatternByCardID(ctx, cardID)
		if err != nil {
			return fmt.Errorf("error get card pattern %w", err)
		}
		result = *result.FromSelectCardPatternByCardIDRow(cardPattern)
		return nil
	})
	return result, err
}
