package htmlget

import (
	"context"

	"github.com/gocolly/colly"
)

// SelectorKey はセレクタのキーを表す型
type SelectorKey string

const (
	CardNameEn      SelectorKey = "CardNameEn"
	CardNameJa      SelectorKey = "CardNameJa"
	CardTextJa1     SelectorKey = "CardTextJa1"
	CardTextJa2     SelectorKey = "CardTextJa2"
	PendulumTextJa  SelectorKey = "PendulumTextJa"
	BASE_URL_FORMAT string      = "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=%d&request_locale=ja"
)

// SelectorMap はセレクタのマッピングを表す型
type SelectorMap map[SelectorKey]ElementInfo

var selectorMap = SelectorMap{
	CardNameEn: {
		Selector: "#cardname > h1 > span:nth-child(2)",
	},
	CardNameJa: {
		Selector: "#cardname > h1",
	},
	CardTextJa1: {
		Selector: "#CardSet > div.top > div:nth-child(4) > div",
	},
	CardTextJa2: {
		Selector: "#CardSet > div.top > div:nth-child(5) > div",
	},
	PendulumTextJa: {
		Selector: "#CardSet > div.top > div.CardText.pen > div.frame.pen_effect > div",
	},
}

type neuronHtmlGetterImpl struct {
	*htmlGetter
}

type NeuronHtmlGetter interface {
	HTMLGetter
	VisitSite(ctx context.Context, url string) (map[SelectorKey]string, error)
}

// コンストラクタ
func NewNeuronHtmlGetter() NeuronHtmlGetter {
	return NewHtmlGetter(func(h *htmlGetter) NeuronHtmlGetter {
		return &neuronHtmlGetterImpl{htmlGetter: h}
	}, colly.NewCollector(
		colly.AllowedDomains("www.db.yugioh-card.com"),
		colly.DetectCharset(),
	), &selectorMap)
}

func (h *neuronHtmlGetterImpl) VisitSite(ctx context.Context, url string) (map[SelectorKey]string, error) {
	results, err := h.htmlGetter.Visit(ctx, url)
	if err != nil {
		return map[SelectorKey]string{}, err
	}

	return results, nil
}
