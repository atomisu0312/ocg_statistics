package usecase

import (
	"atomisu.com/ocg-statics/infoInsert/config"

	"github.com/samber/do"
)

// ハンドラから直接呼ばれるのがユースケース
type useCase struct {
	dbConn *config.DbConn
}

type UseCase interface {
	emptyFunc()
}

// NewUseCase は新しい UseCase インスタンスを作成します
func NewUseCase[T interface{ UseCase }](i *do.Injector, constructor func(*useCase) T) (T, error) {
	dbConn := do.MustInvoke[*config.DbConn](i)

	return constructor(&useCase{
		dbConn: dbConn,
	}), nil
}
