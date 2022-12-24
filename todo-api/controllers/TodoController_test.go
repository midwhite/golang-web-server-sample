package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_URL", "postgres://postgres@todo-api-db:5432/todo_api_test?sslmode=disable")
	db.Setup()
	db.Migrate()

	status := m.Run()

	db.Rollback()
	os.Exit(status)
}

func TestTodoControllerIndex(t *testing.T) {
	t.Run("when todo does not exist", func(t *testing.T) {
		t.Run("responds empty array", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/todos", nil)
			got := httptest.NewRecorder()

			HandleTodos(got, req)

			if got.Code != http.StatusOK {
				t.Errorf("expected %d, but got %d", http.StatusOK, got.Code)
			}

			data := new(GetTodoListResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if len(data.Todos) != 0 {
				t.Errorf("expected todo list to be empty, but got %+v", data.Todos)
			}
		})
	})

	t.Run("when some todos exist", func(t *testing.T) {
		todo1 := models.Todo{Title: "Todo 1", CreatedAt: time.Now()}
		todo1.Insert(context.Background(), db.Conn, boil.Infer())
		todo2 := models.Todo{Title: "Todo 2", CreatedAt: time.Now()}
		todo2.Insert(context.Background(), db.Conn, boil.Infer())

		t.Run("responds todo list", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/todos", nil)
			got := httptest.NewRecorder()

			HandleTodos(got, req)

			if got.Code != http.StatusOK {
				t.Errorf("expected %d, but got %d", http.StatusOK, got.Code)
			}

			data := new(GetTodoListResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if len(data.Todos) != 2 {
				t.Errorf("expected todo list to have 2 items, but got %+v", data.Todos)
			}

			data1, data2 := data.Todos[0], data.Todos[1]

			if data1.Title != "Todo 1" || data2.Title != "Todo 2" {
				t.Errorf("expected todo has correct title, but got %+v, %+v", data1, data2)
			}
		})
	})
}
