package main

import (
	"fmt"

	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/router"
)

func main() {
	close := db.Setup()
	defer close()

	fmt.Println("HTTP server is starting...")
	router.StartServer()
}
