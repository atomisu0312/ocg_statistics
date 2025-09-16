package neon

import (
	"context"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

// InsertMonsterCardInfoはStandardCardを引数として、適切なテーブル群に必要なレコードを挿入する
func (n *neonUseCaseImpl) InsertMonsterCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error) {
	tr := transaction.NewTx(n.ProduceConnDB())

	result := int64(0)

	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {
		// リポジトリの準備
		cardRepo := repository.NewCardRepository(q)
		monsterRepo := repository.NewMonsterRepository(q)

		// カードの挿入
		card, err := cardRepo.InsertCard(ctx, cardInfo.ToInsertCardParamsExceptMonster())
		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		raceId := int32(9)      // TODO
		attributeId := int32(1) // TODO
		typeIds := []int32{1}   // TODO

		// モンスターの挿入
		_, err = monsterRepo.InsertMonster(ctx, card.ID, raceId, attributeId, cardInfo.Atk, cardInfo.Def, cardInfo.Level, typeIds)
		if err != nil {
			return fmt.Errorf("error create monster %w", err)
		}

		result = card.ID
		return nil
	})

	return result, err
}

// GetMonsterCardByID は、モンスターのカードを取得。
func (n *neonUseCaseImpl) GetMonsterCardByID(ctx context.Context, cardID int64) (cardrecord.MonsterCardSelectResult, error) {
	monsterRepo := repository.NewMonsterRepository(sqlc_gen.New(n.ProduceConnDB()))
	return monsterRepo.GetMonsterByCardID(ctx, cardID)
}

// GetMonsterTypeLinesEnByCardID は、モンスターの種類を取得。
func (n *neonUseCaseImpl) GetMonsterTypeLinesEnByCardID(ctx context.Context, cardID int64) ([]string, error) {
	monsterRepo := repository.NewMonsterRepository(sqlc_gen.New(n.ProduceConnDB()))
	typeLineNames := []string{}
	typeLineSelectResult, err := monsterRepo.GetMonsterTypeLineByCardID(ctx, cardID)
	if err != nil {
		return nil, err
	}

	if typeLineSelectResult.IsNormal {
		typeLineNames = append(typeLineNames, "Normal")
	}
	if typeLineSelectResult.IsEffect {
		typeLineNames = append(typeLineNames, "Effect")
	}
	if typeLineSelectResult.IsToon {
		typeLineNames = append(typeLineNames, "Toon")
	}
	if typeLineSelectResult.IsSpirit {
		typeLineNames = append(typeLineNames, "Spirit")
	}
	if typeLineSelectResult.IsUnion {
		typeLineNames = append(typeLineNames, "Union")
	}
	if typeLineSelectResult.IsDual {
		typeLineNames = append(typeLineNames, "Dual")
	}
	if typeLineSelectResult.IsTuner {
		typeLineNames = append(typeLineNames, "Tuner")
	}
	if typeLineSelectResult.IsReverse {
		typeLineNames = append(typeLineNames, "Reverse")
	}
	if typeLineSelectResult.IsRitual {
		typeLineNames = append(typeLineNames, "Ritual")
	}
	if typeLineSelectResult.IsXyz {
		typeLineNames = append(typeLineNames, "Xyz")
	}
	if typeLineSelectResult.IsSynchro {
		typeLineNames = append(typeLineNames, "Synchro")
	}
	if typeLineSelectResult.IsFusion {
		typeLineNames = append(typeLineNames, "Fusion")
	}
	if typeLineSelectResult.IsLink {
		typeLineNames = append(typeLineNames, "Link")
	}
	if typeLineSelectResult.IsPendulum {
		typeLineNames = append(typeLineNames, "Pendulum")
	}

	return typeLineNames, nil
}
