package main

import (
	"fmt"

	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/router"
)

func main() {
	db.Setup()
	defer db.Close()

	fmt.Println("HTTP server is starting...")
	router.StartServer()
}
