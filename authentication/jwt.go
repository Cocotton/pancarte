package authentication

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims contains the informations of the logged in user
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateJWTCookie will create the JWT and return cookie containing it
func CreateJWTCookie(username string, jwtSecret string) *http.Cookie {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	expireCookie := time.Now().Add(time.Hour * 1)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(jwtSecret))

	return &http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
}
