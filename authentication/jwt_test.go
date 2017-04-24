package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var expireCookie = time.Now().Add(time.Hour * 1)
var jwtSecret = "jwtSecret"
var validToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJUZXN0IElzc3VlciIsImlhdCI6MTQ5MjgwMTI4MCwiZXhwIjo0MTEyMDE3MjgwLCJhdWQiOiJ3d3cuZXhhbXBsZS5jb20iLCJzdWIiOiJ0ZXN0dXNlciIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ._r4WBZghrAZb2d5i_ur9nntIEtuydUDeVDJJozbUiR4"
var expiredToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJUZXN0IElzc3VlciIsImlhdCI6MTQ5MjYyODQ4MCwiZXhwIjo5NTYzNDU1MzMsImF1ZCI6Ind3dy5leGFtcGxlLmNvbSIsInN1YiI6InRlc3R1c2VyIiwidXNlcm5hbWUiOiJ0ZXN0dXNlciJ9.trtIuzdxasWJaFUwZJHsAgz-XBqi-YtJCBCepHrWevY"
var badSecretToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJUZXN0IElzc3VlciIsImlhdCI6MTQ5MjYyODQ4MCwiZXhwIjo5NTYzNDU1MzMsImF1ZCI6Ind3dy5leGFtcGxlLmNvbSIsInN1YiI6InRlc3R1c2VyIiwidXNlcm5hbWUiOiJ0ZXN0dXNlciJ9.pQJkJnbwTX8Hqj74KbO8EN7DpeP0JKTMK8pG8fMD72E"
var badSigningAlgoToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzb21lIGlzc3VlciIsInN1YiI6InNvbWUgc3ViamVjeSIsIm5iZiI6MTQ5MjgwMzM3MywiZXhwIjoxNDkyODA2OTczLCJpYXQiOjE0OTI4MDMzNzMsImp0aSI6InNvbWUganRpIiwidHlwIjoic29tZSB0eXBlIn0.ZUuo8p9CBS5xHmbX_5RWrkgP1HA4hoXE6bzM7w5y0ZQQnEFFgwM1F4wMXHFsGKnud7p_wSRYdcahpscBRPsL2JyjzAngwoyZ2CEUQK-PrfblCDV0pyADg4dMGmY-VCLQNGfT2q_PMIyui5X3-rI6MWQNk3Jcm9C-52euXL1DSrX_FQczXTcPZ0snAGdwOb3c05kEprWZ9ph9gUmcXKPYPye7rR7mCAKo0N2M7F3CjPXzr-qtVa8bjMeh9AWL4XJO3Ej1PcFdeWO2epmn6tRXAxi3-vdht8Xzuz_5v8ELQYjHRrbGlNz2t1jL51_rvaQHKgLIqMQ58UiBqcYFgDSAZg"
var invalidToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJUZXN0IElzc3VlciIsImlhdCI6MTQ5MjgwMTI4MCwiZXhwIjo0MTEyMDE3MjgwLCJhdWQiOiJ3d3cuZXhhbXBsZS5jb20iLCJzdWIiOiJ0ZXN0dXNlciIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ._r4WBZghrAZb2d5i_ur9nntIEtuydUDeVDJJozbUiR41"

func TestCreateJWTCookie(t *testing.T) {
	assert := assert.New(t)
	var CreateJWTCookieTests = []struct {
		username    string
		jwtSecret   string
		expectError bool
	}{
		{"user", "secret", false},
		{"", "secret", true},
		{"user", "", true},
	}

	for _, test := range CreateJWTCookieTests {
		actual, err := CreateJWTCookie(test.username, test.jwtSecret)
		if test.expectError {
			assert.Error(err, "Expected an error. None was returned.")
		} else {
			assert.IsType(&http.Cookie{}, actual, "Expected an *http.Cookie. Did not receive one.")
		}
	}
}

func TestGetJWT(t *testing.T) {
	assert := assert.New(t)

	var GetJWTtests = []struct {
		cookie      http.Cookie
		jwtSecret   string
		expectError bool
	}{
		{http.Cookie{Name: "Auth", Value: validToken, Expires: expireCookie, HttpOnly: true}, jwtSecret, false},
		{http.Cookie{Name: "Auth", Value: expiredToken, Expires: expireCookie, HttpOnly: true}, jwtSecret, true},
		{http.Cookie{Name: "Auth", Value: badSecretToken, Expires: expireCookie, HttpOnly: true}, jwtSecret, true},
		{http.Cookie{Name: "Auth", Value: badSigningAlgoToken, Expires: expireCookie, HttpOnly: true}, jwtSecret, true},
		{http.Cookie{Name: "Auth", Value: invalidToken, Expires: expireCookie, HttpOnly: true}, jwtSecret, true},
	}

	for _, test := range GetJWTtests {
		actual, err := GetJWT(test.cookie, test.jwtSecret)
		if test.expectError {
			assert.Error(err, "Expected an error. None was returned.")
		} else {
			assert.IsType(&jwt.Token{}, actual, "Expected a *jwt.Token. Did not receive one.")
		}
	}
}

func TestGetJWTClaims(t *testing.T) {
	assert := assert.New(t)
	cookie := http.Cookie{Name: "Auth", Value: validToken, Expires: expireCookie, HttpOnly: true} //CreateJWTCookie("Username", jwtSecret)
	validJWT, _ := GetJWT(cookie, jwtSecret)
	invalidJWT := *validJWT
	invalidJWT.Claims = jwt.StandardClaims{}

	var GetJWTClaimsTests = []struct {
		token       *jwt.Token
		expectError bool
	}{
		{validJWT, false},
		{&invalidJWT, true},
	}

	for _, test := range GetJWTClaimsTests {
		actual, err := GetJWTClaims(test.token)
		if test.expectError {
			assert.Error(err, "Expected an error. None was returned.")
		} else {
			assert.IsType(&Claims{}, actual, "Expected a *Claims. Did not receive one.")
		}
	}
}

func TestGetContextWithClaims(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	claims := Claims{
		"Username",
		jwt.StandardClaims{
			Issuer: "localhost:8080",
		},
	}

	actual := GetContextWithClaims(r, &claims)
	assert.NotEmpty(t, actual, "Got empty context.")
}
