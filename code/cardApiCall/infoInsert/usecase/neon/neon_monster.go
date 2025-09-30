package neon

import (
	"context"
	"fmt"
	"slices"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/repository"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/transaction"
)

// InsertMonsterCardInfoはStandardCardを引数として、適切なテーブル群に必要なレコードを挿入する
func (n *neonUseCaseImpl) InsertMonsterCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error) {
	tr := transaction.NewTx(n.ProduceConnDB())

	result := int64(0)

	// 以下の一連の流れをトランザクション境界内で実行
	err := tr.ExecTx(ctx, func(q *sqlc_gen.Queries) error {

		// リポジトリの準備
		cardRepo := repository.NewCardRepository(q)
		monsterRepo := repository.NewMonsterRepository(q)
		attributeRepo := repository.NewAttributeRepository(q)
		raceRepo := repository.NewRaceRepository(q)
		monsterTypeRepo := repository.NewMonsterTypeRepository(q)
		fusionMonsterRepo := repository.NewFusionMonsterRepository(q)
		synchroMonsterRepo := repository.NewSynchroMonsterRepository(q)
		xyzMonsterRepo := repository.NewXyzMonsterRepository(q)
		ritualMonsterRepo := repository.NewRitualMonsterRepository(q)

		// カードの挿入
		card, err := cardRepo.InsertCard(ctx, cardInfo.ToInsertCardParamsExceptMonster())
		if err != nil {
			return fmt.Errorf("error create card %w", err)
		}

		// 種族をIDに変換
		raceEntity, err := raceRepo.GetRaceByNameEn(ctx, cardInfo.Race)
		if err != nil {
			return fmt.Errorf("error get race %w", err)
		}
		raceID := raceEntity.ID

		// 属性をIDに変換
		attributeEntity, err := attributeRepo.GetAttributeByNameEn(ctx, cardInfo.Attribute)
		if err != nil {
			return fmt.Errorf("error get attribute %w", err)
		}
		attributeID := attributeEntity.ID

		// モンスターのタイプをIDに変換
		typeIDs := []int32{}
		for _, typeLine := range cardInfo.TypeLines {
			typeEntity, err := monsterTypeRepo.GetMonsterTypeByNameEn(ctx, typeLine)
			if err != nil {
				continue
			}
			// 無効値なのでスキップ
			if typeEntity.ID == 0 {
				continue
			}
			typeIDs = append(typeIDs, typeEntity.ID)
		}

		// モンスターの挿入
		_, err = monsterRepo.InsertMonster(ctx, card.ID, raceID, attributeID, cardInfo.Atk, cardInfo.Def, cardInfo.Level, typeIDs)
		if err != nil {
			return fmt.Errorf("error create monster %w", err)
		}

		// モンスターの種類がFusionの場合は、Fusionテーブルへの挿入
		if slices.Contains(cardInfo.TypeLines, "Fusion") {
			_, err = fusionMonsterRepo.InsertFusionMonster(ctx, card.ID)
			if err != nil {
				return fmt.Errorf("error create fusion monster %w", err)
			}
		}

		// モンスターの種類がFusionの場合は、Fusionテーブルへの挿入
		if slices.Contains(cardInfo.TypeLines, "Synchro") {
			_, err = synchroMonsterRepo.InsertSynchroMonster(ctx, card.ID)
			if err != nil {
				return fmt.Errorf("error create synchro monster %w", err)
			}
		}

		// モンスターの種類がXyzの場合は、Xyzテーブルへの挿入
		if slices.Contains(cardInfo.TypeLines, "Xyz") {
			_, err = xyzMonsterRepo.InsertXyzMonster(ctx, card.ID)
			if err != nil {
				return fmt.Errorf("error create xyz monster %w", err)
			}
		}

		// モンスターの種類がRitualの場合は、Ritualテーブルへの挿入
		if slices.Contains(cardInfo.TypeLines, "Ritual") {
			_, err = ritualMonsterRepo.InsertRitualMonster(ctx, card.ID)
			if err != nil {
				return fmt.Errorf("error create ritual monster %w", err)
			}
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
	if typeLineSelectResult.IsGemini {
		typeLineNames = append(typeLineNames, "Gemini")
	}
	if typeLineSelectResult.IsTuner {
		typeLineNames = append(typeLineNames, "Tuner")
	}
	if typeLineSelectResult.IsFlip {
		typeLineNames = append(typeLineNames, "Flip")
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
