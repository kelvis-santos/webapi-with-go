package main

import "github.com/kelvis-santos/webapi-with-go/server"

func main() {
	server := server.NewServer()

	server.Run()
}
