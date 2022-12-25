package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/models"
	"github.com/midwhite/golang-web-server-sample/todo-api/serializers"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_URL", "postgres://postgres@todo-api-db:5432/todo_api_test?sslmode=disable")
	db.Setup()
	db.Reset()

	status := m.Run()

	db.Rollback()
	os.Exit(status)
}

func TestTodoControllerIndex(t *testing.T) {
	t.Run("when todo does not exist", func(t *testing.T) {
		t.Run("responds empty array", func(t *testing.T) {
			t.Cleanup(func() {
				db.Reset()
			})

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
		t.Cleanup(func() {
			db.Reset()
		})

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

func TestTodoControllerCreate(t *testing.T) {
	t.Run("when title is blank", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		t.Run("responds correct error message", func(t *testing.T) {
			reqBody := CreateTodoParams{Title: ""}
			encodedReqBody, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/todos", bytes.NewBuffer(encodedReqBody))
			got := httptest.NewRecorder()

			HandleTodos(got, req)

			data := new(serializers.ErrorResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if data.Message != "件名は必須です。" {
				t.Errorf("expected %+v, but got %+v", "件名は必須です。", data.Message)
			}
		})
	})

	t.Run("when parameters are valid", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		t.Run("creates new todo", func(t *testing.T) {
			prevCount, _ := models.Todos().Count(context.Background(), db.Conn)

			reqBody := CreateTodoParams{Title: "New Todo Title"}
			encodedReqBody, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/todos", bytes.NewBuffer(encodedReqBody))
			got := httptest.NewRecorder()

			HandleTodos(got, req)

			currentCount, _ := models.Todos().Count(context.Background(), db.Conn)

			if prevCount+1 != currentCount {
				t.Errorf("expected %+v, but got %+v", prevCount+1, currentCount)
			}

			data := new(models.Todo)
			json.Unmarshal(got.Body.Bytes(), data)

			if data.Title != reqBody.Title {
				t.Errorf("expected %+v, but got %+v", reqBody.Title, data.Title)
			}
		})
	})
}
