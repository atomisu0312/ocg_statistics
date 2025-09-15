package usecase

import (
	"atomisu.com/ocg-statics/infoInsert/config"

	"github.com/samber/do"
)

// useCase は、ユースケースの実装です。
type UseCaseImpl struct {
	DbConn *config.DbConn
}

// UseCase は、ユースケースのインターフェースです。
type UseCase interface {
}

// NewUseCase は渡されたコンストラクタ関数を使って新規のユースケースを作成す∂る
func NewUseCase[T interface{ UseCase }](i *do.Injector, constructor func(*UseCaseImpl) T) (T, error) {
	dbConn, err := do.Invoke[*config.DbConn](i)
	if err != nil {
		var zero T
		return zero, err
	}

	return constructor(&UseCaseImpl{
		DbConn: dbConn,
	}), nil
}
