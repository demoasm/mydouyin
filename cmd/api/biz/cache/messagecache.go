package cache

import (
	"context"
	"crypto/md5"
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

		// msgKey := md5.Sum([]byte(strconv.FormatInt(messages[i].ID, 10) + strconv.FormatInt(messages[i].ToUserId, 10) + c.keyname))
		msgKey := strconv.FormatUint(c.hash_key(messages[i].FromUserId, messages[i].ToUserId), 10) + c.keyname
		_, err := redisClient.ZAdd(c.ctx, string(msgKey[:]), &redis.Z{Score: float64(messages[i].CreateTime), Member: message}).Result()
		if err != nil {
			return err
		}
		_, err1 := redisClient.Expire(c.ctx, string(msgKey[:]), time.Minute*30).Result()
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
		msgkey := md5.Sum([]byte(strconv.FormatInt(fromUserID, 10) + strconv.FormatInt(friendIds[i], 10) + c.keyname))
		ex, err := redisClient.Exists(c.ctx, string(msgkey[:])).Result()
		if err != nil {
			return err
		}
		if ex == 1 {
			// 从数据库拉取所有我发的消息
			resp, err := rpc.GetMessageList(c.ctx, &message.GetMessageListRequest{
				FromUserId: fromUserID,
				ToUserId:   friendIds[i],
				PreMsgTime: time.Now().Unix(),
			})
			if err != nil {
				return nil
			}
			// 存入缓存
			err = c.SaveMessage(apimodel.PackMessages(resp.MessageList))
			if err != nil {
				return err
			}
		}
		msgkey = md5.Sum([]byte(strconv.FormatInt(friendIds[i], 10) + strconv.FormatInt(fromUserID, 10) + c.keyname))
		ex, err = redisClient.Exists(c.ctx, string(msgkey[:])).Result()
		if err != nil {
			return err
		}
		if ex == 1 {
			// 从数据库拉取所有发给我的消息
			resp, err := rpc.GetMessageList(c.ctx, &message.GetMessageListRequest{
				FromUserId: friendIds[i],
				ToUserId:   fromUserID,
				PreMsgTime: time.Now().Unix(),
			})
			if err != nil {
				return nil
			}
			// 存入缓存
			err = c.SaveMessage(apimodel.PackMessages(resp.MessageList))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *MessageCache) GetMessage(fromUserID int64, toUserID int64, preMsgTime int64) ([]*apimodel.Message, bool, error) {
	messageList := make([]*apimodel.Message, 0)
	// log.Println(fromUserID, toUserID)
	// msgkey := md5.Sum([]byte(strconv.FormatInt(fromUserID, 10) + strconv.FormatInt(toUserID, 10) + c.keyname))
	msgkey := strconv.FormatUint(c.hash_key(fromUserID, toUserID), 10) + c.keyname
	ex, err := redisClient.Exists(c.ctx, msgkey).Result()
	if err != nil {
		return nil, false, err
	}
	if ex == 0 {
		return nil, false, nil
	}
	values, err := redisClient.ZRange(c.ctx, msgkey, preMsgTime, -1).Result()
	// log.Println(values)
	if err != nil {
		return nil, false, err
	}
	for i := 0; i < len(values); i++ {
		message := new(apimodel.Message)
		err = json.Unmarshal([]byte(values[i]), &message)
		messageList = append(messageList, message)
		if err != nil {
			continue
		}
	}
	return messageList, true, nil

	// if(len(values) == 0){
	// 	rpc_resp, err := rpc.GetMessageList(c.ctx, &message.GetMessageListRequest{
	// 		FromUserId: fromUserID,
	// 		ToUserId: toUserID,
	// 		PreMsgTime: preMsgTime,
	// 	})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	messages := apimodel.PackMessages(rpc_resp.MessageList)
	// 	c.SaveMessage(messages)
	// 	return messages, nil
	// }else{

	// }
}
