package http

import (
	"context"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type rest interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type restImpl struct {
	logger *zap.Logger
}

func NewRest[T rest](constructor func(*restImpl) T, logger *zap.Logger) T {
	return constructor(&restImpl{logger: logger})
}

func (h *restImpl) Get(ctx context.Context, url string) ([]byte, error) {
	h.logger.Info("HTTP GET リクエスト開始", zap.String("url", url))

	resp, err := http.Get(url)
	if err != nil {
		h.logger.Error("HTTP GET リクエストエラー", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	h.logger.Info("HTTP レスポンス受信", zap.String("url", url), zap.String("status", resp.Status))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		h.logger.Error("レスポンスボディ読み取りエラー", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	h.logger.Info("レスポンスボディ読み取り完了", zap.String("url", url), zap.Int("body_length", len(body)))

	return body, nil
}
