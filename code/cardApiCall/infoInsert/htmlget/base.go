package htmlget

import (
	"context"

	"go.uber.org/zap"

	"github.com/gocolly/colly"
)

// HTMLGetter は、HTMLGetterの基本インターフェースです。
type HTMLGetter interface {
	Visit(ctx context.Context, url string) (map[SelectorKey]string, error)
}

// htmlGetter は、HTMLGetterの基本構造体です。
type htmlGetter struct {
	logger      *zap.Logger
	collector   *colly.Collector
	selectorMap *SelectorMap
	results     map[SelectorKey]string
}

// ElementInfo は要素の情報を表す構造体
type ElementInfo struct {
	Selector string
}

func NewHtmlGetter[T HTMLGetter](
	constructor func(*htmlGetter) T,
	collector *colly.Collector,
	selectorMap *SelectorMap,
) T {
	logger, _ := zap.NewDevelopment()
	htmlGetter := &htmlGetter{logger: logger, collector: collector, selectorMap: selectorMap, results: make(map[SelectorKey]string)}

	for key, elementInfo := range *selectorMap {
		collector.OnHTML(elementInfo.Selector, func(e *colly.HTMLElement) {
			htmlGetter.results[key] = e.Text
		})
	}

	return constructor(htmlGetter)
}

func (h *htmlGetter) Visit(ctx context.Context, url string) (map[SelectorKey]string, error) {
	err := h.collector.Visit(url)
	if err != nil {
		h.logger.Error("Failed to visit URL", zap.String("url", url), zap.Error(err))
	}
	return h.results, nil
}
