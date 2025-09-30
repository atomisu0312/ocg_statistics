package neon_test

import (
	"context"
	"fmt"
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
	assert.ElementsMatch(t, typeLines, typeLinesEn)

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

	t.Run("カード情報の挿入&取得テスト（通常モンスター）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Melffy Rabby",
			NameJa:         "メルフィー・ラビィ",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       15250,
			TcgID:          20129614,
			Def:            2100,
			Atk:            0,
			Type:           "Normal Monster",
			Level:          2,
			Race:           "Beast",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Normal"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "獣族"
		TypeLines := []string{"Normal"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（効果モンスター）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Ext Ryzeal",
			NameJa:         "エクス・ライゼオル",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       20575,
			TcgID:          34022970,
			Def:            2000,
			Atk:            500,
			Type:           "Effect Monster",
			Level:          4,
			Race:           "Pyro",
			LinkMarkers:    []string{},
			Attribute:      "LIGHT",
			LinkVal:        0,
			TypeLines:      []string{"Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "光"
		RaceNameJa := "炎族"
		TypeLines := []string{"Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（デュアル）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Elemental HERO Neos Alius",
			NameJa:         "E・HERO アナザー・ネオス",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       7195,
			TcgID:          69884162,
			Def:            1300,
			Atk:            1900,
			Type:           "Gemini Monster",
			Level:          4,
			Race:           "Warrior",
			LinkMarkers:    []string{},
			Attribute:      "LIGHT",
			LinkVal:        0,
			TypeLines:      []string{"Gemini", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "光"
		RaceNameJa := "戦士族"
		TypeLines := []string{"Gemini", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（トゥーン）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Toon Cannon Soldier",
			NameJa:         "トゥーン・キャノン・ソルジャー",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       5477,
			TcgID:          79875176,
			Def:            1300,
			Atk:            1400,
			Type:           "Toon Monster",
			Level:          4,
			Race:           "Machine",
			LinkMarkers:    []string{},
			Attribute:      "DARK",
			LinkVal:        0,
			TypeLines:      []string{"Toon", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "闇"
		RaceNameJa := "機械族"
		TypeLines := []string{"Toon", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（効果チューナー）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Al-Lumi'raj",
			NameJa:         "イルミラージュ",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       12145,
			TcgID:          25795273,
			Def:            1000,
			Atk:            1600,
			Type:           "Tuner Monster",
			Level:          3,
			Race:           "Wyrm",
			LinkMarkers:    []string{},
			Attribute:      "WIND",
			LinkVal:        0,
			TypeLines:      []string{"Tuner", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "風"
		RaceNameJa := "幻竜族"
		TypeLines := []string{"Tuner", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（ユニオン）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Fortune Chariot",
			NameJa:         "運命の戦車",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       13996,
			TcgID:          39299733,
			Def:            2000,
			Atk:            1000,
			Type:           "Union Effect Monster",
			Level:          6,
			Race:           "Fairy",
			LinkMarkers:    []string{},
			Attribute:      "WIND",
			LinkVal:        0,
			TypeLines:      []string{"Union", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "風"
		RaceNameJa := "天使族"
		TypeLines := []string{"Union", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（リバース）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Morphing Jar #2",
			NameJa:         "カオスポッド",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       4969,
			TcgID:          79106360,
			Def:            700,
			Atk:            800,
			Type:           "Flip Effect Monster",
			Level:          3,
			Race:           "Rock",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Flip", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "岩石族"
		TypeLines := []string{"Flip", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（リバースチューナー）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Shaddoll Falco",
			NameJa:         "シャドール・ファルコン",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       11232,
			TcgID:          37445295,
			Def:            1400,
			Atk:            600,
			Type:           "Flip Tuner Effect Monster",
			Level:          2,
			Race:           "Spellcaster",
			LinkMarkers:    []string{},
			Attribute:      "DARK",
			LinkVal:        0,
			TypeLines:      []string{"Flip", "Tuner", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "闇"
		RaceNameJa := "魔法使い族"
		TypeLines := []string{"Flip", "Tuner", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（スピリット）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Dark Dust Spirit",
			NameJa:         "砂塵の悪霊",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       5414,
			TcgID:          89111398,
			Def:            1800,
			Atk:            2200,
			Type:           "Spirit Monster",
			Level:          6,
			Race:           "Zombie",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Spirit", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "アンデット族"
		TypeLines := []string{"Spirit", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

}

// 融合・シンクロ・儀式モンスターに関するテスト
func TestNeonMonsterUseCase2(t *testing.T) {
	t.Run("カード情報の挿入&取得テスト（融合バニラ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Karbonala Warrior",
			NameJa:         "カルボナーラ戦士",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       4049,
			TcgID:          54541900,
			Def:            1200,
			Atk:            1500,
			Type:           "Fusion Monster",
			Level:          4,
			Race:           "Warrior",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Fusion"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "戦士族"
		TypeLines := []string{"Fusion"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（融合）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Azamina Mu Rcielago",
			NameJa:         "告死聖徒ルシエラーゴ",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       20507,
			TcgID:          73391962,
			Def:            2400,
			Atk:            2000,
			Type:           "Fusion Monster",
			Level:          6,
			Race:           "Illusion",
			LinkMarkers:    []string{},
			Attribute:      "DARK",
			LinkVal:        0,
			TypeLines:      []string{"Fusion", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "闇"
		RaceNameJa := "幻想魔族"
		TypeLines := []string{"Fusion", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（儀式バニラ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Crab Turtle",
			NameJa:         "クラブ・タートル",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       4716,
			TcgID:          91782219,
			Def:            2500,
			Atk:            2550,
			Type:           "Ritual Monster",
			Level:          8,
			Race:           "Aqua",
			LinkMarkers:    []string{},
			Attribute:      "WATER",
			LinkVal:        0,
			TypeLines:      []string{"Ritual"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "水"
		RaceNameJa := "水族"
		TypeLines := []string{"Ritual"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（儀式）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Water Leviathan @Ignister",
			NameJa:         "ウォーターリヴァイアサン＠イグニスター",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       14843,
			TcgID:          37061511,
			Def:            2000,
			Atk:            2300,
			Type:           "Ritual Monster",
			Level:          7,
			Race:           "Cyberse",
			LinkMarkers:    []string{},
			Attribute:      "WATER",
			LinkVal:        0,
			TypeLines:      []string{"Ritual", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "水"
		RaceNameJa := "サイバース族"
		TypeLines := []string{"Ritual", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（XYZバニラ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Gem-Knight Pearl",
			NameJa:         "ジェムナイト・パール",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       9601,
			TcgID:          71594310,
			Def:            1900,
			Atk:            2600,
			Type:           "XYZ Monster",
			Level:          4,
			Race:           "Rock",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Xyz"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "岩石族"
		TypeLines := []string{"Xyz"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（エクシーズ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Gigantic Spright",
			NameJa:         "ギガンティック・スプライト",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       17449,
			TcgID:          54498517,
			Def:            1600,
			Atk:            1600,
			Type:           "XYZ Monster",
			Level:          2,
			Race:           "Thunder",
			LinkMarkers:    []string{},
			Attribute:      "DARK",
			LinkVal:        0,
			TypeLines:      []string{"Xyz", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "闇"
		RaceNameJa := "雷族"
		TypeLines := []string{"Xyz", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（シンクロバニラ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Scrap Archfiend",
			NameJa:         "スクラップ・デスデーモン",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       9202,
			TcgID:          45815891,
			Def:            1800,
			Atk:            2700,
			Type:           "Synchro Monster",
			Level:          7,
			Race:           "Fiend",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Synchro"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "悪魔族"
		TypeLines := []string{"Synchro"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（シンクロ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Ukiyoe-P.U.N.K. Amazing Dragon",
			NameJa:         "Uk－P.U.N.K.アメイジング・ドラゴン",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       16731,
			TcgID:          44708154,
			Def:            2800,
			Atk:            3000,
			Type:           "Synchro Monster",
			Level:          11,
			Race:           "Sea Serpent",
			LinkMarkers:    []string{},
			Attribute:      "WIND",
			LinkVal:        0,
			TypeLines:      []string{"Synchro", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "風"
		RaceNameJa := "海竜族"
		TypeLines := []string{"Synchro", "Effect"}
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（シンクロチューナー）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Garden Rose Flora",
			NameJa:         "ガーデン・ローズ・フローラ",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       15997,
			TcgID:          76524506,
			Def:            1600,
			Atk:            800,
			Type:           "Synchro Tuner Monster",
			Level:          5,
			Race:           "Plant",
			LinkMarkers:    []string{},
			Attribute:      "LIGHT",
			LinkVal:        0,
			TypeLines:      []string{"Synchro", "Tuner", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "光"
		RaceNameJa := "植物族"
		TypeLines := []string{"Synchro", "Tuner", "Effect"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		assert.NoError(t, err)
		assert.NotNil(t, results)
	})
}

// ペンデュラムモンスターに関するテスト
func TestNeonMonsterUseCase3(t *testing.T) {
	t.Run("カード情報の挿入&取得テスト（ペンデュラムバニラ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Sky Dragoons of Draconia",
			NameJa:         "ドラコニアの翼竜騎兵",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       12031,
			TcgID:          68182934,
			Def:            200,
			Atk:            2200,
			Type:           "Pendulum Normal Monster",
			Level:          5,
			Race:           "Winged Beast",
			LinkMarkers:    []string{},
			Attribute:      "WIND",
			LinkVal:        0,
			TypeLines:      []string{"Normal", "Pendulum"},
			PendulumScale:  7,
			PendulumTextJa: "テキスト２",
			PendulumTextEn: "Text2",
		}

		AttributeNameJa := "風"
		RaceNameJa := "鳥獣族"
		TypeLines := []string{"Normal", "Pendulum"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		//results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		//assert.NoError(t, err)
		//assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（ペンデュラム）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Amorphage Lechery",
			NameJa:         "アモルファージ・ルクス",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       12297,
			TcgID:          70917315,
			Def:            0,
			Atk:            1350,
			Type:           "Pendulum Effect Monster",
			Level:          2,
			Race:           "Dragon",
			LinkMarkers:    []string{},
			Attribute:      "EARTH",
			LinkVal:        0,
			TypeLines:      []string{"Pendulum", "Effect"},
			PendulumScale:  5,
			PendulumTextJa: "テキスト２",
			PendulumTextEn: "text2",
		}

		AttributeNameJa := "地"
		RaceNameJa := "ドラゴン族"
		TypeLines := []string{"Pendulum", "Effect"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		//results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		//assert.NoError(t, err)
		//assert.NotNil(t, results)
	})
	t.Run("カード情報の挿入&取得テスト（融合ペンデュラム）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Supreme King Z-ARC",
			NameJa:         "覇王龍ズァーク",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       12953,
			TcgID:          13331639,
			Def:            4000,
			Atk:            4000,
			Type:           "Fusion Pendulum Effect Monster",
			Level:          12,
			Race:           "Dragon",
			LinkMarkers:    []string{},
			Attribute:      "DARK",
			LinkVal:        0,
			TypeLines:      []string{"Fusion", "Pendulum", "Effect"},
			PendulumScale:  1,
			PendulumTextJa: "テキスト２",
			PendulumTextEn: "text2",
		}

		AttributeNameJa := "闇"
		RaceNameJa := "ドラゴン族"
		TypeLines := []string{"Fusion", "Pendulum", "Effect"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		//results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		//assert.NoError(t, err)
		//assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（Xyzペンデュラム）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "D/D/D Deviser King Deus Machinex",
			NameJa:         "DDD赦俿王デス・マキナ",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       16849,
			TcgID:          46593546,
			Def:            3000,
			Atk:            3000,
			Type:           "XYZ Pendulum Effect Monster",
			Level:          8,
			Race:           "Fiend",
			LinkMarkers:    []string{},
			Attribute:      "DARK",
			LinkVal:        0,
			TypeLines:      []string{"Xyz", "Pendulum", "Effect"},
			PendulumScale:  10,
			PendulumTextJa: "テキスト２",
			PendulumTextEn: "text2",
		}

		AttributeNameJa := "闇"
		RaceNameJa := "悪魔族"
		TypeLines := []string{"Xyz", "Pendulum", "Effect"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		//results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		//assert.NoError(t, err)
		//assert.NotNil(t, results)
	})
}

// リンクモンスターに関するテスト
func TestNeonMonsterUseCase4(t *testing.T) {
	t.Run("カード情報の挿入&取得テスト（リンクバニラ）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "LANphorhynchus",
			NameJa:         "LANフォリンクス",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       13530,
			TcgID:          77637979,
			Def:            0, //
			Atk:            1200,
			Type:           "Link Monster",
			Level:          0,
			Race:           "Cyberse",
			LinkMarkers:    []string{"Bottom-Left", "Bottom-Right"},
			Attribute:      "LIGHT",
			LinkVal:        2,
			TypeLines:      []string{"Link"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "光"
		RaceNameJa := "サイバース族"
		TypeLines := []string{"Link"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		//results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		//assert.NoError(t, err)
		//assert.NotNil(t, results)
	})

	t.Run("カード情報の挿入&取得テスト（リンク）", func(t *testing.T) {
		sampleData := cardrecord.StandardCard{
			NameEn:         "Traptrix Atypus",
			NameJa:         "アティプスの蟲惑魔",
			DescEn:         "text1",
			DescJa:         "テキスト１",
			NeuronID:       18314,
			TcgID:          48183890,
			Def:            0,
			Atk:            1800,
			Type:           "Link Monster",
			Level:          0,
			Race:           "Insect",
			LinkMarkers:    []string{"Left", "Right", "Bottom"},
			Attribute:      "EARTH",
			LinkVal:        3,
			TypeLines:      []string{"Link", "Effect"},
			PendulumScale:  0,
			PendulumTextJa: "",
			PendulumTextEn: "",
		}

		AttributeNameJa := "地"
		RaceNameJa := "昆虫族"
		TypeLines := []string{"Link", "Effect"}
		fmt.Println(sampleData)
		fmt.Println(AttributeNameJa)
		fmt.Println(RaceNameJa)
		fmt.Println(TypeLines)
		//results, err := testMonsterCommon(t, sampleData, AttributeNameJa, RaceNameJa, TypeLines)
		//assert.NoError(t, err)
		//assert.NotNil(t, results)
	})
}
