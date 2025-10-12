package master_test

import (
	"context"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/usecase/master"
	"atomisu.com/ocg-statics/infoInsert/usecase/neon"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestInsertCardInfo(t *testing.T) {
	testCases := []struct {
		card cardrecord.StandardCard
	}{
		{
			card: cardrecord.StandardCard{
				NameEn:    "Mokey Mokey",
				NameJa:    "もけもけ",
				DescEn:    "An outcast angel.",
				DescJa:    "天使のはみだし者",
				NeuronID:  6018,
				TcgID:     27288416,
				Def:       100,
				Atk:       300,
				Type:      "Normal Monster",
				Level:     1,
				Race:      "Fairy",
				Attribute: "LIGHT",
				TypeLines: []string{"Normal"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Melffy Rabby",
				NameJa:    "メルフィー・ラビィ",
				NeuronID:  15250,
				TcgID:     20129614,
				Def:       2100,
				Atk:       0,
				Type:      "Normal Monster",
				Level:     2,
				Race:      "Beast",
				Attribute: "EARTH",
				TypeLines: []string{"Normal"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Ext Ryzeal",
				NameJa:    "エクス・ライゼオル",
				NeuronID:  20575,
				TcgID:     34022970,
				Def:       2000,
				Atk:       500,
				Type:      "Effect Monster",
				Level:     4,
				Race:      "Pyro",
				Attribute: "LIGHT",
				TypeLines: []string{"Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Elemental HERO Neos Alius",
				NameJa:    "E・HERO アナザー・ネオス",
				NeuronID:  7195,
				TcgID:     69884162,
				Def:       1300,
				Atk:       1900,
				Type:      "Gemini Monster",
				Level:     4,
				Race:      "Warrior",
				Attribute: "LIGHT",
				TypeLines: []string{"Gemini", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Toon Cannon Soldier",
				NameJa:    "トゥーン・キャノン・ソルジャー",
				NeuronID:  5477,
				TcgID:     79875176,
				Def:       1300,
				Atk:       1400,
				Type:      "Toon Monster",
				Level:     4,
				Race:      "Machine",
				Attribute: "DARK",
				TypeLines: []string{"Toon", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Al-Lumi'raj",
				NameJa:    "イルミラージュ",
				NeuronID:  12145,
				TcgID:     25795273,
				Def:       1000,
				Atk:       1600,
				Type:      "Tuner Monster",
				Level:     3,
				Race:      "Wyrm",
				Attribute: "WIND",
				TypeLines: []string{"Tuner", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Fortune Chariot",
				NameJa:    "運命の戦車",
				NeuronID:  13996,
				TcgID:     39299733,
				Def:       2000,
				Atk:       1000,
				Type:      "Union Effect Monster",
				Level:     6,
				Race:      "Fairy",
				Attribute: "WIND",
				TypeLines: []string{"Union", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Morphing Jar #2",
				NameJa:    "カオスポッド",
				NeuronID:  4969,
				TcgID:     79106360,
				Def:       700,
				Atk:       800,
				Type:      "Flip Effect Monster",
				Level:     3,
				Race:      "Rock",
				Attribute: "EARTH",
				TypeLines: []string{"Flip", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Shaddoll Falco",
				NameJa:    "シャドール・ファルコン",
				NeuronID:  11232,
				TcgID:     37445295,
				Def:       1400,
				Atk:       600,
				Type:      "Flip Tuner Effect Monster",
				Level:     2,
				Race:      "Spellcaster",
				Attribute: "DARK",
				TypeLines: []string{"Flip", "Tuner", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Dark Dust Spirit",
				NameJa:    "砂塵の悪霊",
				NeuronID:  5414,
				TcgID:     89111398,
				Def:       1800,
				Atk:       2200,
				Type:      "Spirit Monster",
				Level:     6,
				Race:      "Zombie",
				Attribute: "EARTH",
				TypeLines: []string{"Spirit", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Karbonala Warrior",
				NameJa:    "カルボナーラ戦士",
				NeuronID:  4049,
				TcgID:     54541900,
				Def:       1200,
				Atk:       1500,
				Type:      "Fusion Monster",
				Level:     4,
				Race:      "Warrior",
				Attribute: "EARTH",
				TypeLines: []string{"Fusion"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Azamina Mu Rcielago",
				NameJa:    "告死聖徒ルシエラーゴ",
				NeuronID:  20507,
				TcgID:     73391962,
				Def:       2400,
				Atk:       2000,
				Type:      "Fusion Monster",
				Level:     6,
				Race:      "Illusion",
				Attribute: "DARK",
				TypeLines: []string{"Fusion", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Crab Turtle",
				NameJa:    "クラブ・タートル",
				NeuronID:  4716,
				TcgID:     91782219,
				Def:       2500,
				Atk:       2550,
				Type:      "Ritual Monster",
				Level:     8,
				Race:      "Aqua",
				Attribute: "WATER",
				TypeLines: []string{"Ritual"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Water Leviathan @Ignister",
				NameJa:    "ウォーターリヴァイアサン＠イグニスター",
				NeuronID:  14843,
				TcgID:     37061511,
				Def:       2000,
				Atk:       2300,
				Type:      "Ritual Monster",
				Level:     7,
				Race:      "Cyberse",
				Attribute: "WATER",
				TypeLines: []string{"Ritual", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Gem-Knight Pearl",
				NameJa:    "ジェムナイト・パール",
				NeuronID:  9601,
				TcgID:     71594310,
				Def:       1900,
				Atk:       2600,
				Type:      "XYZ Monster",
				Level:     4,
				Race:      "Rock",
				Attribute: "EARTH",
				TypeLines: []string{"Xyz"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Gigantic Spright",
				NameJa:    "ギガンティック・スプライト",
				NeuronID:  17449,
				TcgID:     54498517,
				Def:       1600,
				Atk:       1600,
				Type:      "XYZ Monster",
				Level:     2,
				Race:      "Thunder",
				Attribute: "DARK",
				TypeLines: []string{"Xyz", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Scrap Archfiend",
				NameJa:    "スクラップ・デスデーモン",
				NeuronID:  9202,
				TcgID:     45815891,
				Def:       1800,
				Atk:       2700,
				Type:      "Synchro Monster",
				Level:     7,
				Race:      "Fiend",
				Attribute: "EARTH",
				TypeLines: []string{"Synchro"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Ukiyoe-P.U.N.K. Amazing Dragon",
				NameJa:    "Uk－P.U.N.K.アメイジング・ドラゴン",
				NeuronID:  16731,
				TcgID:     44708154,
				Def:       2800,
				Atk:       3000,
				Type:      "Synchro Monster",
				Level:     11,
				Race:      "Sea Serpent",
				Attribute: "WIND",
				TypeLines: []string{"Synchro", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:    "Garden Rose Flora",
				NameJa:    "ガーデン・ローズ・フローラ",
				NeuronID:  15997,
				TcgID:     76524506,
				Def:       1600,
				Atk:       800,
				Type:      "Synchro Tuner Monster",
				Level:     5,
				Race:      "Plant",
				Attribute: "LIGHT",
				TypeLines: []string{"Synchro", "Tuner", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:        "Sky Dragoons of Draconia",
				NameJa:        "ドラコニアの翼竜騎兵",
				NeuronID:      12031,
				TcgID:         68182934,
				Def:           200,
				Atk:           2200,
				Type:          "Pendulum Normal Monster",
				Level:         5,
				Race:          "Winged Beast",
				Attribute:     "WIND",
				TypeLines:     []string{"Normal", "Pendulum"},
				PendulumScale: 7,
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:        "Amorphage Lechery",
				NameJa:        "アモルファージ・ルクス",
				NeuronID:      12297,
				TcgID:         70917315,
				Def:           0,
				Atk:           1350,
				Type:          "Pendulum Effect Monster",
				Level:         2,
				Race:          "Dragon",
				Attribute:     "EARTH",
				TypeLines:     []string{"Pendulum", "Effect"},
				PendulumScale: 5,
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:        "Supreme King Z-ARC",
				NameJa:        "覇王龍ズァーク",
				NeuronID:      12953,
				TcgID:         13331639,
				Def:           4000,
				Atk:           4000,
				Type:          "Fusion Pendulum Effect Monster",
				Level:         12,
				Race:          "Dragon",
				Attribute:     "DARK",
				TypeLines:     []string{"Fusion", "Pendulum", "Effect"},
				PendulumScale: 1,
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:        "D/D/D Deviser King Deus Machinex",
				NameJa:        "DDD赦俿王デス・マキナ",
				NeuronID:      16849,
				TcgID:         46593546,
				Def:           3000,
				Atk:           3000,
				Type:          "XYZ Pendulum Effect Monster",
				Level:         8,
				Race:          "Fiend",
				Attribute:     "DARK",
				TypeLines:     []string{"Xyz", "Pendulum", "Effect"},
				PendulumScale: 10,
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "LANphorhynchus",
				NameJa:      "LANフォリンクス",
				NeuronID:    13530,
				TcgID:       77637979,
				Atk:         1200,
				Type:        "Link Monster",
				Race:        "Cyberse",
				LinkMarkers: []string{"Bottom-Left", "Bottom-Right"},
				Attribute:   "LIGHT",
				LinkVal:     2,
				TypeLines:   []string{"Link"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Traptrix Atypus",
				NameJa:      "アティプスの蟲惑魔",
				NeuronID:    18314,
				TcgID:       48183890,
				Atk:         1800,
				Type:        "Link Monster",
				Race:        "Insect",
				LinkMarkers: []string{"Left", "Right", "Bottom"},
				Attribute:   "EARTH",
				LinkVal:     3,
				TypeLines:   []string{"Link", "Effect"},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Infinite Impermanence",
				NameJa:      "無限泡影",
				NeuronID:    13631,
				TcgID:       10045474,
				Atk:         0,
				Type:        "Trap Card",
				Race:        "Normal",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Macro Cosmos",
				NameJa:      "マクロコスモス",
				NeuronID:    6682,
				TcgID:       30241314,
				Atk:         0,
				Type:        "Trap Card",
				Race:        "Continuous",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Solemn Judgment",
				NameJa:      "神の宣告",
				NeuronID:    4861,
				TcgID:       41420027,
				Atk:         0,
				Type:        "Trap Card",
				Race:        "Counter",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Kashtiratheosis",
				NameJa:      "六世壊他化自在天",
				NeuronID:    18203,
				TcgID:       34447918,
				Atk:         0,
				Type:        "Spell Card",
				Race:        "Normal",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Kashtira Birth",
				NameJa:      "クシャトリラ・バース",
				NeuronID:    17815,
				TcgID:       69540484,
				Atk:         0,
				Type:        "Spell Card",
				Race:        "Continuous",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Majesty with Eyes of Blue",
				NameJa:      "青き眼の威光",
				NeuronID:    12587,
				TcgID:       2783661,
				Atk:         0,
				Type:        "Spell Card",
				Race:        "Quick-Play",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Primeval Planet Perlereino",
				NameJa:      "壱世壊＝ペルレイノ",
				NeuronID:    17462,
				TcgID:       77103950,
				Atk:         0,
				Type:        "Spell Card",
				Race:        "Field",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Lightning Blade",
				NameJa:      "稲妻の剣",
				NeuronID:    5110,
				TcgID:       55226821,
				Atk:         0,
				Type:        "Spell Card",
				Race:        "Equip",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
		{
			card: cardrecord.StandardCard{
				NameEn:      "Shinobird's Calling",
				NameJa:      "霊魂の降神",
				NeuronID:    12795,
				TcgID:       73055622,
				Atk:         0,
				Type:        "Spell Card",
				Race:        "Ritual",
				LinkMarkers: []string{},
				Attribute:   "",
				LinkVal:     0,
				TypeLines:   []string{},
			},
		},
	}
	config.BeforeEachForUnitTest()
	defer config.AfterEachForUnitTest()
	injector := app.SetupDIContainer()
	do.Override(injector, config.TestDbConnection)
	masterUseCase := do.MustInvoke[master.MasterUseCase](injector)
	neonUseCase := do.MustInvoke[neon.NeonUseCase](injector)

	for _, tc := range testCases {
		t.Run("正常系: "+tc.card.NameJa, func(t *testing.T) {
			t.Parallel()
			cardID, err := masterUseCase.InsertCardInfo(context.Background(), tc.card.NeuronID)
			assert.NoError(t, err)
			assert.NotEqual(t, int64(0), cardID)

			switch tc.card.Type {
			case "Spell Card":
				results, _ := neonUseCase.GetSpellCardByID(context.Background(), cardID)
				assert.NotNil(t, results)
				assert.Equal(t, tc.card.NameEn, results.NameEn)
				assert.Equal(t, tc.card.NameJa, results.NameJa)
				assert.Equal(t, tc.card.NeuronID, results.NeuronID)
				assert.Equal(t, tc.card.TcgID, results.OcgApiID)
				switch tc.card.Race {
				case "Normal":
					assert.Equal(t, "通常", results.SpellTypeNameJa)
					assert.Equal(t, "Normal", results.SpellTypeNameEn)
				case "Continuous":
					assert.Equal(t, "永続", results.SpellTypeNameJa)
					assert.Equal(t, "Continuous", results.SpellTypeNameEn)
				case "Equip":
					assert.Equal(t, "装備", results.SpellTypeNameJa)
					assert.Equal(t, "Equip", results.SpellTypeNameEn)
				case "Field":
					assert.Equal(t, "フィールド", results.SpellTypeNameJa)
					assert.Equal(t, "Field", results.SpellTypeNameEn)
				case "Quick-Play":
					assert.Equal(t, "速攻", results.SpellTypeNameJa)
					assert.Equal(t, "Quick-Play", results.SpellTypeNameEn)
				case "Ritual":
					assert.Equal(t, "儀式", results.SpellTypeNameJa)
					assert.Equal(t, "Ritual", results.SpellTypeNameEn)
				}
			case "Trap Card":
				results, _ := neonUseCase.GetTrapCardByID(context.Background(), cardID)
				assert.NotNil(t, results)
				assert.Equal(t, tc.card.NameEn, results.NameEn)
				assert.Equal(t, tc.card.NameJa, results.NameJa)
				assert.Equal(t, tc.card.NeuronID, results.NeuronID)
				assert.Equal(t, tc.card.TcgID, results.OcgApiID)
				switch tc.card.Race {
				case "Normal":
					assert.Equal(t, "通常", results.TrapTypeNameJa)
					assert.Equal(t, "Normal", results.TrapTypeNameEn)
				case "Continuous":
					assert.Equal(t, "永続", results.TrapTypeNameJa)
					assert.Equal(t, "Continuous", results.TrapTypeNameEn)
				case "Counter":
					assert.Equal(t, "カウンター", results.TrapTypeNameJa)
					assert.Equal(t, "Counter", results.TrapTypeNameEn)
				}
			default:
				monster, err := neonUseCase.GetMonsterCardExtendedByID(context.Background(), cardID)
				assert.NotNil(t, monster)
				assert.Equal(t, tc.card.NameEn, monster.NameEn)
				assert.Equal(t, tc.card.NameJa, monster.NameJa)
				assert.Equal(t, tc.card.NeuronID, monster.NeuronID)
				assert.Equal(t, tc.card.TcgID, monster.OcgApiID)
				typeLines, err := neonUseCase.GetMonsterTypeLinesEnByCardID(context.Background(), cardID)
				assert.NoError(t, err)
				assert.NotNil(t, typeLines)
				assert.ElementsMatch(t, tc.card.TypeLines, typeLines)

			}
		})
	}

	t.Run("異常系_01: 存在しないカードID", func(t *testing.T) {
		cardID, err := masterUseCase.InsertCardInfo(context.Background(), 1999999)
		assert.Error(t, err)
		assert.Equal(t, int64(0), cardID)
	})

	t.Run("異常系_02: 英字情報しか取得できないカード", func(t *testing.T) {
		cardID, err := masterUseCase.InsertCardInfo(context.Background(), 21967)
		assert.Error(t, err)
		assert.Equal(t, int64(0), cardID)
	})
}
