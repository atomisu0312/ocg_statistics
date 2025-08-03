package main

import (
	"context"
	"os/user"

	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, user user.User) (string, error) {
	return fmt.Sprintf("Hello %s!", user.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
