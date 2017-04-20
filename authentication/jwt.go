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

// CreateJWTToken will create the JWT and add it in a cookie
func CreateJWTToken(w http.ResponseWriter, r *http.Request, username string, jwtSecret string) {
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

	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(w, &cookie)
}
