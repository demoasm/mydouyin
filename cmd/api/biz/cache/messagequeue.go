package cache

import (
	"context"
)

//消息队列，调用ProductionMessage将指令结构体放入redis的list中，另起一个线程调用ConsumeMessage接收消息，反序列化指令结构体交给自定义的handler函数执行
//在videoHandel中有使用
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
