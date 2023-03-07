package test

import (
	//"fmt"
	"net/http"
	"testing"
)

func BenchmarkFavoriteAct(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	_ = e.POST("/douyin/favorite/action/").
		WithQuery("token", token).WithQuery("video_id", "545").WithQuery("action_type", 1).
		WithFormField("token", token).WithFormField("video_id", "545").WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	_ = e.POST("/douyin/favorite/action/").
		WithQuery("token", token).WithQuery("video_id", "545").WithQuery("action_type", 2).
		WithFormField("token", token).WithFormField("video_id", "545").WithFormField("action_type", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}
func BenchmarkGetFavoriteList(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	_ = e.GET("/douyin/favorite/list/").
		WithQuery("token", token).WithQuery("user_id", "121").
		WithFormField("token", token).WithFormField("user_id", "121").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}

