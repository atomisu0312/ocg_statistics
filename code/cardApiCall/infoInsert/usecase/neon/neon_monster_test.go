package neon_test

import (
	"context"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/usecase/neon"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func testMonsterCommon(t *testing.T, sampleData cardrecord.StandardCard, attrNameJa string, raceNameJa string, typeLines []string) (cardrecord.MonsterCardSelectResult, error) {
	config.BeforeEachForUnitTest()      // テスト前処理
	defer config.AfterEachForUnitTest() // テスト後処理

	injector := app.SetupDIContainer()
	do.Override(injector, config.TestDbConnection)

	neonUseCase := do.MustInvoke[neon.NeonUseCase](injector)

	resultInsert, err := neonUseCase.InsertMonsterCardInfo(context.Background(), sampleData)

	assert.NoError(t, err)
	assert.NotNil(t, resultInsert)
	assert.NotEqual(t, int32(-1), resultInsert)
	cardID := resultInsert

	resultsGet, err := neonUseCase.GetMonsterCardByID(context.Background(), cardID)
	assert.NoError(t, err)
	assert.NotNil(t, resultsGet)
	assert.Equal(t, sampleData.NameEn, resultsGet.NameEn)
	assert.Equal(t, sampleData.NameJa, resultsGet.NameJa)
	assert.Equal(t, sampleData.DescEn, resultsGet.CardTextEn)
	assert.Equal(t, sampleData.DescJa, resultsGet.CardTextJa)
	assert.Equal(t, sampleData.NeuronID, resultsGet.NeuronID)
	assert.Equal(t, sampleData.TcgID, resultsGet.OcgApiID)
	assert.Equal(t, raceNameJa, resultsGet.RaceNameJa)
	assert.Equal(t, attrNameJa, resultsGet.AttributeNameJa)
	assert.Equal(t, sampleData.Def, resultsGet.Defense)
	assert.Equal(t, sampleData.Atk, resultsGet.Attack)
	assert.Equal(t, sampleData.Level, resultsGet.Level)

	typeLinesEn, err := neonUseCase.GetMonsterTypeLinesEnByCardID(context.Background(), cardID)
	assert.NoError(t, err)
	assert.NotNil(t, typeLinesEn)
	assert.Equal(t, typeLines, typeLinesEn)

	return resultsGet, nil
}

func TestNeonMonsterUseCase(t *testing.T) {
	t.Run("カード情報の挿入&取得テスト（通常モンスター）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Mokey Mokey",
			NameJa:         "もけもけ",
			DescEn:         "An outcast angel.",
			DescJa:         "天使のはみだし者",
			NeuronID:       6018,
			TcgID:          27288416,
			Def:            100,
			Atk:            300,
			Type:           "Normal Monster",
			Level:          1,
			Race:           "Fairy",
			LinkMarkers:    []string{},
			Attribute:      "LIGHT",
			LinkVal:        0,
			TypeLines:      []string{"Normal"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "光"
		RaceNameJa := "天使族"
		TypeLines := []string{"Normal"}

		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)

	})
}
