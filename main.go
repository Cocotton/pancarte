package main

import (
	"net/http"

	"github.com/cocotton/pancarte/pancarte"
)

func main() {
	p := pancarte.Pancarte{}

	p.InitDB("localhost", "pancarte")
	p.InitRouter()
	http.ListenAndServe(":8080", p.Router)
}
