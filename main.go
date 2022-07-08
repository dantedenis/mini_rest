package main

import (
	"wb_test/web"
)

func main() {
	s := web.NewServer("localhost", "8080")
	s.Run()
}