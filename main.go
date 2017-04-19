package main

import "github.com/cocotton/pancarte/pancarte"

func main() {
	p := pancarte.Pancarte{}

	p.InitDB("localhost", "pancarte")
	p.InitRouter()
}
