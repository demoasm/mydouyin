package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"log"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/errno"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type MessageCache struct {
	keyName string
	mq      *CommandQueue
}

type CreateMessageCommand struct {
	FromUserId int64
	ToUserId   int64
	Content    string
}

var MC *MessageCache

func initMessageCache() {
	MC = new(MessageCache)
	MC.keyName = "message_list"
	MC.mq = NewCommandQueue(context.Background(), "message")
	go MC.listen()
}

func (c *MessageCache) listen() {
	for {
		msg, err := c.mq.ConsumeMessage()
		if err != nil {
			continue
		}
		var cmd CreateMessageCommand
		json.Unmarshal(msg, &cmd)
		log.Printf("[********MessageCache********] recover command:%v", cmd)
		err = c.execCommand(&cmd)
		if err != nil {
			log.Printf("[********MessageCache********] command exec fail, error:%v", err)
			data, _ := json.Marshal(cmd)
			c.mq.ProductionMessage(data)
		} else {
			log.Printf("[********MessageCache********] command exec success!!!")
		}
	}
}

func (c *MessageCache) execCommand(cmd *CreateMessageCommand) error {
	resp, err := rpc.CreateMessage(c.mq.ctx, &message.CreateMessageRequest{
		FromUserId: cmd.FromUserId,
		ToUserId:   cmd.ToUserId,
		Content:    cmd.Content,
	})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	//上传成功，将新消息写入缓存
	msg := apimodel.Message{
		ID:         resp.Id,
		CreateTime: resp.CreateTime,
		ToUserId:   cmd.ToUserId,
		FromUserId: cmd.FromUserId,
		Content:    cmd.Content,
	}
	data, _ := json.Marshal(msg)
	key := md5.Sum([]byte(strconv.FormatInt(cmd.FromUserId, 10) + strconv.FormatInt(cmd.ToUserId, 10) + c.keyName))
	_, err = redisClient.SAdd(c.mq.ctx, string(key[:]), data).Result()
	return err
}

func (c *MessageCache) getMsgFromSet(from_id, to_id int64) (msg *apimodel.Message, err error) {
	key := md5.Sum([]byte(strconv.FormatUint(c.hash_key(from_id, to_id), 10) + c.keyName))
	exist, err := redisClient.Exists(c.mq.ctx, string(key[:])).Result()
	if err != nil {
		return nil, err
	}
	if exist != 1 {
		return nil, errors.New("key not exist")
	}
	result, err := redisClient.Get(c.mq.ctx, string(key[:])).Result()
	defer redisClient.Expire(c.mq.ctx, string(key[:]), 12*time.Hour)
	if err != nil {
		return nil, err
	}
	msg = new(apimodel.Message)
	json.Unmarshal([]byte(result), msg)
	return msg, nil
}

func (c *MessageCache) setMsgToSet(msg *apimodel.Message) error {
	key := md5.Sum([]byte(strconv.FormatUint(c.hash_key(msg.FromUserId, msg.ToUserId), 10) + c.keyName))
	val, _ := json.Marshal(msg)
	result, err := redisClient.Set(c.mq.ctx, string(key[:]), val, 12*time.Hour).Result()
	if err != nil || result != "OK" {
		return err
	}
	return nil
}

func (c *MessageCache) hash_key(id1, id2 int64) uint64 {
	if id1 > id2 {
		return uint64(id2 | id1<<32)
	} else {
		return uint64(id1 | (id2 << 32))
	}
}

func (c *MessageCache) GetFirstMessage(me int64, friendIds []int64) (frist_msg_list []*apimodel.FristMessage) {
	frist_msg_list = make([]*apimodel.FristMessage, 0, len(friendIds))
	for _, friendId := range friendIds {
		msg, err := c.getMsgFromSet(friendId, me)
		if err != nil {
			frist_msg_list = append(frist_msg_list, &apimodel.FristMessage{
				FriendId: friendId,
				MsgType:  -1,
			})
		} else {
			if msg.FromUserId == me {
				frist_msg_list = append(frist_msg_list, &apimodel.FristMessage{
					FriendId: friendId,
					Content:  msg.Content,
					MsgType:  1,
				})
			} else {
				frist_msg_list = append(frist_msg_list, &apimodel.FristMessage{
					FriendId: friendId,
					Content:  msg.Content,
					MsgType:  0,
				})
			}
		}
	}
	return
}

//设置最新消息
func (c *MessageCache) SetFirstMessage(msg *apimodel.Message) (err error) {
	return c.setMsgToSet(msg)
}

func (c *MessageCache) SaveMessage(messages []*apimodel.Message) error {
	for i := 0; i < len(messages); i++ {
		message, _ := json.Marshal(messages[i])

		// msgKey := md5.Sum([]byte(strconv.FormatInt(messages[i].ID, 10) + strconv.FormatInt(messages[i].ToUserId, 10) + c.keyname))
		msgKey := strconv.FormatUint(c.hash_key(messages[i].FromUserId, messages[i].ToUserId), 10) + c.keyName
		_, err := redisClient.ZAdd(c.mq.ctx, string(msgKey[:]), &redis.Z{Score: float64(messages[i].CreateTime), Member: message}).Result()
		if err != nil {
			return err
		}
		_, err1 := redisClient.Expire(c.mq.ctx, string(msgKey[:]), time.Minute*30).Result()
		if err1 != nil {
			return err1
		}
	}
	return nil
}

func (c *MessageCache) InitMessageFromDB(fromUserID int64) error {
	resp, err := rpc.GetFriend(c.mq.ctx, &relation.GetFriendRequest{
		MeId: fromUserID,
	})
	if err != nil {
		return nil
	}
	friendIds := resp.FriendIds
	for i := 0; i < len(friendIds); i++ {
		msgkey := md5.Sum([]byte(strconv.FormatInt(fromUserID, 10) + strconv.FormatInt(friendIds[i], 10) + c.keyName))
		ex, err := redisClient.Exists(c.mq.ctx, string(msgkey[:])).Result()
		if err != nil {
			return err
		}
		if ex == 1 {
			// 从数据库拉取所有我发的消息
			resp, err := rpc.GetMessageList(c.mq.ctx, &message.GetMessageListRequest{
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
		msgkey = md5.Sum([]byte(strconv.FormatInt(friendIds[i], 10) + strconv.FormatInt(fromUserID, 10) + c.keyName))
		ex, err = redisClient.Exists(c.mq.ctx, string(msgkey[:])).Result()
		if err != nil {
			return err
		}
		if ex == 1 {
			// 从数据库拉取所有发给我的消息
			resp, err := rpc.GetMessageList(c.mq.ctx, &message.GetMessageListRequest{
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
	msgkey := strconv.FormatUint(c.hash_key(fromUserID, toUserID), 10) + c.keyName
	ex, err := redisClient.Exists(c.mq.ctx, msgkey).Result()
	if err != nil {
		return nil, false, err
	}
	if ex == 0 {
		return nil, false, nil
	}
	values, err := redisClient.ZRange(c.mq.ctx, msgkey, preMsgTime, -1).Result()
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
