package authentication

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims contains the informations of the logged in user
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type key int

const jwtClaimsKey key = 0

// CreateJWTCookie will create the JWT and return cookie containing it
func CreateJWTCookie(username string, jwtSecret string) *http.Cookie {
	expireToken := time.Now().Add(time.Hour * 72).Unix()
	expireCookie := time.Now().Add(time.Hour * 72)

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

// GetJWT validates and returns a JWT
func GetJWT(cookie http.Cookie, jwtSecret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("The JWT fetched from the provided cookie is not valid")
	}

	return token, nil
}

// GetJWTClaims validates and returns the claims from a JWT
func GetJWTClaims(token *jwt.Token) (*Claims, error) {
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("The claims found in the JWT are not valid")
	}

	return claims, nil
}

// GetContextWithClaims adds the claims with the ContextKey key to an HTTP context and returns it
func GetContextWithClaims(r *http.Request, claims *Claims) context.Context {
	return context.WithValue(r.Context(), jwtClaimsKey, *claims)
}
