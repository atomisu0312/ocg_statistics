package http

import (
	"context"
	"fmt"
	"net/url"

	"go.uber.org/zap"
)

type tcgRestImpl struct {
	*restImpl
}

// TcgRest is an interface for the TcgRest.
type TcgRest interface {
	rest
	Get(ctx context.Context, url string) ([]byte, error)
	GetEnInfoByEnName(ctx context.Context, name string) ([]byte, error)
	GetJaInfoByJaName(ctx context.Context, name string) ([]byte, error)
}

// NewTCGRest is a constructor for TcgRest.
func NewTCGRest() TcgRest {
	return NewRest(func(r *restImpl) TcgRest {
		return &tcgRestImpl{restImpl: r}
	}, zap.NewExample())
}

// GetEnInfoByName is a method to get the English information of a card by name.
func (r *tcgRestImpl) GetEnInfoByEnName(ctx context.Context, name string) ([]byte, error) {
	// URLエンコーディングを使用
	encodedName := url.QueryEscape(name)
	url := fmt.Sprintf("https://db.ygoprodeck.com/api/v7/cardinfo.php?name=%s", encodedName)

	return r.Get(ctx, url)
}

// GetJaInfoByName is a method to get the Japanese information of a card by name.
func (r *tcgRestImpl) GetJaInfoByJaName(ctx context.Context, name string) ([]byte, error) {
	// URLエンコーディングを使用
	encodedName := url.QueryEscape(name)
	url := fmt.Sprintf("https://db.ygoprodeck.com/api/v7/cardinfo.php?name=%s&language=ja", encodedName)
	return r.Get(ctx, url)
}
