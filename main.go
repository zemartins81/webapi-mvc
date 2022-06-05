package main

import (
	"github.com/jcmartins81/webapi-with-go/database"
	"github.com/jcmartins81/webapi-with-go/server"
)

func main() {

	database.StartDB()

	newServer := server.NewServer()

	newServer.Run()
}
