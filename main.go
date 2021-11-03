package main

import (
	"web-api-go/database"
	"web-api-go/server"
)

func main() {

	database.StartDB()
	server := server.NewServer()

	server.Run()
}
