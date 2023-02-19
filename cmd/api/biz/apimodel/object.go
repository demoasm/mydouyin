package apimodel

import (
	"mydouyin/kitex_gen/douyincomment"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/consts"
)

type User struct {
	UserID        int64  `form:"user_id" json:"id" query:"user_id"`
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

type Video struct {
	VideoID       int64  `form:"id" json:"id" query:"id"`
	Author        User   `form:"author" json:"author" query:"author"`
	PlayUrl       string `form:"play_url" json:"play_url" query:"play_url"`
	CoverUrl      string `form:"cover_url" json:"cover_url" query:"cover_url"`
	FavoriteCount int    `form:"favorite_count" json:"favorite_count" query:"favorite_count"`
	CommentCount  int    `form:"comment_count" json:"comment_count" query:"comment_count"`
	IsFavorite    bool   `form:"is_favorite" json:"is_favorite" query:"is_favorite"`
	Title         string `form:"title" json:"title" query:"title"`
	UploadTime    string `form:"upload" json:"upload" query:"upload"`
}

func PackVideo(douyin_video *douyinvideo.Video) *Video {
	return &Video{
		VideoID: douyin_video.VideoId,
		// Author:        douyin_video.Author,
		PlayUrl:       consts.CDNURL + douyin_video.PlayUrl,
		CoverUrl:      consts.CDNURL + douyin_video.CoverUrl,
		FavoriteCount: int(douyin_video.FavoriteCount),
		CommentCount:  int(douyin_video.CommentCount),
		IsFavorite:    douyin_video.IsFavorite,
		Title:         douyin_video.Title,
		UploadTime:    douyin_video.UploadTime,
	}
}

type FriendUser struct {
	User
	Avatar  string `form:"avatar" json:"avatar" query:"avatar"`
	Message string `form:"message" json:"message" query:"message"`
	MsgType int64  `form:"msgType" json:"msgType" query:"msgType"`
}

func PackFriendUser(douyin_user *douyinuser.User) *FriendUser {
	return &FriendUser{
		User{
			UserID:        douyin_user.UserId,
			Username:      douyin_user.Username,
			FollowCount:   douyin_user.FollowCount,
			FollowerCount: douyin_user.FollowerCount,
			IsFollow:      false,
		},
		"url",
		"123",
		1,
	}
}

type Comment struct {
	CommentID  int64  `form:"id" json:"id" query:"id"`
	Commentor  User   `form:"user" json:"user" query:"user"`
	Content    string `form:"content" json:"content" query:"content"`
	CreateDate string `form:"create_data" json:"create_data" query:"create_data"`
}

func PackComment(douyin_comment *douyincomment.Comment) *Comment {
	return &Comment{
		CommentID: douyin_comment.CommentId,
		// User: douyin_comment.User
		Content:    douyin_comment.Content,
		CreateDate: douyin_comment.CreateDate,
	}
}

type Favorite struct {
	FavoriteID int64 `form:"id" json:"id" query:"id"`
	UserID     int64 `form:"user_id" json:"user_id" query:"user_id"`
	VideoID    int64 `form:"video_id" json:"video_id" query:"video_id"`
}

func PackVFavorite(douyin_favorite *douyinfavorite.Favorite) *Favorite {
	return &Favorite{
		FavoriteID: douyin_favorite.FavoriteId,
		UserID:     douyin_favorite.UserId,
		VideoID:    douyin_favorite.VideoId,
	}
}
