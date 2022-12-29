package router

import (
	"log"
	"net/http"

	"github.com/midwhite/golang-web-server-sample/todo-api/controllers"
)

func StartServer() {
	http.HandleFunc("/todos", controllers.HandleTodos)
	http.HandleFunc("/todos/", controllers.HandleTodoDetail)
	http.HandleFunc("/users/1", controllers.GetUserDetail)

	log.Fatal(
		http.ListenAndServe(":3000", nil),
	)
}
