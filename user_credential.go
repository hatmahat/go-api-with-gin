package main

// type UserCredential struct {
// 	Username string `form:"username" binding:"required"`
// 	Password string `form:"password" binding:"required"`
// }

type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"user_password"`
}

type AuthHandler struct {
	Authorizationheader string `header:"Authorization"`
}
