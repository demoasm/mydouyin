package apimodel

type CreateUserRequest struct {
	Username string `json:"username" query:"username" vd:"len($) > 0"`
	Password string `json:"password" query:"password" vd:"len($) > 0"`
}

type CheckUserRequest struct {
	Username string `json:"username" query:"username" vd:"len($) > 0"`
	Password string `json:"password" query:"password" vd:"len($) > 0"`
}

type GetUserRequest struct {
	UserID string `json:"user_id" query:"user_id"`
	Token  string `json:"token" query:"token"`
}

