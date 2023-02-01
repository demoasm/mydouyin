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
