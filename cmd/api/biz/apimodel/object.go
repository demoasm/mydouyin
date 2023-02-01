package apimodel

import "mydouyin/kitex_gen/douyinuser"

type User struct {
	UserID        int64  `form:"user_id" json:"user_id" query:"user_id"`
	Username      string `form:"name" json:"name" query:"name"`
	FollowCount   int64  `form:"follow_count" json:"follow_count" query:"follow_count"`
	FollowerCount int64  `form:"follower_count" json:"follower_count" query:"follower_count"`
	IsFollow      bool   `form:"is_follow" json:"is_follow" query:"is_follow"`
}

func PackUser(douyin_user *douyinuser.User) *User {
	return &User{
		UserID:        douyin_user.UserId,
		Username:      douyin_user.Username,
		FollowCount:   douyin_user.FollowCount,
		FollowerCount: douyin_user.FollowerCount,
		IsFollow:      false,
	}
}
