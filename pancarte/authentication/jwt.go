package authentication

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cocotton/pancarte/pancarte/response"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims struct contains the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Key bla
type Key int

// MyKey bla
const MyKey Key = 0

// SetToken bla
func SetToken(w http.ResponseWriter, r *http.Request, username string) {
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

	signedToken, _ := token.SignedString([]byte("secret"))

	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(w, &cookie)

}

// Validate bla
func Validate(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			response.ErrorWithText(w, "Auth cookie not found", http.StatusForbidden)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			a := *claims
			fmt.Println(a)
			ctx := context.WithValue(r.Context(), MyKey, *claims)

			protectedPage(w, r.WithContext(ctx))
		} else {
			http.NotFound(w, r)
			return
		}
	})
}

// ProtectedPage bla
func ProtectedPage(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(MyKey).(Claims)
	if !ok {
		http.NotFound(w, r)
	}
	response.SuccessWithJSON(w, []byte("Hello: "+claims.Username), http.StatusOK)
}

// Logout deletes the cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	return
}
