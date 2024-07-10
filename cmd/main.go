package main

import "github.com/ratheeshkumar25/pkg/di"

func main() {
	server := di.Init()
	server.StartServer()

}