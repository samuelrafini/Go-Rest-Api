package model

type JWT struct {
	Token string `json:"token"`
}

type Student struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}