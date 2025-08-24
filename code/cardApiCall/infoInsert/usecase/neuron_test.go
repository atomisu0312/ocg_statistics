package usecase_test

import (
	"context"
	"strconv"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/htmlget"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestNeuronUseCase(t *testing.T) {
	t.Run("カード情報の取得テスト(アイツ)", func(t *testing.T) {
		neuronID := int64(5642)
		expected := map[htmlget.SelectorKey]string{
			"CardID":     "5642",
			"CardNameEn": "Aitsu",
			"CardNameJa": "アイツ",
			"CardTextJa": "非常に頼りない姿をしているが、実はとてつもない潜在能力を隠し持っているらしい。",
		}

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neuronUseCase := do.MustInvoke[usecase.NeuronUseCase](injector)

		results, err := neuronUseCase.GetCardInfo(context.Background(), neuronID)
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expected["CardID"], strconv.FormatInt(results.CardID, 10))
		assert.Equal(t, expected["CardNameEn"], results.CardNameEn)
		assert.Equal(t, expected["CardNameJa"], results.CardNameJa)
		assert.Equal(t, expected["CardTextJa"], results.CardTextJa)
		assert.Empty(t, results.PendulumTextJa)

	})

	t.Run("カード情報の取得テスト(針淵のヴァリアンツ－アルクトスⅩⅡ)", func(t *testing.T) {
		neuronID := int64(18182)
		expected := map[htmlget.SelectorKey]string{
			"CardID":         "18182",
			"CardNameEn":     "Arktos XII - Chronochasm Vaylantz",
			"CardNameJa":     "針淵のヴァリアンツ－アルクトスⅩⅡ",
			"CardTextJa":     "レベル５以上の「ヴァリアンツ」モンスター×２EXデッキの裏側表示のこのカードは、自分フィールドの上記カードをリリースした場合のみEXデッキから特殊召喚できる。このカード名の①②のモンスター効果はそれぞれ１ターンに１度しか使用できない。①：自分・相手ターンに発動できる。自分のモンスター２体または相手のモンスター２体をメインモンスターゾーンから選び、その２体の位置を入れ替える。②：モンスターゾーンのカードが他のモンスターゾーンに移動した場合に発動できる。フィールドのカード１枚を選んで破壊する。",
			"PendulumTextJa": "このカード名のP効果は１ターンに１度しか使用できない。①：以下の効果から１つを選択して発動できる。●このカードを正面の自分のメインモンスターゾーンに特殊召喚する。●自分のメインモンスターゾーンのモンスター１体を選び、その位置をその隣のモンスターゾーンに移動する。",
		}

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neuronUseCase := do.MustInvoke[usecase.NeuronUseCase](injector)

		results, err := neuronUseCase.GetCardInfo(context.Background(), neuronID)
		assert.NoError(t, err)
		assert.NotNil(t, results)

		assert.Equal(t, expected["CardID"], strconv.FormatInt(results.CardID, 10))
		assert.Equal(t, expected["CardNameEn"], results.CardNameEn)
		assert.Equal(t, expected["CardNameJa"], results.CardNameJa)
		assert.Equal(t, expected["CardTextJa"], results.CardTextJa)
		assert.Equal(t, expected["PendulumTextJa"], results.PendulumTextJa) // 有効
	})

	t.Run("カード情報の取得テスト(古代の機械掌)", func(t *testing.T) {
		neuronID := int64(6844)
		expected := map[htmlget.SelectorKey]string{
			"CardID":     "6844",
			"CardNameEn": "Ancient Gear Fist",
			"CardNameJa": "古代の機械掌",
			"CardTextJa": "「アンティーク・ギア」と名のついたモンスターにのみ装備可能。装備モンスターと戦闘を行ったモンスターを、そのダメージステップ終了時に破壊する。",
		}

		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		neuronUseCase := do.MustInvoke[usecase.NeuronUseCase](injector)

		results, err := neuronUseCase.GetCardInfo(context.Background(), neuronID)
		assert.NoError(t, err)
		assert.NotNil(t, results)

		assert.Equal(t, expected["CardID"], strconv.FormatInt(results.CardID, 10))
		assert.Equal(t, expected["CardNameEn"], results.CardNameEn)
		assert.Equal(t, expected["CardNameJa"], results.CardNameJa)
		assert.Equal(t, expected["CardTextJa"], results.CardTextJa)
		assert.Empty(t, results.PendulumTextJa)
	})
}
