package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTestRelation(b *testing.B) {
	e := newExpect(b)
	userIdA, tokenA := getTestUserToken(testUserA, e)
	userIdB, tokenB := getTestUserToken(testUserB, e)


	relationResp := e.POST("/douyin/relation/action/").
	WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).
	WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
relationResp.Value("status_code").Number().Equal(0)

followListResp := e.GET("/douyin/relation/follow/list/").
	WithQuery("token", tokenA).WithQuery("user_id", userIdA).
	WithFormField("token", tokenA).WithFormField("user_id", userIdA).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
followListResp.Value("status_code").Number().Equal(0)

containTestUserB := false
for _, element := range followListResp.Value("user_list").Array().Iter() {
	user := element.Object()
	user.ContainsKey("id")
	if int(user.Value("id").Number().Raw()) == userIdB {
		containTestUserB = true
	}
}
assert.True(b, containTestUserB, "Follow test user failed")

followerListResp := e.GET("/douyin/relation/follower/list/").
	WithQuery("token", tokenB).WithQuery("user_id", userIdB).
	WithFormField("token", tokenB).WithFormField("user_id", userIdB).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
followerListResp.Value("status_code").Number().Equal(0)

containTestUserA := false
for _, element := range followerListResp.Value("user_list").Array().Iter() {
	user := element.Object()
	user.ContainsKey("id")
	if int(user.Value("id").Number().Raw()) == userIdA {
		containTestUserA = true
	}
}
assert.True(b, containTestUserA, "Follower test user failed")

relationResp = e.POST("/douyin/relation/action/").
	WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).WithQuery("action_type", 1).
	WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).WithFormField("action_type", 1).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
relationResp.Value("status_code").Number().Equal(0)

friendResp := e.GET("/douyin/relation/friend/list/").
	WithQuery("token", tokenA).WithQuery("user_id", userIdA).
	WithFormField("token", tokenA).WithFormField("user_id", userIdA).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
friendResp.Value("status_code").Number().Equal(0)

containTestUserB = false
for _, element := range friendResp.Value("user_list").Array().Iter() {
	user := element.Object()
	user.ContainsKey("id")
	if int(user.Value("id").Number().Raw()) == userIdB {
		containTestUserB = true
	}
}
assert.True(b, containTestUserB, "Friend test user failed")

relationResp = e.POST("/douyin/relation/action/").
	WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 2).
	WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 2).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
relationResp.Value("status_code").Number().Equal(0)

friendResp = e.GET("/douyin/relation/friend/list/").
	WithQuery("token", tokenA).WithQuery("user_id", userIdA).
	WithFormField("token", tokenA).WithFormField("user_id", userIdA).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
friendResp.Value("status_code").Number().Equal(0)

containTestUserB = false
for _, element := range friendResp.Value("user_list").Array().Iter() {
	user := element.Object()
	user.ContainsKey("id")
	if int(user.Value("id").Number().Raw()) == userIdB {
		containTestUserB = true
	}
}
assert.False(b, containTestUserB, "Cancle follow test user failed")

friendResp = e.GET("/douyin/relation/friend/list/").
	WithQuery("token", tokenB).WithQuery("user_id", userIdB).
	WithFormField("token", tokenB).WithFormField("user_id", userIdB).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
friendResp.Value("status_code").Number().Equal(0)

containTestUserA = false
for _, element := range friendResp.Value("user_list").Array().Iter() {
	user := element.Object()
	user.ContainsKey("id")
	if int(user.Value("id").Number().Raw()) == userIdA {
		containTestUserA = true
	}
}
assert.False(b, containTestUserA, "Cancle follow test user failed")

relationResp = e.POST("/douyin/relation/action/").
	WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).WithQuery("action_type", 2).
	WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).WithFormField("action_type", 2).
	Expect().
	Status(http.StatusOK).
	JSON().Object()
relationResp.Value("status_code").Number().Equal(0)

}

func BenchmarkTestChat(b *testing.B) {
	e := newExpect(b)

	userIdA, tokenA := getTestUserToken(testUserA, e)
	userIdB, tokenB := getTestUserToken(testUserB, e)

	messageResp := e.POST("/douyin/message/action/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).WithQuery("action_type", 1).WithQuery("content", "Send to UserB").
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).WithFormField("action_type", 1).WithQuery("content", "Send to UserB").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	messageResp.Value("status_code").Number().Equal(0)

	chatResp := e.GET("/douyin/message/chat/").
		WithQuery("token", tokenA).WithQuery("to_user_id", userIdB).
		WithFormField("token", tokenA).WithFormField("to_user_id", userIdB).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	chatResp.Value("status_code").Number().Equal(0)
	chatResp.Value("message_list").Array().Length().Gt(0)

	chatResp = e.GET("/douyin/message/chat/").
		WithQuery("token", tokenB).WithQuery("to_user_id", userIdA).
		WithFormField("token", tokenB).WithFormField("to_user_id", userIdA).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	chatResp.Value("status_code").Number().Equal(0)
	chatResp.Value("message_list").Array().Length().Gt(0)
}
