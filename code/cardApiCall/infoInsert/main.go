package main

import (
	"context"
	"encoding/json"
	"fmt"

	"atomisu.com/ocg-statics/infoInsert/app"
	"atomisu.com/ocg-statics/infoInsert/usecase/master"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/samber/do"
)

type Input struct {
	StartId int64 `json:"startId"`
	Delta   int64 `json:"delta"`
}

type Output struct {
	FailedIds []int64 `json:"failedIds"`
}

// DIコンテナはグローバルで1回だけ初期化（Lambda cold start を考慮）
var (
	globalInjector = app.SetupDIContainer()
)

func init() {
	// 必要であればここでログやメトリクス初期化
}

func HandleRequest(ctx context.Context, event json.RawMessage) (Output, error) {
	var input Input
	if err := json.Unmarshal(event, &input); err != nil {
		return Output{}, err
	}

	// 入力バリデーション
	if input.Delta <= 0 {
		return Output{}, fmt.Errorf("delta must be > 0")
	}

	masterUseCase := do.MustInvoke[master.MasterUseCase](globalInjector)

	// カード情報を挿入
	failedIds, err := masterUseCase.InsertCardInfoList(ctx, input.StartId, input.Delta)
	if err != nil {
		return Output{}, err
	}

	return Output{FailedIds: failedIds}, nil
}

func main() {
	defer globalInjector.Shutdown()
	lambda.Start(HandleRequest)
}
