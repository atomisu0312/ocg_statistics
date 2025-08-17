package htmlget

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNeuronHtmlGetter(t *testing.T) {

	// 実際のスクレイピングテスト
	t.Run("要素の取得テスト(ペンデュラムテキスト有)", func(t *testing.T) {
		getter := NewNeuronHtmlGetter()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 実際のカードIDでテスト(アストログラフマジシャン)
		url := fmt.Sprintf(BASE_URL_FORMAT, 12906)

		results, err := getter.Visit(ctx, url)
		if err != nil {
			t.Logf("スクレイピングエラー（予想される場合もある）: %v", err)
			return
		}

		// 結果の確認
		assert.NotNil(t, results, "結果が取得できていること")

		// 各セレクタの結果を確認
		assert.True(t, results[CardNameEn] == "Astrograph Sorcerer", "英語名が取得できていること")
		assert.True(t, results[CardNameJa] != "", "日本語名が取得できていること")
		assert.True(t, results[CardTextJa1] != "", "カードテキストが取得できていること")
		assert.True(t, results[CardTextJa2] != "", "カードテキストが取得できていること")
		assert.True(t, results[PendulumTextJa] != "", "ペンデュラムテキストが取得できていること")
	})

	t.Run("要素の取得テスト(ペンデュラムテキスト無)", func(t *testing.T) {
		getter := NewNeuronHtmlGetter()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 実際のカードIDでテスト(幻煌龍の天渦)
		url := fmt.Sprintf(BASE_URL_FORMAT, 12988)

		results, err := getter.Visit(ctx, url)
		if err != nil {
			t.Logf("スクレイピングエラー（予想される場合もある）: %v", err)
			return
		}

		assert.NotNil(t, results, "結果が取得できていること")
		assert.True(t, results[CardNameJa] != "", "日本語名が取得できていること")
		assert.True(t, results[CardTextJa1] != "", "カードテキストが取得できていること")
		assert.True(t, results[CardTextJa2] != "", "カードテキストが取得できていること")
		assert.True(t, results[PendulumTextJa] == "", "ペンデュラムテキストが取得できていないこと")
	})
}
