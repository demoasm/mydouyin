package test

import (
	//"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavorite(t *testing.T) {
	e := newExpect(t)

	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().Equal(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	firstVideo := feedResp.Value("video_list").Array().First().Object()
	videoId := firstVideo.Value("id").Number().Raw()

	userId, token := getTestUserToken(testUserA, e)

	favoriteResp := e.POST("/douyin/favorite/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 1).
		WithFormField("token", token).WithFormField("video_id", videoId).WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteResp.Value("status_code").Number().Equal(0)

	favoriteListResp := e.GET("/douyin/favorite/list/").
		WithQuery("token", token).WithQuery("user_id", userId).
		WithFormField("token", token).WithFormField("user_id", userId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteListResp.Value("status_code").Number().Equal(0)

	//fmt.Println("*************", favoriteListResp)

	for _, element := range favoriteListResp.Value("video_list").Array().Iter() {

		video := element.Object()
		//fmt.Println("*************", video.Value("id").Number().Raw())
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}

	delfavoriteResp := e.POST("/douyin/favorite/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 2).
		WithFormField("token", token).WithFormField("video_id", videoId).WithFormField("action_type", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	delfavoriteResp.Value("status_code").Number().Equal(0)

	favoriteListResp = e.GET("/douyin/favorite/list/").
		WithQuery("token", token).WithQuery("user_id", userId).
		WithFormField("token", token).WithFormField("user_id", userId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteListResp.Value("status_code").Number().Equal(0)
}

func TestComment(t *testing.T) {
	e := newExpect(t)

	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().Equal(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	firstVideo := feedResp.Value("video_list").Array().First().Object()
	videoId := firstVideo.Value("id").Number().Raw()

	_, token := getTestUserToken(testUserA, e)

	addCommentResp := e.POST("/douyin/comment/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 1).WithQuery("comment_text", "测试评论").
		WithFormField("token", token).WithFormField("video_id", videoId).WithFormField("action_type", 1).WithFormField("comment_text", "测试评论").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	addCommentResp.Value("status_code").Number().Equal(0)
	addCommentResp.Value("comment").Object().Value("id").Number().Gt(0)
	commentId := int(addCommentResp.Value("comment").Object().Value("id").Number().Raw())

	commentListResp := e.GET("/douyin/comment/list/").
		WithQuery("token", token).WithQuery("video_id", videoId).
		WithFormField("token", token).WithFormField("video_id", videoId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	commentListResp.Value("status_code").Number().Equal(0)

	containTestComment := false
	for _, element := range commentListResp.Value("comment_list").Array().Iter() {

		comment := element.Object()
		if int(comment.Value("id").Number().Raw()) == commentId {
			containTestComment = true
		}
		comment.ContainsKey("id")
		comment.ContainsKey("user")
		comment.Value("content").String().NotEmpty()
		//fmt.Println("****************", comment.Value("create_date").String())
		comment.Value("create_date").String().NotEmpty()
	}
	assert.True(t, containTestComment, "Can't find test comment in list")

	delCommentResp := e.POST("/douyin/comment/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 2).WithQuery("comment_id", commentId).
		WithFormField("token", token).WithFormField("video_id", videoId).WithFormField("action_type", 2).WithFormField("comment_id", commentId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	delCommentResp.Value("status_code").Number().Equal(0)
	commentListResp = e.GET("/douyin/comment/list/").
		WithQuery("token", token).WithQuery("video_id", videoId).
		WithFormField("token", token).WithFormField("video_id", videoId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	commentListResp.Value("status_code").Number().Equal(0)

	commentListResp = e.GET("/douyin/comment/list/").
		WithQuery("token", token).WithQuery("video_id", videoId).
		WithFormField("token", token).WithFormField("video_id", videoId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	commentListResp.Value("status_code").Number().Equal(0)

	containTestComment = false
	for _, element := range commentListResp.Value("comment_list").Array().Iter() {

		comment := element.Object()
		if int(comment.Value("id").Number().Raw()) == commentId {
			containTestComment = true
		}
		comment.ContainsKey("id")
		comment.ContainsKey("user")
		comment.Value("content").String().NotEmpty()
		//fmt.Println("****************", comment.Value("create_date").String())
		comment.Value("create_date").String().NotEmpty()
	}
	assert.False(t, containTestComment, "Can't find test comment in list")


}
