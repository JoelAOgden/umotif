package mock

import (
	"context"
	"fmt"
)

type SqsClient struct {
}

func (s SqsClient) Push(ctx context.Context, message string) error {
	// todo: this
	fmt.Println("message pushed to sqs- ", message)
	return nil
}
