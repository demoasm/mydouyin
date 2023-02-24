package test

import (
	"net/http"
	"testing"

)

func BenchmarkTestFeed(b *testing.B) {
	e:=newExpect(b)
	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().Equal(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	for _, element := range feedResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}

}



func BenchmarkTestPublish(b *testing.B) {
	e := newExpect(b)

	userId, token := getTestUserToken(testUserA, e)

	publishResp := e.POST("/douyin/publish/action/").
		WithMultipart().
		WithFile("data", "../public/bear.mp4").
		WithFormField("token", token).
		WithFormField("title", "Bear").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishResp.Value("status_code").Number().Equal(0)

	publishListResp := e.GET("/douyin/publish/list/").
		WithQuery("user_id", userId).WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishListResp.Value("status_code").Number().Equal(0)
	publishListResp.Value("video_list").Array().Length().Gt(0)

	for _, element := range publishListResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}
