package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func BenchmarkRegister(b *testing.B) {
	e := newExpect(b)
	rand.Seed(time.Now().UnixNano())
	registerValue := fmt.Sprintf("douyin%d", rand.Intn(65536))
	_ = e.POST("/douyin/user/register/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	// registerResp.Value("status_code").Number().Equal(0)
	// registerResp.Value("user_id").Number().Gt(0)
	// registerResp.Value("token").String().Length().Gt(0)
}

func BenchmarkLogin(b *testing.B) {
	e := newExpect(b)
	_ = e.POST("/douyin/user/login/").
		WithQuery("username", "dousheng").WithQuery("password", "123456").
		WithFormField("username", "douysheng").WithFormField("password", "123456").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	// loginResp.Value("status_code").Number().Equal(0)
	// loginResp.Value("user_id").Number().Gt(0)
	// loginResp.Value("token").String().Length().Gt(0)
}

func BenchmarkGetUser(b *testing.B) {
	e := newExpect(b)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODk4ODIsImlkIjoxMjEsIm9yaWdfaWF0IjoxNjc3MjAzNDgyfQ.Lkuj8M7p7P1Zg-XqFpNpff1NBrioEHeBEZGVFHbUmkQ"
	user_id := "121"
	_ = e.GET("/douyin/user/").
		WithQuery("token", token).
		WithQuery("user_id", user_id).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	// userResp.Value("status_code").Number().Equal(0)
	// userInfo := userResp.Value("user").Object()
	// userInfo.NotEmpty()
	// userInfo.Value("id").Number().Gt(0)
	// userInfo.Value("name").String().Length().Gt(0)

}
