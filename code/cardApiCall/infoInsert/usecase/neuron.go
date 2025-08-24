package usecase

import (
	"context"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/htmlget"
	"atomisu.com/ocg-statics/infoInsert/util"

	"github.com/samber/do"
)

// neuronUseCaseImpl は、NeuronUseCaseの実装です。
type neuronUseCaseImpl struct {
	*useCase
}

// NeuronUseCase は、NeuronUseCaseのインターフェースです。
type NeuronUseCase interface {
	UseCase
	GetCardInfo(ctx context.Context, cardID int64) (NeuronExtractedData, error)
}

// NewNeuronUseCase は、NeuronUseCaseのコンストラクタです。
func NewNeuronUseCase(i *do.Injector) (NeuronUseCase, error) {
	return NewUseCase(i, func(u *useCase) NeuronUseCase {
		return &neuronUseCaseImpl{u}
	})
}

// emptyFunc は、空の関数です。
func (n *neuronUseCaseImpl) emptyFunc() {
}

// NeuronExtractedData は、NeuronUseCaseの抽出データです。
type NeuronExtractedData struct {
	CardID         int64
	CardNameEn     string
	CardNameJa     string
	CardTextJa     string
	PendulumTextJa string
}

// GetCardInfo により、NeuronUseCaseを使ってカードの情報を取得する
func (n *neuronUseCaseImpl) GetCardInfo(ctx context.Context, cardID int64) (NeuronExtractedData, error) {
	htmlGetter := htmlget.NewNeuronHtmlGetter()
	results, err := htmlGetter.VisitSite(ctx, fmt.Sprintf(htmlget.BASE_URL_FORMAT, cardID))

	if err != nil {
		return NeuronExtractedData{}, err
	}

	// テキストを適切に処理
	cardNameJaParts := util.SplitByNewlinesAndTabs(results[htmlget.CardNameJa])
	cardTextJaParts := util.SplitByNewlinesAndTabs(results[htmlget.CardTextJa1])
	cardTextJa2Parts := util.SplitByNewlinesAndTabs(results[htmlget.CardTextJa2])
	pendulumTextJaParts := util.SplitByNewlinesAndTabs(results[htmlget.PendulumTextJa])

	var cardNameJa, cardTextJa, pendulumTextJa string
	if len(cardNameJaParts) > 0 {
		cardNameJa = cardNameJaParts[1]
	}

	if len(pendulumTextJaParts) > 0 {
		pendulumTextJa = pendulumTextJaParts[0]
	}

	if pendulumTextJa == "" {
		cardTextJa = cardTextJaParts[1]
	} else {
		cardTextJa = cardTextJa2Parts[1]
	}

	return NeuronExtractedData{
		CardID:         cardID,
		CardNameEn:     results[htmlget.CardNameEn],
		CardNameJa:     cardNameJa,
		CardTextJa:     cardTextJa,
		PendulumTextJa: pendulumTextJa,
	}, nil
}
