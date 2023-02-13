package apimodel

import "mime/multipart"

type CreateUserRequest struct {
	Username string `json:"username" query:"username" vd:"len($) > 0"`
	Password string `json:"password" query:"password" vd:"len($) > 0"`
}

type CheckUserRequest struct {
	Username string `json:"username" query:"username" vd:"len($) > 0"`
	Password string `json:"password" query:"password" vd:"len($) > 0"`
}

type GetUserRequest struct {
	UserID string `json:"user_id" query:"user_id"`
	Token  string `json:"token" query:"token"`
}

type GetFeedRequest struct {
	LatestTime string `json:"latest_time" query:"latest_time"`
	Token      string `json:"token" query:"token"`
}

type PublishVideoRequest struct {
	Data  *multipart.FileHeader `json:"data" form:"data"`
	Token string                `json:"token" form:"token"`
	Title string                `json:"title" form:"title"`
}

type GetPublishListRequest struct {
	Token  string `json:"token" query:"token"`
	UserId string `json:"user_id" query:"user_id"`
}

type RelationActionRequest struct {
	Token      string `json:"token" query:"token"`
	ToUserId   string `json:"to_user_id" query:"to_user_id"`
	ActionType string `json:"action_type" query:"action_type"`
}

type FollowAndFollowerListRequest struct {
	UserId string `json:"user_id" query:"user_id"`
	Token  string `json:"token" query:"token"`
}
