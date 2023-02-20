package cache

import (
	"context"
)

type MessageQueue struct {
	ListName string
	ctx      context.Context
}

func NewMessageQueue(ctx context.Context, listName string) *MessageQueue {
	return &MessageQueue{
		ListName: listName,
		ctx:      ctx,
	}
}

func (mq *MessageQueue) ProductionMessage(message []byte) error {
	_, err := redisClient.LPush(mq.ctx, mq.ListName, message).Result()
	return err
}

func (mq *MessageQueue) ConsumeMessage() ([]byte, error) {
	item, err := redisClient.BRPop(mq.ctx, 0, mq.ListName).Result()
	if err != nil {
		return nil, err
	}
	return []byte(item[1]), err
}
