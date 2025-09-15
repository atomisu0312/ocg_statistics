package neon_test

import (
	"context"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/dto/carddto"
	"atomisu.com/ocg-statics/infoInsert/usecase/neon"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestNeonTrapUseCase(t *testing.T) {
	//t.Parallel()

	t.Run("カード情報の挿入&取得テスト（カウンター罠）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Solemn Judgment",
			NameJa:   "神の宣告",
			DescEn:   "Sample Text.",
			DescJa:   "サンプルテキスト。",
			Type:     "Trap Card",
			NeuronID: 8916,
			TcgID:    41420027,
			Race:     "Counter",
		}

		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[neon.NeonUseCase](injector)

		results, err := neonUseCase.InsertTrapCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetTrapCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "カウンター", results2.TrapTypeNameJa)

	})

	t.Run("カード情報の挿入&取得テスト（永続罠）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Skill Drain",
			NameJa:   "スキルドレイン",
			DescEn:   "Sample Text.",
			DescJa:   "サンプルテキスト。",
			Type:     "Trap Card",
			NeuronID: 5740,
			TcgID:    82732705,
			Race:     "Continuous",
		}

		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[neon.NeonUseCase](injector)

		results, err := neonUseCase.InsertTrapCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetTrapCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "永続", results2.TrapTypeNameJa)

	})

	t.Run("カード情報の挿入&取得テスト（通常罠）", func(t *testing.T) {
		sampleData := carddto.StandardCard{
			NameEn:   "Welcome Labrynth",
			NameJa:   "ウェルカム・ラビュリンス",
			DescEn:   "Sample Text.",
			DescJa:   "サンプルテキスト。",
			Type:     "Trap Card",
			NeuronID: 17369,
			TcgID:    5380979,
			Race:     "Normal",
		}

		config.BeforeEachForUnitTest()      // テスト前処理
		defer config.AfterEachForUnitTest() // テスト後処理

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neonUseCase := do.MustInvoke[neon.NeonUseCase](injector)

		results, err := neonUseCase.InsertTrapCardInfo(context.Background(), sampleData)

		assert.NoError(t, err)
		assert.NotNil(t, results)

		results2, err := neonUseCase.GetTrapCardByID(context.Background(), results)
		assert.NoError(t, err)
		assert.NotNil(t, results2)
		assert.Equal(t, sampleData.NameEn, results2.NameEn)
		assert.Equal(t, sampleData.NameJa, results2.NameJa)
		assert.Equal(t, sampleData.DescEn, results2.CardTextEn)
		assert.Equal(t, sampleData.DescJa, results2.CardTextJa)
		assert.Equal(t, sampleData.NeuronID, results2.NeuronID)
		assert.Equal(t, sampleData.TcgID, results2.OcgApiID)
		assert.Equal(t, "通常", results2.TrapTypeNameJa)

	})
}
