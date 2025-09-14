package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"atomisu.com/ocg-statics/infoInsert/http"
	"github.com/samber/do"
)

type TcgApiCard struct {
	Desc                  string   `json:"desc"`
	Name                  string   `json:"name"`
	ID                    int64    `json:"id"`
	Def                   int64    `json:"def"`
	Atk                   int64    `json:"atk"`
	Type                  string   `json:"type"`
	Level                 int64    `json:"level"`
	Race                  string   `json:"race"`
	LinkMarkers           []string `json:"linkMarkers"`
	Attribute             string   `json:"attribute"`
	LinkVal               int64    `json:"linkVal"`
	TypeLines             []string `json:"typeLines"`
	HumanReadableCardType string   `json:"humanReadableCardType"`
	PendulumText          string   `json:"pendulumText"`
}

// TcgApiResponse is a response from the TcgRest.
type TcgApiResponse struct {
	Data []TcgApiCard `json:"data"`
}

// TcgUseCase is an interface for the TcgUseCase.
type TcgUseCase interface {
	UseCase
	GetCardInfoByEnName(ctx context.Context, name string) (TcgApiCard, error)
}

type tcgUseCaseImpl struct {
	*useCase
}

// NewTcgUseCase is a constructor for TcgUseCase.
func NewTcgUseCase(i *do.Injector) (TcgUseCase, error) {
	return NewUseCase(i, func(u *useCase) TcgUseCase {
		return &tcgUseCaseImpl{u}
	})
}

func (t *tcgUseCaseImpl) emptyFunc() {
}

func (t *tcgUseCaseImpl) GetCardInfoByEnName(ctx context.Context, name string) (TcgApiCard, error) {
	tcgRest := http.NewTCGRest()
	results, err := tcgRest.GetEnInfoByEnName(ctx, name)
	if err != nil {
		return TcgApiCard{}, err
	}

	var response TcgApiResponse
	err = json.Unmarshal(results, &response)
	if err != nil {
		return TcgApiCard{}, err
	}

	if len(response.Data) == 0 {
		return TcgApiCard{}, errors.New("no data found")
	}

	return response.Data[0], nil
}
