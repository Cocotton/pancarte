package main

import (
	"net/http"
	"os"

	"github.com/cocotton/pancarte/pancarte"
)

func main() {
	p := pancarte.Pancarte{}

	p.InitDB("localhost", "pancarte")
	p.InitRouter()
	p.SetJWTSecret(loadJWTSecret())
	http.ListenAndServe(":8080", p.Router)
}

func loadJWTSecret() string {
	return os.Getenv("PANCARTE_JWT_SECRET")
}
