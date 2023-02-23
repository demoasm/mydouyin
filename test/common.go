package test

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

var serverAddr = "http://localhost:8080"
var testUserA = "douyinTestUserA"
var testUserB = "douyinTestUserB"

func newExpect(b *testing.B) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client:   http.DefaultClient,
		BaseURL:  serverAddr,
		Reporter: httpexpect.NewAssertReporter(b),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(b, true),
		},
	})
}



func getTestUserToken(user string, e *httpexpect.Expect) (int, string) {
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", user).WithQuery("password", user).
		WithFormField("username", user).WithFormField("password", user).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	userId := 0
	token := registerResp.Value("token").String().Raw()
	if len(token) == 0 {
		loginResp := e.POST("/douyin/user/login/").
			WithQuery("username", user).WithQuery("password", user).
			WithFormField("username", user).WithFormField("password", user).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		loginToken := loginResp.Value("token").String()
		loginToken.Length().Gt(0)
		token = loginToken.Raw()
		userId = int(loginResp.Value("user_id").Number().Raw())
	} else {
		userId = int(registerResp.Value("user_id").Number().Raw())
	}
	return userId, token
}
