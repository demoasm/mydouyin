package test

import (
	//"fmt"
	"net/http"
	"testing"
)

func BenchmarkCommentAct(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	CommentResp := e.POST("/douyin/comment/action/").
		WithQuery("token", token).WithQuery("video_id", "545").WithQuery("action_type", 1).WithQuery("comment_text", "测试评论").
		WithFormField("token", token).WithFormField("video_id", "545").WithFormField("action_type", 1).WithFormField("comment_text", "测试评论").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	commentId := int(CommentResp.Value("comment").Object().Value("id").Number().Raw())
	_ = e.POST("/douyin/comment/action/").
		WithQuery("token", token).WithQuery("video_id", "545").WithQuery("action_type", 2).WithQuery("comment_id", commentId).
		WithFormField("token", token).WithFormField("video_id", "545").WithFormField("action_type", 2).WithFormField("comment_id", commentId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}

func BenchmarkGetCommentList(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	_ = e.GET("/douyin/comment/list/").
		WithQuery("token", token).WithQuery("video_id", "545").
		WithFormField("token", token).WithFormField("video_id", "545").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}
