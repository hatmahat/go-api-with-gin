package main

type UserCredential struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
