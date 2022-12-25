package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/models"
	"github.com/midwhite/golang-web-server-sample/todo-api/serializers"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HandleTodos(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		createTodo(w, req)
	case "GET", "":
		getTodoList(w, req)
	default:
		data := serializers.ErrorResponse{Message: "Not Found"}
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
		data := serializers.ErrorResponse{Message: "ID is not set."}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
		return
	}

	switch req.Method {
	case "GET", "":
		getTodoDetail(w, req, todoId)
	case "PUT":
		updateTodo(w, req, todoId)
	case "DELETE":
		deleteTodo(w, req, todoId)
	default:
		data := serializers.ErrorResponse{Message: "No route matches."}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
	}
}

type CreateTodoParams struct {
	Title string `json:"title"`
}

func createTodo(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := io.ReadAll(req.Body)
	params := new(CreateTodoParams)
	json.Unmarshal(reqBody, params)

	todo := models.Todo{Title: params.Title, CreatedAt: time.Now()}
	todo.Insert(context.Background(), db.Conn, boil.Infer())
	body, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

type GetTodoListResponse struct {
	Todos []*models.Todo `json:"todos"`
}

func getTodoList(w http.ResponseWriter, _ *http.Request) {
	todos, _ := models.Todos(qm.OrderBy("created_at")).All(context.Background(), db.Conn)
	if todos == nil {
		todos = make([]*models.Todo, 0)
	}
	response := GetTodoListResponse{Todos: todos}
	body, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func getTodoDetail(w http.ResponseWriter, _ *http.Request, todoId string) {
	todo, err := models.FindTodo(context.Background(), db.Conn, todoId, "id", "title")

	if err != nil {
		data := serializers.ErrorResponse{Message: err.Error()}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
	} else {
		body, _ := json.Marshal(todo)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

type UpdateTodoParams struct {
	Title string `json:"title"`
}

func updateTodo(w http.ResponseWriter, req *http.Request, todoId string) {
	todo, err := models.FindTodo(context.Background(), db.Conn, todoId, "id", "title", "created_at")

	if err != nil {
		data := serializers.ErrorResponse{Message: err.Error()}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
	} else {
		reqBody, _ := io.ReadAll(req.Body)

		params := new(UpdateTodoParams)
		json.Unmarshal(reqBody, params)

		todo.Title = params.Title
		todo.Update(context.Background(), db.Conn, boil.Infer())

		body, _ := json.Marshal(todo)
		w.Write(body)
	}
}

func deleteTodo(w http.ResponseWriter, _ *http.Request, todoId string) {
	todo, err := models.FindTodo(context.Background(), db.Conn, todoId, "id")

	if err != nil {
		data := serializers.ErrorResponse{Message: err.Error()}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusNotFound)
		w.Write(body)
	} else {
		todo.Delete(context.Background(), db.Conn)

		data := map[string]bool{
			"success": true,
		}
		body, _ := json.Marshal(data)

		w.Write(body)
	}
}
