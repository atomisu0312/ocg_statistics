package usecase

import (
	"context"

	"atomisu.com/ocg-statics/infoInsert/dto/carddto"

	"github.com/samber/do"
)

// neonUseCaseImpl は、NeonUseCaseの実装です。
type neonUseCaseImpl struct {
	*useCase
}

// NeonUseCase は、NeonUseCaseのインターフェースです。
type NeonUseCase interface {
	UseCase
	InsertTrapCardInfo(ctx context.Context, cardInfo carddto.StandardCard) (int64, error)
	GetTrapCardByID(ctx context.Context, cardID int64) (carddto.TrapCardSelectResult, error)
}

// NewNeonUseCase は、NeonUseCaseのコンストラクタです。
func NewNeonUseCase(i *do.Injector) (NeonUseCase, error) {
	return NewUseCase(i, func(u *useCase) NeonUseCase {
		return &neonUseCaseImpl{u}
	})
}

// emptyFunc は、空の関数です。
func (n *neonUseCaseImpl) emptyFunc() {
}
