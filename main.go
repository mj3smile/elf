package main

import (
	"elf/route"
)

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":9000"))
}
