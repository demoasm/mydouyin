package test

import (
	"net/http"
	"testing"

)

func BenchmarkPublish(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"

	_ = e.POST("/douyin/publish/action/").
		WithMultipart().
		WithFile("data", "../public/bear.mp4").
		WithFormField("token", token).
		WithFormField("title", "Bear").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	// publishResp.Value("status_code").Number().Equal(0)



}

func BenchmarkGetList(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	 _ = e.GET("/douyin/publish/list/").
		WithQuery("user_id", 121).WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}


func BenchmarkFeed(b *testing.B) {
	e:=newExpect(b)
	_ = e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
}