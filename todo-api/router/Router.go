package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/midwhite/golang-web-server-sample/todo-api/controllers"
)

func StartServer() {
	http.HandleFunc("/todos/", controllers.GetTodos)
	http.HandleFunc("/users/1", controllers.GetUserDetail)

	fmt.Println("router is set up.")

	log.Fatal(
		http.ListenAndServe(":3000", nil),
	)
}
