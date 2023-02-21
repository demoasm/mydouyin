// Code generated by hertz generator.

package douyinapi

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/mw"
	"mydouyin/cmd/api/biz/service"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

//	基础接口
//
// FavoriteAction
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	var err error
	var req apimodel.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	resp, err := service.NewFavoriteService(ctx).FavoriteAction(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// GetFavoriteList
// @router /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {

	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}

	var err error
	var req apimodel.GetFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	resp, err := service.NewFavoriteService(ctx).GetFavoriteList(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(errno.Success)
		resp.Send(c)
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// GetFeed
// @router /douyin/feed/ [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	user, exists := c.Get(consts.IdentityKey)
	var userId int64 = -1
	if exists {
		userId = user.(*apimodel.User).UserID
	}
	var err error
	var req apimodel.GetFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	resp, err := service.NewFeedService(ctx).GetFeed(req, userId)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// GetPublishList
// @router /douyin/publish/list [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}

	var err error
	var req apimodel.GetPublishListRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	resp, err := service.NewPublishService(ctx).GetPublishList(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// Publish Video
// @router /douyin/publish/action/ [POST]
func PublishVideo(ctx context.Context, c *app.RequestContext) {
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	var err error
	var req apimodel.PublishVideoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}
	resp, err := service.NewPublishService(ctx).PublishVideo(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// CreateUser .
// @router /douyin/user/register/ [POST]
func RegistUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.RegistUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	resp, err := service.NewUserService(ctx).RegistUser(req)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// CheckUser .
// @router /douyin/user/login/ [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)

}

// GetUser .
// @router /douyin/user/ [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.GetUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	resp, err := service.NewUserService(ctx).GetUser(req)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
	}


	resp.SetErr(errno.Success)
	resp.Send(c)
}

// CommentAction .
// @router /douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	var err error
	var req apimodel.CommentActionRequest
	err = c.BindAndValidate((&req))
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	resp, err := service.NewCommentService(ctx).CommentAction(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}

	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	resp, err := service.NewCommentService(ctx).CommentList(req)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	resp, err := service.NewRelationService(ctx).RelationAction(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/relation/follow/list/ [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.FollowAndFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	resp, err := service.NewRelationService(ctx).FollowAndFollowerList(req, user.(*apimodel.User), 1)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/relation/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.FollowAndFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	resp, err := service.NewRelationService(ctx).FollowAndFollowerList(req, user.(*apimodel.User), 2)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/relation/friend/list/ [GET] 开发中....
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.FriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	resp, err := service.NewRelationService(ctx).FriendList(req)
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/message/chat/ [GET] 开发中....
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	resp, err := service.NewMessageService(ctx).MessageChat(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}

// @router /douyin/message/action/ [POST] 开发中....
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	user, exists := c.Get(consts.IdentityKey)
	if !exists {
		SendResponse(c, errno.AuthorizationFailedErr, nil)
		return
	}
	resp, err := service.NewMessageService(ctx).MessageAction(req, user.(*apimodel.User))
	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
		return
	}
	resp.SetErr(errno.Success)
	resp.Send(c)
}
