package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/midwhite/golang-web-server-sample/todo-api/models"
)

func HandleTodos(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET", "":
		getTodoList(w, req)
	default:
		data := models.ErrorResponse{Message: "Not Found"}
		body, _ := json.Marshal(data)

		w.WriteHeader(404)
		w.Write(body)
	}
}

func HandleTodoDetail(w http.ResponseWriter, req *http.Request) {
	pathname := strings.TrimPrefix(req.URL.Path, "/todos/")
	paths := regexp.MustCompile("[/?]").Split(pathname, -1)
	todoId := paths[0]

	if todoId == "" {
		data := models.ErrorResponse{Message: "ID is not set."}
		body, _ := json.Marshal(data)

		w.WriteHeader(404)
		w.Write(body)
		return
	}

	switch req.Method {
	case "GET", "":
		getTodoDetail(w, req, todoId)
	}
}

type GetTodoListResponse struct {
	Todos []models.Todo `json:"todos"`
}

func getTodoList(w http.ResponseWriter, _ *http.Request) {
	todo := models.Todo{Id: "1", Title: "Todo Sample 1"}
	todos := []models.Todo{todo}
	response := GetTodoListResponse{Todos: todos}

	body, err := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")

	if err == nil {
		w.WriteHeader(200)
		w.Write(body)
	} else {
		w.WriteHeader(400)

		res := models.ErrorResponse{Message: "unexpected error occured."}
		errorRes, _ := json.Marshal(res)

		w.Write(errorRes)
	}
}

func getTodoDetail(w http.ResponseWriter, _ *http.Request, todoId string) {
	todo := models.Todo{Id: todoId, Title: "Sample Todo"}
	body, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)
}
