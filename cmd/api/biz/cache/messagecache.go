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
	"mydouyin/pkg/errno"
	"strconv"
	"time"
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

//当接收到messageChat请求时调用该函数从缓存中拉取全部最新消息，如果
func (c *MessageCache) GetMessageList(from_user_id, to_user_id int64) []*apimodel.Message {
	key := md5.Sum([]byte(strconv.FormatInt(from_user_id, 10) + strconv.FormatInt(to_user_id, 10) + c.keyName))
	results, err := redisClient.SMembers(c.mq.ctx, string(key[:])).Result()
	if err != nil {
		return []*apimodel.Message{}
	}
	msg_list := make([]*apimodel.Message, 0, 50)
	for _, result := range results {
		msg := new(apimodel.Message)
		err = json.Unmarshal([]byte(result), msg)
		if err != nil {
			continue
		}
		msg_list = append(msg_list, msg)
	}
	return msg_list
}

func (c *MessageCache) AddMessage(from_user_id, to_user_id int64, content string) error {
	//提交增加message的指令
	cmd, _ := json.Marshal(CreateMessageCommand{
		FromUserId: from_user_id,
		ToUserId:   to_user_id,
		Content:    content,
	})
	err := c.mq.ProductionMessage(cmd)
	if err != nil {
		return err
	}
	return nil
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

//传入我的id和我朋友列表的id，返回对应朋友的第一条消息，若未命中则未nil，以及err
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
