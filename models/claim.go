package models

import "github.com/dgrijalva/jwt-go"

// Claim model
type Claim struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	jwt.StandardClaims
}
