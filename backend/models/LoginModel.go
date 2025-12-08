package models

type LoginModel struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	Keep_me_logged_in bool   `json:"keep_me_logged_in"`
}
