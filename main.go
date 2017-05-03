package main

import (
	"net/http"
	"os"

	"github.com/cocotton/pancarte/pancarte"
)

func main() {
	p := pancarte.Pancarte{}

	p.InitDB(loadEnvVar("PANCARTE_DB_HOST"), loadEnvVar("PANCARTE_DB_NAME"))
	p.InitRouter()
	p.SetJWTSecret(loadEnvVar("PANCARTE_JWT_SECRET"))
	http.ListenAndServe(":"+loadEnvVar("PANCARTE_PORT"), p.Router)
}

func loadEnvVar(variable string) string {
	return os.Getenv(variable)
}
