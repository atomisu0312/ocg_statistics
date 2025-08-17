package usecase

import (
	"atomisu.com/ocg-statics/infoInsert/config"

	"github.com/samber/do"
)

// useCase は、ユースケースの実装です。
type useCase struct {
	dbConn *config.DbConn
}

// UseCase は、ユースケースのインターフェースです。
type UseCase interface {
	emptyFunc()
}

// NewUseCase は渡されたコンストラクタ関数を使って新規のユースケースを作成す∂る
func NewUseCase[T interface{ UseCase }](i *do.Injector, constructor func(*useCase) T) (T, error) {
	dbConn, err := do.Invoke[*config.DbConn](i)
	if err != nil {
		return constructor(&useCase{}), err
	}

	return constructor(&useCase{
		dbConn: dbConn,
	}), nil
}
