package main

import (
	"fmt"

	"github.com/midwhite/golang-web-server-sample/todo-api/router"
)

func main() {
	fmt.Println("HTTP server is starting...")

	router.StartServer()
}
