package tcgapi_test

import (
	"context"
	"strconv"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/usecase/tcgapi"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestTcgUseCase(t *testing.T) {
	t.Run("カード情報の取得テスト(デコード・トーカー)", func(t *testing.T) {
		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		expected := map[string]string{
			"id":   "1861629",
			"name": "Decode Talker",
			"desc": "2+ Effect Monsters\r\nGains 500 ATK for each monster it points to. When your opponent activates a card or effect that targets a card(s) you control (Quick Effect): You can Tribute 1 monster this card points to; negate the activation, and if you do, destroy that card.",
		}

		tcgUseCase := do.MustInvoke[tcgapi.TcgUseCase](injector)

		results, err := tcgUseCase.GetCardInfoByEnName(context.Background(), "Decode Talker")
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expected["id"], strconv.FormatInt(results.ID, 10))
		assert.Equal(t, expected["name"], results.Name)
		assert.Equal(t, expected["desc"], results.Desc)
	})

	t.Run("カード情報の取得テスト(デコード・トーカー?)", func(t *testing.T) {
		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		tcgUseCase := do.MustInvoke[tcgapi.TcgUseCase](injector)

		_, err := tcgUseCase.GetCardInfoByEnName(context.Background(), "Decode Tolker")
		assert.Error(t, err)
	})

	t.Run("カード情報の取得テスト(砦を守る翼竜)_特殊文字が入る場合", func(t *testing.T) {
		injector := app.SetupDIContainer()
		do.Override(injector, config.TestDbConnection)

		expected := map[string]string{
			"id":   "87796900",
			"name": "Winged Dragon, Guardian of the Fortress #1",
			"desc": "A dragon commonly found guarding mountain fortresses. Its signature attack is a sweeping dive from out of the blue.",
		}

		tcgUseCase := do.MustInvoke[tcgapi.TcgUseCase](injector)

		results, err := tcgUseCase.GetCardInfoByEnName(context.Background(), "Winged Dragon, Guardian of the Fortress #1")
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expected["id"], strconv.FormatInt(results.ID, 10))
		assert.Equal(t, expected["name"], results.Name)
		assert.Equal(t, expected["desc"], results.Desc)
	})
}
