package main

import (
	"context"
	"encoding/json"

	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event json.RawMessage) (string, error) {

	var input Input
	if err := json.Unmarshal(event, &input); err != nil {
		return "", err
	}

	return fmt.Sprintf("name: %s", input.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
