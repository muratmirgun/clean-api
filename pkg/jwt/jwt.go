package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// MakeToken helper to make a token from user
func MakeToken(id string) (string, error) {

	secret := "clean-api-secret"
	claims := Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "clean-api",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

// JWTGetUserFromToken  helper to retrieve the user from token
// Get user from token
func JWTGetUserFromToken(tokenString string) (*Claims, error) {
	secret := "clean-api-secret"

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Claims struct
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
