package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/midwhite/golang-web-server-sample/todo-api/models"
)

func GetTodos(w http.ResponseWriter, _ *http.Request) {
	todo := models.Todo{Id: "1", Title: "Todo Sample"}
	data, err := json.Marshal(todo)

	if err == nil {
		w.Write(data)
	} else {
		io.WriteString(w, "error!")
	}
}
