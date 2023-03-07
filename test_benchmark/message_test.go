package test

import (
	"net/http"
	"testing"
)

func BenchmarkChat(b *testing.B) {
	e := newExpect(b)
	token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzgyNDMwNDEsImlkIjoxNCwib3JpZ19pYXQiOjE2NzgxNTY2NDF9.MiShRh6jvx9XQOxoVA1YqB0I8UM-W9DuGO3-2Za7EwI"
	//token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyOTQxMjgsImlkIjoxMzAsIm9yaWdfaWF0IjoxNjc3MjA3NzI4fQ.Xy2iBkUH-JLDIGzcpYBmcYe01yxiP9DQ-hAQAnjxNig"
	_ = e.POST("/douyin/message/action/").
		WithQuery("token", token1).WithQuery("to_user_id", "130").WithQuery("action_type", 1).WithQuery("content", "Send to UserB").
		WithFormField("token", token1).WithFormField("to_user_id", "130").WithFormField("action_type", 1).WithQuery("content", "Send to UserB").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}
func BenchmarkGetMessageList(b *testing.B) {
	e := newExpect(b)
	token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzgyNDMwNDEsImlkIjoxNCwib3JpZ19pYXQiOjE2NzgxNTY2NDF9.MiShRh6jvx9XQOxoVA1YqB0I8UM-W9DuGO3-2Za7EwI"
	//token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyOTQxMjgsImlkIjoxMzAsIm9yaWdfaWF0IjoxNjc3MjA3NzI4fQ.Xy2iBkUH-JLDIGzcpYBmcYe01yxiP9DQ-hAQAnjxNig"
	_ = e.GET("/douyin/message/chat/").
		WithQuery("token", token1).WithQuery("to_user_id", "130").
		WithFormField("token", token1).WithFormField("to_user_id", "130").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}
