// Code generated by hertz generator.

package douyinapi

import (
	"context"
	"strconv"

	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/mw"
	"mydouyin/cmd/api/biz/rpc"
	videohandel "mydouyin/cmd/api/biz/videoHandel"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

//基础接口
// GetFeed
// @router /douyin/feed/ [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.GetFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	resp := new(apimodel.GetFeedResponse)

	defer func() {
		resp.SetErr(err)
		resp.Send(c)
	}()

	resp.VideoList, resp.NextTime, err = rpc.GetFeed(context.Background(), &douyinvideo.GetFeedRequest{
		LatestTime: req.LatestTime,
		UserId:     -1,
	})

	if err != nil {
		return
	}
	err = errno.Success
}

// GetPublishList
// @router /douyin/publish/list [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.GetPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	resp := new(apimodel.GetPublishListResponse)

	defer func() {
		resp.SetErr(err)
		resp.Send(c)
	}()

	resp.VideoList, err = rpc.GetPublishList(context.Background(), &douyinvideo.GetListRequest{
		UserId: int64(userId),
	})
	if err != nil {
		return
	}
	err = errno.Success
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
	videourl, coverurl, err := videohandel.VH.UpLoadFile(req.Data)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	// hlog.Infof("上传的url为%v", videourl, coverurl)
	resp := new(apimodel.PublishVideoResponse)

	err = rpc.PublishVideo(context.Background(), &douyinvideo.CreateVideoRequest{
		Author:   user.(*apimodel.User).UserID,
		PlayUrl:  videourl,
		CoverUrl: coverurl,
		Title:    req.Title,
	})

	if err != nil {
		resp.SetErr(err)
		resp.Send(c)
	}

	resp.SetErr(errno.Success)
	resp.Send(c)
}

// CreateUser .
// @router /douyin/user/register/ [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req apimodel.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		// c.String(consts.StatusBadRequest, err.Error())
		SendResponse(c, err, nil)
		return
	}

	resp := new(apimodel.CreateUserResponse)

	err = rpc.CreateUser(context.Background(), &douyinuser.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

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
	resp := new(apimodel.GetUserResponse)
	defer func() {
		resp.SetErr(err)
		resp.Send(c)
	}()
	id, err := strconv.Atoi(req.UserID)
	if err != nil {
		err = errno.ParamErr
		return
	}
	user, err1 := rpc.GetUser(context.Background(), &douyinuser.MGetUserRequest{[]int64{int64(id)}})
	if err1 != nil {
		err = err1
		return
	}
	resp.User = *user
	err = errno.Success
}
