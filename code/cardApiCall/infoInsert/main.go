package main

import (
	"context"
	"encoding/json"

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

func HandleRequest(ctx context.Context, event json.RawMessage) (string, error) {
	var input Input
	if err := json.Unmarshal(event, &input); err != nil {
		return "", err
	}

	injector := app.SetupDIContainer()
	defer injector.Shutdown()

	masterUseCase := do.MustInvoke[master.MasterUseCase](injector)

	failedIds, err := masterUseCase.InsertCardInfoList(ctx, input.StartId, input.Delta)
	if err != nil {
		return "", err
	}

	output := Output{FailedIds: failedIds}
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func main() {
	lambda.Start(HandleRequest)
}