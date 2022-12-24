package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/midwhite/golang-web-server-sample/todo-api/models"
	"github.com/midwhite/golang-web-server-sample/todo-api/utils"
)

func HandleTodos(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		createTodo(w, req)
	case "GET", "":
		getTodoList(w, req)
	default:
		data := models.ErrorResponse{Message: "Not Found"}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
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

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
		return
	}

	switch req.Method {
	case "GET", "":
		getTodoDetail(w, req, todoId)
	default:
		data := models.ErrorResponse{Message: "No route matches."}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
	}
}

type CreateTodoParams struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func createTodo(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := utils.ReadRequestBody(req)
	params := new(CreateTodoParams)
	json.Unmarshal(reqBody, params)

	todo := models.Todo{ID: strconv.Itoa(params.Id), Title: params.Title}
	body, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

type GetTodoListResponse struct {
	Todos []models.Todo `json:"todos"`
}

func getTodoList(w http.ResponseWriter, _ *http.Request) {
	todo := models.Todo{ID: "1", Title: "Todo Sample 1", CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	todos := []models.Todo{todo}
	response := GetTodoListResponse{Todos: todos}

	body, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func getTodoDetail(w http.ResponseWriter, _ *http.Request, todoId string) {
	todo := models.Todo{ID: todoId, Title: "Sample Todo", CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	body, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
