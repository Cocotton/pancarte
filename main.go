package main

import "github.com/cocotton/pancarte/pancarte"

func main() {
	p := pancarte.Init("localhost")
	p.Run(":8080")
}
