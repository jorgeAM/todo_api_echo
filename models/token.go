package models

// Token model
type Token struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
