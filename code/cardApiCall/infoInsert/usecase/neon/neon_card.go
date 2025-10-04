package neon

import (
	"context"
	"fmt"
	"strings"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
)

func isMonsterCard(cardInfo cardrecord.StandardCard) bool {
	return strings.Contains(cardInfo.Type, "Monster")
}

func isTrapCard(cardInfo cardrecord.StandardCard) bool {
	return strings.Contains(cardInfo.Type, "Trap")
}

func isSpellCard(cardInfo cardrecord.StandardCard) bool {
	return strings.Contains(cardInfo.Type, "Spell")
}

// StandardCardを引数として、魔法・罠・モンスターを判定して適切なテーブルに挿入する
func (n *neonUseCaseImpl) InsertCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error) {
	if isMonsterCard(cardInfo) {
		return n.InsertMonsterCardInfo(ctx, cardInfo)
	}
	if isTrapCard(cardInfo) {
		return n.InsertTrapCardInfo(ctx, cardInfo)
	}
	if isSpellCard(cardInfo) {
		return n.InsertSpellCardInfo(ctx, cardInfo)
	}
	return 0, fmt.Errorf("invalid card type")
}
