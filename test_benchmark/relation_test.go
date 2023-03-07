package test

import (
	"net/http"
	"testing"
)

func BenchmarkFollowAct(b *testing.B) {
	e := newExpect(b)
	token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	//token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyOTQxMjgsImlkIjoxMzAsIm9yaWdfaWF0IjoxNjc3MjA3NzI4fQ.Xy2iBkUH-JLDIGzcpYBmcYe01yxiP9DQ-hAQAnjxNig"
	_ = e.POST("/douyin/relation/action/").
		WithQuery("token", token1).WithQuery("to_user_id", "130").WithQuery("action_type", 1).
		WithFormField("token", token1).WithFormField("to_user_id", "130").WithFormField("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	_ = e.POST("/douyin/relation/action/").
		WithQuery("token", token1).WithQuery("to_user_id", "130").WithQuery("action_type", 2).
		WithFormField("token", token1).WithFormField("to_user_id", "130").WithFormField("action_type", 2).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}

func BenchmarkGetFollowList(b *testing.B) {
	e := newExpect(b)
	token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	//token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyOTQxMjgsImlkIjoxMzAsIm9yaWdfaWF0IjoxNjc3MjA3NzI4fQ.Xy2iBkUH-JLDIGzcpYBmcYe01yxiP9DQ-hAQAnjxNig"
	_ = e.GET("/douyin/relation/follow/list/").
		WithQuery("token", token1).WithQuery("user_id", "121").
		WithFormField("token", token1).WithFormField("user_id", "121").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}

func BenchmarkGetFollowerList(b *testing.B) {
	e := newExpect(b)
	//token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyOTQxMjgsImlkIjoxMzAsIm9yaWdfaWF0IjoxNjc3MjA3NzI4fQ.Xy2iBkUH-JLDIGzcpYBmcYe01yxiP9DQ-hAQAnjxNig"
	_ = e.GET("/douyin/relation/follower/list/").
		WithQuery("token", token2).WithQuery("user_id", "130").
		WithFormField("token", token2).WithFormField("user_id", "130").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}

func BenchmarkGetFriendList(b *testing.B) {
	e := newExpect(b)
	token1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	//token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyOTQxMjgsImlkIjoxMzAsIm9yaWdfaWF0IjoxNjc3MjA3NzI4fQ.Xy2iBkUH-JLDIGzcpYBmcYe01yxiP9DQ-hAQAnjxNig"
	_ = e.GET("/douyin/relation/friend/list/").
		WithQuery("token", token1).WithQuery("user_id", "121").
		WithFormField("token", token1).WithFormField("user_id", "121").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
}
