package router

import (
	"log"
	"net/http"

	"github.com/midwhite/golang-web-server-sample/todo-api/controllers"
)

func StartServer() {
	http.HandleFunc("/todos", controllers.GetTodos)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
