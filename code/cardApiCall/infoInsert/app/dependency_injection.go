package app

import (
	"atomisu.com/ocg-statics/infoInsert/config"
	"atomisu.com/ocg-statics/infoInsert/usecase"

	_ "github.com/lib/pq" // PostgreSQL ドライバをインポート
	"github.com/samber/do"
)

// SetupDIContainer はDIコンテナをセットアップします
// テストなどでコンテナの中身をオーバーライドしたい場合には、
// samba/doのOverride関数を使って上書きしてください
func SetupDIContainer() *do.Injector {
	injector := do.New()
	// DBコネクション
	do.Provide(injector, config.NewDbConnection)

	// ユースケース
	do.Provide(injector, usecase.NewNeuronUseCase)
	//do.Provide(injector, usecase.NewAuthUseCase)

	// ハンドラー
	//do.Provide(injector, handler.NewAccountHandler)
	//do.Provide(injector, handler.NewAuthHandler)

	return injector
}
