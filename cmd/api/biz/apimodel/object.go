package apimodel

import (
	"context"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyincomment"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/consts"
)

type User struct {
	UserID          int64  `form:"user_id" json:"id" query:"user_id"`
	Username        string `form:"name" json:"name" query:"name"`
	FollowCount     int64  `form:"follow_count" json:"follow_count" query:"follow_count"`
	FollowerCount   int64  `form:"follower_count" json:"follower_count" query:"follower_count"`
	IsFollow        bool   `form:"is_follow" json:"is_follow" query:"is_follow"`
	Avatar          string `form:"avatar" json:"avatar" query:"avatar"`
	BackgroundImage string `form:"background_image" json:"background_image" query:"background_image"`
	Signature       string `form:"signature" json:"signature" query:"signature"`
	TotalFavoried   int64  `form:"total_favoried" json:"total_favoried" query:"total_favoried"`
	WorkCount       int64  `form:"work_count" json:"work_count" query:"work_count"`
	FavoriteCount   int64  `form:"favorite_count" json:"favorite_count" query:"favorite_count"`
}

var avatar_list map[int]string = map[int]string{
	0: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3uldzkb7ij309q09qjsn.jpg",
	1: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3uldztsvxj309q09qdha.jpg",
	2: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule03d3zj309q09qjsm.jpg",
	3: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule0ckvpj309q09qwfh.jpg",
	4: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule0jgguj309q09qmya.jpg",
	5: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule0vqnhj309q09qwg2.jpg",
	6: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule1a2d3j309q09q0tp.jpg",
	7: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule1j42xj309q09qjsx.jpg",
	8: "https://maomint.maomint.cn/douyin/avatar/006LfQcply1g3ule1szakj309q09qta0.jpg",
}

var background_list map[int]string = map[int]string{
	0: "https://maomint.maomint.cn/douyin/background/125615ape48gysysgxbx0y.jpg",
	1: "https://maomint.maomint.cn/douyin/background/125620l6lecc441lilqej6.jpg",
	2: "https://maomint.maomint.cn/douyin/background/125631yyvjdud5j5tjm9m1.jpg",
	3: "https://maomint.maomint.cn/douyin/background/index.jpg",
}

func PackUser(douyin_user *douyinuser.User) *User {
	return &User{
		UserID:          douyin_user.UserId,
		Username:        douyin_user.Username,
		FollowCount:     douyin_user.FollowCount,
		FollowerCount:   douyin_user.FollowerCount,
		Avatar:          avatar_list[int(douyin_user.UserId)%len(avatar_list)],
		BackgroundImage: background_list[int(douyin_user.UserId)%len(background_list)],
		Signature:       "你妈死了",
		TotalFavoried:   0,
		WorkCount:       0,
		FavoriteCount:   0,
		IsFollow:        false,
	}
}

func PackUserRelation(douyin_user *douyinuser.User, me int64) *User {
	user := PackUser(douyin_user)
	r, err := rpc.ValidIfFollowRequest(context.Background(), &relation.ValidIfFollowRequest{FollowId: user.UserID, FollowerId: me})
	if err != nil || r.BaseResp.StatusCode != 0 {
		return user
	}
	user.IsFollow = r.IfFollow
	return user
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

func PackVideos(douyin_videos []*douyinvideo.Video) []*Video {
	res := make([]*Video, 0, 30)
	for _, douyin_video := range douyin_videos {
		res = append(res, PackVideo(douyin_video))
	}
	return res
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


