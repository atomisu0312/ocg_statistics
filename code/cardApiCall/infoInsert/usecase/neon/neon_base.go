package neon

import (
	"context"

	"atomisu.com/ocg-statics/infoInsert/dto/carddto"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"github.com/samber/do"
)

// neonUseCaseImpl は、NeonUseCaseの実装です。
type NeonUseCaseImpl struct {
	*usecase.UseCaseImpl
}

// NeonUseCase は、NeonUseCaseのインターフェースです。
type NeonUseCase interface {
	usecase.UseCase
	InsertTrapCardInfo(ctx context.Context, cardInfo carddto.StandardCard) (int64, error)
	GetTrapCardByID(ctx context.Context, cardID int64) (carddto.TrapCardSelectResult, error)
	InsertSpellCardInfo(ctx context.Context, cardInfo carddto.StandardCard) (int64, error)
	GetSpellCardByID(ctx context.Context, cardID int64) (carddto.SpellCardSelectResult, error)
}

// NewNeonUseCase は、NeonUseCaseのコンストラクタです。
func NewNeonUseCase(i *do.Injector) (NeonUseCase, error) {
	return usecase.NewUseCase(i, func(u *usecase.UseCaseImpl) NeonUseCase {
		return &NeonUseCaseImpl{u}
	})
}
