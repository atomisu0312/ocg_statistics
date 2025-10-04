package neuron

import (
	"context"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/htmlget"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"atomisu.com/ocg-statics/infoInsert/util"
	"github.com/samber/do"
)

// neuronUseCaseImpl は、NeuronUseCaseの実装です。
type neuronUseCaseImpl struct {
	*usecase.UseCaseImpl
}

// NeuronUseCase は、NeuronUseCaseのインターフェースです。
type NeuronUseCase interface {
	usecase.UseCase
	GetCardInfo(ctx context.Context, cardID int64) (cardrecord.NeuronExtractedData, error)
}

// NewNeuronUseCase は、NeuronUseCaseのコンストラクタです。
func NewNeuronUseCase(i *do.Injector) (NeuronUseCase, error) {
	return usecase.NewUseCase(i, func(u *usecase.UseCaseImpl) NeuronUseCase {
		return &neuronUseCaseImpl{u}
	})
}

// GetCardInfo により、NeuronUseCaseを使ってカードの情報を取得する
func (n *neuronUseCaseImpl) GetCardInfo(ctx context.Context, cardID int64) (cardrecord.NeuronExtractedData, error) {
	htmlGetter := htmlget.NewNeuronHtmlGetter()
	results, err := htmlGetter.VisitSite(ctx, fmt.Sprintf(htmlget.BASE_URL_FORMAT, cardID))

	if err != nil {
		return cardrecord.NeuronExtractedData{}, err
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

	return cardrecord.NeuronExtractedData{
		CardID:         cardID,
		CardNameEn:     results[htmlget.CardNameEn],
		CardNameJa:     cardNameJa,
		CardTextJa:     cardTextJa,
		PendulumTextJa: pendulumTextJa,
	}, nil
}
