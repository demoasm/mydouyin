package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/errno"
	"strconv"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) RelationAction(req apimodel.RelationActionRequest, user *apimodel.User) (*apimodel.RelationActionResponse, error) {
	resp := new(apimodel.RelationActionResponse)
	userId := user.UserID
	to_user_id, err := strconv.Atoi(req.ToUserId)
	if err != nil {
		return resp, errno.ParamErr
	}

	switch req.ActionType {
	case "1":
		rpc_resp, err := rpc.CreateRelation(s.ctx, &relation.CreateRelationRequest{
			FollowId:   int64(to_user_id),
			FollowerId: userId,
		})
		if err != nil {
			return resp, err
		}
		if rpc_resp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.StatusCode, rpc_resp.StatusMessage)
		}
	case "2":
		rpc_resp, err := rpc.DeleteRelation(s.ctx, &relation.DeleteRelationRequest{
			FollowId:   int64(to_user_id),
			FollowerId: userId,
		})
		if err != nil {
			return resp, err
		}
		if rpc_resp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.StatusCode, rpc_resp.StatusMessage)
		}
	default:
		err = errno.ParamErr
		return resp, err
	}
	return resp, nil
}

//获取关注或粉丝列表，option表示操作类型(1：关注列表，2：粉丝列表)
func (s *RelationService) FollowAndFollowerList(req apimodel.FollowAndFollowerListRequest, user *apimodel.User, option int) (*apimodel.FollowAndFollowerListReponse, error) {
	resp := new(apimodel.FollowAndFollowerListReponse)
	var err error
	// users := make([]*apimodel.User, 0)
	userIds := make([]int64, 0)
	switch option {
	case 1:
		rpc_resp, err := rpc.GetFollow(s.ctx, &relation.GetFollowListRequest{FollowerId: int64(user.UserID)})
		if err != nil {
			return resp, err
		}
		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
		if len(rpc_resp.FollowIds) < 1 {
			return resp, nil
		}
		userIds = rpc_resp.FollowIds
	case 2:
		rpc_resp, err := rpc.GetFollower(s.ctx, &relation.GetFollowerListRequest{FollowId: int64(user.UserID)})
		if err != nil {
			return resp, err
		}
		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
		if len(rpc_resp.FollowerIds) < 1 {
			return resp, nil
		}
		userIds = rpc_resp.FollowerIds
	}
	ur, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}
	if ur.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
	}
	for _, rpc_user := range ur.Users {
		switch option {
		case 1:
			u := apimodel.PackUser(rpc_user)
			u.IsFollow = true
			resp.UserList = append(resp.UserList, u)
		case 2:
			u := apimodel.PackUserRelation(rpc_user, int64(user.UserID))
			resp.UserList = append(resp.UserList, u)
		}
	}
	// resp.UserList = users
	return resp, errno.Success
}
