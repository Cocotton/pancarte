package main

import "github.com/cocotton/pancarte/pancarte"

func main() {
	var p pancarte.Pancarte
	p.InitDB("localhost")
	p.InitRouter()
	p.Run(":8080")
}
