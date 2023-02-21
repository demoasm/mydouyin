package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/cache"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/message"
	"mydouyin/pkg/errno"
	"sort"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) MessageAction(req apimodel.MessageActionRequest, user *apimodel.User) (resp *apimodel.MessageActionResponse, err error) {
	resp = new(apimodel.MessageActionResponse)
	rpc_resp, err := rpc.CreateMessage(s.ctx, &message.CreateMessageRequest{
		FromUserId: user.UserID,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	return resp, nil
}

func (s *MessageService) MessageChat(req apimodel.MessageChatRequest, user *apimodel.User) (resp *apimodel.MessageChatResponse, err error) {
	resp = new(apimodel.MessageChatResponse)
	messageList, hit, err := cache.MC.GetMessage(user.UserID, req.ToUserId, req.PreMsgTime)
	if err != nil{
		return
	}
	if hit{
		resp.MessageList = messageList
		return 
	}
	rpc_resp_from, err := rpc.GetMessageList(s.ctx, &message.GetMessageListRequest{
		FromUserId: user.UserID,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp_from.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp_from.BaseResp.StatusCode, rpc_resp_from.BaseResp.StatusMessage)
	}
	message_list_from := apimodel.PackMessages(rpc_resp_from.MessageList)
	rpc_resp_to, err := rpc.GetMessageList(s.ctx, &message.GetMessageListRequest{
		FromUserId: req.ToUserId,
		ToUserId:   user.UserID,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp_to.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp_to.BaseResp.StatusCode, rpc_resp_to.BaseResp.StatusMessage)
	}
	message_list_to := apimodel.PackMessages(rpc_resp_to.MessageList)
	resp.MessageList = append(message_list_from, message_list_to...)
	sort.Sort(apimodel.MessageSorter(resp.MessageList))
	cache.MC.SaveMessage(resp.MessageList)
	return
}
