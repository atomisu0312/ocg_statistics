package usecase

import (
	"database/sql"

	"atomisu.com/ocg-statics/infoInsert/config"

	"github.com/samber/do"
)

// useCase は、ユースケースの実装です。
type UseCaseImpl struct {
	dbConn *config.DbConn
}

func (u *UseCaseImpl) ProduceConnDB() *sql.DB {
	if u.dbConn == nil {
		panic("dbConn is nil")
	}
	return u.dbConn.DB
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
		dbConn: dbConn,
	}), nil
}
