package usecase_test

import (
	"context"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/dto/carddto"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestNeonSpellUseCase(t *testing.T) {
	//t.Parallel()

	t.Run("カード情報の挿入&取得テスト（通常魔法）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Raigeki",
			NameJa:   "サンダー・ボルト",
			DescEn:   "Sample Text.",
			DescJa:   "サンプルテキスト。",
			Type:     "Spell Card",
			NeuronID: 4343,
			TcgID:    12580477,
			Race:     "Normal",
		}

		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[usecase.NeonUseCase](injector)

		results, err := neonUseCase.InsertSpellCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetSpellCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "通常", results2.SpellTypeNameJa)

	})

	t.Run("カード情報の挿入&取得テスト（永続魔法）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Continuous Spell Card",
			NameJa:   "永続魔法カード",
			DescEn:   "Sample Text for Continuous Spell.",
			DescJa:   "永続魔法のサンプルテキスト。",
			Type:     "Spell Card",
			NeuronID: 10001,
			TcgID:    10000001,
			Race:     "Continuous",
		}

		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[usecase.NeonUseCase](injector)

		results, err := neonUseCase.InsertSpellCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetSpellCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "永続", results2.SpellTypeNameJa)
	})

	t.Run("カード情報の挿入&取得テスト（装備魔法）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Equip Spell Card",
			NameJa:   "装備魔法カード",
			DescEn:   "Sample Text for Equip Spell.",
			DescJa:   "装備魔法のサンプルテキスト。",
			Type:     "Spell Card",
			NeuronID: 10002,
			TcgID:    10000002,
			Race:     "Equip",
		}

		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[usecase.NeonUseCase](injector)

		results, err := neonUseCase.InsertSpellCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetSpellCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "装備", results2.SpellTypeNameJa)
	})

	t.Run("カード情報の挿入&取得テスト（フィールド魔法）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Field Spell Card",
			NameJa:   "フィールド魔法カード",
			DescEn:   "Sample Text for Field Spell.",
			DescJa:   "フィールド魔法のサンプルテキスト。",
			Type:     "Spell Card",
			NeuronID: 10003,
			TcgID:    10000003,
			Race:     "Field",
		}

		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[usecase.NeonUseCase](injector)

		results, err := neonUseCase.InsertSpellCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetSpellCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "フィールド", results2.SpellTypeNameJa)
	})

	t.Run("カード情報の挿入&取得テスト（速攻魔法）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Quick-Play Spell Card",
			NameJa:   "速攻魔法カード",
			DescEn:   "Sample Text for Quick-Play Spell.",
			DescJa:   "速攻魔法のサンプルテキスト。",
			Type:     "Spell Card",
			NeuronID: 10004,
			TcgID:    10000004,
			Race:     "Quick-Play",
		}

		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[usecase.NeonUseCase](injector)

		results, err := neonUseCase.InsertSpellCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetSpellCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "速攻", results2.SpellTypeNameJa)
	})

	t.Run("カード情報の挿入&取得テスト（儀式魔法）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Ritual Spell Card",
			NameJa:   "儀式魔法カード",
			DescEn:   "Sample Text for Ritual Spell.",
			DescJa:   "儀式魔法のサンプルテキスト。",
			Type:     "Spell Card",
			NeuronID: 10005,
			TcgID:    10000005,
			Race:     "Ritual",
		}

		config.BeforeEachForUnitTest()
		defer config.AfterEachForUnitTest()

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[usecase.NeonUseCase](injector)

		results, err := neonUseCase.InsertSpellCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetSpellCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "儀式", results2.SpellTypeNameJa)
	})
}
