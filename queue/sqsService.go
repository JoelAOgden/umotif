package queue

import "context"

type SqsClient interface {
	Push(ctx context.Context, message string) error // this is wrong but we can change this later
}

type Message struct {
	Body string
}

type SqsService struct {
	Client SqsClient
}

func (s SqsService) PushToQueue(ctx context.Context, msg Message) error {
	return s.Client.Push(ctx, msg.Body)
}
