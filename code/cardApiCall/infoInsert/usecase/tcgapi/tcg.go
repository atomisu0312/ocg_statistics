package tcgapi

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/http"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"github.com/samber/do"
)

// TcgApiResponse is a response from the TcgRest.
type TcgApiResponse struct {
	Data []cardrecord.TcgApiCard `json:"data"`
}

// TcgUseCase is an interface for the TcgUseCase.
type TcgUseCase interface {
	usecase.UseCase
	GetCardInfoByEnName(ctx context.Context, name string) (cardrecord.TcgApiCard, error)
}

type tcgUseCaseImpl struct {
	*usecase.UseCaseImpl
}

// NewTcgUseCase is a constructor for TcgUseCase.
func NewTcgUseCase(i *do.Injector) (TcgUseCase, error) {
	return usecase.NewUseCase(i, func(u *usecase.UseCaseImpl) TcgUseCase {
		return &tcgUseCaseImpl{u}
	})
}

var OmitChars = []string{"＃", "#", "<", ">", "&"}

func (t *tcgUseCaseImpl) GetCardInfoByEnName(ctx context.Context, name string) (cardrecord.TcgApiCard, error) {
	tcgRest := http.NewTCGRest()

	for _, omitChar := range OmitChars {
		name = strings.ReplaceAll(name, omitChar, "")
	}

	results, err := tcgRest.GetEnInfoByEnName(ctx, name)
	if err != nil {
		return cardrecord.TcgApiCard{}, err
	}

	var response TcgApiResponse
	err = json.Unmarshal(results, &response)
	if err != nil {
		return cardrecord.TcgApiCard{}, err
	}

	if len(response.Data) == 0 {
		return cardrecord.TcgApiCard{}, errors.New("no data found")
	}

	tcgCard := response.Data[0]
	return tcgCard, nil
}
