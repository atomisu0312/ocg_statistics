package neon

import (
	"context"

	"atomisu.com/ocg-statics/infoInsert/dto/cardrecord"
	"atomisu.com/ocg-statics/infoInsert/usecase"
	"github.com/samber/do"
)

// neonUseCaseImpl は、NeonUseCaseの実装です。
type neonUseCaseImpl struct {
	*usecase.UseCaseImpl
}

// NeonUseCase は、NeonUseCaseのインターフェースです。
type NeonUseCase interface {
	usecase.UseCase
	InsertTrapCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error)
	GetTrapCardByID(ctx context.Context, cardID int64) (cardrecord.TrapCardSelectResult, error)
	InsertSpellCardInfo(ctx context.Context, cardInfo cardrecord.StandardCard) (int64, error)
	GetSpellCardByID(ctx context.Context, cardID int64) (cardrecord.SpellCardSelectResult, error)
}

// NewNeonUseCase は、NeonUseCaseのコンストラクタです。
func NewNeonUseCase(i *do.Injector) (NeonUseCase, error) {
	return usecase.NewUseCase(i, func(u *usecase.UseCaseImpl) NeonUseCase {
		return &neonUseCaseImpl{u}
	})
}
