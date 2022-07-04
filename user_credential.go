package main

type UserCredential struct {
	// Username string `form:"use_name" binding:"required"`
	// Password string `form:"user_password" binding:"required"`

	Username string `json:"user_name"`
	Password string `json:"user_password"`
}
