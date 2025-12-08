package models

type RegisterModel struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	First_name        string `json:"first_name"`
	Last_name         string `json:"last_name"`
	Keep_me_logged_in bool   `json:"keep_me_logged_in"`
}
