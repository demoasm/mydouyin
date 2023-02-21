package cache

import (
	"context"
	"encoding/json"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/relation"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type MessageCache struct {
	keyname string
	ctx     context.Context
}

var MC *MessageCache

func initMessageCache() {
	MC = new(MessageCache)
	MC.keyname = "message_list"
	//MC.mq = NewCommandQueue(context.Background(), "message")
	//go MC.listen()
	MC.ctx = context.Background()
}

func (c *MessageCache) hash_key(id1, id2 int64) uint64 {
	if id1 > id2 {
		return uint64(id2 | id1<<32)
	} else {
		return uint64(id1 | (id2 << 32))
	}
}

func (c *MessageCache) SaveMessage(messages []*apimodel.Message) error {
	for i := 0; i < len(messages); i++ {
		message, _ := json.Marshal(messages[i])
		msgKey := strconv.FormatUint(c.hash_key(messages[i].FromUserId, messages[i].ToUserId), 10) + c.keyname
		_, err := redisClient.ZAdd(c.ctx, msgKey, &redis.Z{Score: float64(messages[i].CreateTime), Member: message}).Result()
		if err != nil {
			return err
		}
		_, err1 := redisClient.Expire(c.ctx, msgKey, time.Minute*30).Result()
		if err1 != nil {
			return err1
		}
	}
	return nil
}

func (c *MessageCache) InitMessageFromDB(fromUserID int64) error {
	resp, err := rpc.GetFriend(c.ctx, &relation.GetFriendRequest{
		MeId: fromUserID,
	})
	if err != nil {
		return nil
	}
	friendIds := resp.FriendIds
	for i := 0; i < len(friendIds); i++ {
		msgkey := strconv.FormatUint(c.hash_key(fromUserID, friendIds[i]), 10) + c.keyname
		ex, err := redisClient.Exists(c.ctx, msgkey).Result()
		if err != nil {
			return err
		}

		if ex == 0 {

			messageList := make([]*apimodel.Message, 0)
			// 从数据库拉取所有我发的消息
			resp, err := rpc.GetMessageList(c.ctx, &message.GetMessageListRequest{
				FromUserId: fromUserID,
				ToUserId:   friendIds[i],
				PreMsgTime: 0,
			})
			if err != nil {
				return err
			}
			messageList = append(messageList, apimodel.PackMessages(resp.MessageList)...)
			// 从数据库拉取所有发给我的消息
			resp, err = rpc.GetMessageList(c.ctx, &message.GetMessageListRequest{
				FromUserId: friendIds[i],
				ToUserId:   fromUserID,
				PreMsgTime: 0,
			})
			if err != nil {
				return err
			}
			messageList = append(messageList, apimodel.PackMessages(resp.MessageList)...)

			// 存入缓存
			err = c.SaveMessage(messageList)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *MessageCache) GetMessage(fromUserID int64, toUserID int64, preMsgTime int64) ([]*apimodel.Message, bool, error) {
	messageList := make([]*apimodel.Message, 0)
	msgkey := strconv.FormatUint(c.hash_key(fromUserID, toUserID), 10) + c.keyname
	ex, err := redisClient.Exists(c.ctx, msgkey).Result()
	if err != nil {
		return nil, false, err
	}
	if ex == 0 {
		return nil, false, nil
	}
	values, err := redisClient.ZRange(c.ctx, msgkey, preMsgTime, -1).Result()
	if err != nil {
		return nil, false, err
	}
	for i := 0; i < len(values); i++ {
		message := new(apimodel.Message)
		err = json.Unmarshal([]byte(values[i]), &message)
		if err != nil {
			continue
		}
		messageList = append(messageList, message)
	}
	return messageList, true, nil
}
