package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jorgeAM/echo-api/models"
)

var key = os.Getenv("JWT_KEY")

// GenerateJWT generates jwt
func GenerateJWT(user *models.User) (*models.Token, error) {
	c := &models.Claim{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "todo app",
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	stringToken, err := t.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &models.Token{
		Token: stringToken,
		User:  user,
	}, nil
}
