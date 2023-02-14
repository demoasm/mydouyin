package apimodel

import (
	"mydouyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response interface {
	SetErr(err error)
	Send(c *app.RequestContext)
}

// @router /douyin/user/register [POST] Response
type CreateUserResponse struct {
	StatusCode int64  `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `form:"status_msg" json:"status_msg" query:"status_msg"`
	UserId     int64  `form:"user_id" json:"user_id" query:"user_id"`
	Token      string `form:"token" json:"token" query:"token"`
}

func (res *CreateUserResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *CreateUserResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/user/login [POST] Response
type CheckUserResponse struct {
	StatusCode int64  `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `form:"status_msg" json:"status_msg" query:"status_msg"`
	UserId     int64  `form:"user_id" json:"user_id" query:"user_id"`
	Token      string `form:"token" json:"token" query:"token"`
}

func (res *CheckUserResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *CheckUserResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/user/ [GET]
type GetUserResponse struct {
	StatusCode int64  `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `form:"status_msg" json:"status_msg" query:"status_msg"`
	User       User   `form:"user" json:"user" query:"user"`
}

func (res *GetUserResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *GetUserResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/feed/ [GET]
type GetFeedResponse struct {
	StatusCode int64   `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string  `form:"status_msg" json:"status_msg" query:"status_msg"`
	NextTime   int64   `form:"next_time" json:"next_time" query:"next_time"`
	VideoList  []Video `form:"video_list" json:"video_list" query:"video_list"`
}

func (res *GetFeedResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *GetFeedResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/publish/action/ [POST]
type PublishVideoResponse struct {
	StatusCode int64  `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `form:"status_msg" json:"status_msg" query:"status_msg"`
}

func (res *PublishVideoResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *PublishVideoResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/publish/list/ [GET]
type GetPublishListResponse struct {
	StatusCode int64   `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string  `form:"status_msg" json:"status_msg" query:"status_msg"`
	VideoList  []Video `form:"video_list" json:"video_list" query:"video_list"`
}

func (res *GetPublishListResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *GetPublishListResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/comment/action/ [POST]
type CommentActionResponse struct {
	StatusCode int64   `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string  `form:"status_msg" json:"status_msg" query:"status_msg"`
	Comment    Comment `form:"comment" json:"comment" query:"comment"`
}

func (res *CommentActionResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *CommentActionResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}

// @router /douyin/comment/list/ [GET]
type CommentListResponse struct {
	StatusCode  int64     `form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg   string    `form:"status_msg" json:"status_msg" query:"status_msg"`
	CommentList []Comment `form:"comment_list" json:"comment_list" query:"comment_list"`
}

func (res *CommentListResponse) Send(c *app.RequestContext) {
	c.JSON(consts.StatusOK, res)
}

func (res *CommentListResponse) SetErr(err error) {
	Err := errno.ConvertErr(err)
	res.StatusCode = Err.ErrCode
	res.StatusMsg = Err.ErrMsg
}
