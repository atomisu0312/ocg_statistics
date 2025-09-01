package http_test

import (
	"context"
	"io"
	httplib "net/http"
	"testing"

	"atomisu.com/ocg-statics/infoInsert/http"
	"github.com/stretchr/testify/assert"
)

func TestTCGRest(t *testing.T) {
	t.Run("GetEnInfoByName", func(t *testing.T) {
		// 英語名でカードを検索
		body, err := http.NewTCGRest().GetEnInfoByEnName(context.Background(), "Decode Talker")

		// デバッグ用のログ出力
		t.Logf("カスタム実装レスポンス長: %d", len(body))
		t.Logf("カスタム実装エラー: %v", err)

		assert.NoError(t, err)
		assert.NotEmpty(t, string(body), "レスポンスが空であってはならない")
	})

	t.Run("GetJaInfoByName", func(t *testing.T) {
		// 日本語名でカードを検索
		body, err := http.NewTCGRest().GetJaInfoByJaName(context.Background(), "デコード・トーカー")

		// デバッグ用のログ出力
		t.Logf("カスタム実装レスポンス長: %d", len(body))
		t.Logf("カスタム実装エラー: %v", err)

		assert.NoError(t, err)
		assert.NotEmpty(t, string(body), "レスポンスが空であってはならない")
	})

	t.Run("直接HTTPリクエスト", func(t *testing.T) {
		// 直接HTTPリクエストを送信してテスト
		resp, err := httplib.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php?name=Decode%20Talker")
		if err != nil {
			t.Fatalf("HTTPリクエストエラー: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("ボディ読み取りエラー: %v", err)
		}

		t.Logf("直接HTTPレスポンス長: %d", len(body))
		t.Logf("HTTPステータス: %s", resp.Status)

		assert.NotEmpty(t, string(body), "直接HTTPリクエストのレスポンスが空であってはならない")
	})
}
