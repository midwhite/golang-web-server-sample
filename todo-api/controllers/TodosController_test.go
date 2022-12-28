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

	"github.com/gofrs/uuid"
	"github.com/midwhite/golang-web-server-sample/todo-api/db"
	"github.com/midwhite/golang-web-server-sample/todo-api/models"
	"github.com/midwhite/golang-web-server-sample/todo-api/serializers"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_URL", "postgres://postgres@todo-api-db:5432/todo_api_test?sslmode=disable")
	close := db.Setup()
	defer close()

	db.Reset()

	status := m.Run()

	db.Rollback()
	os.Exit(status)
}

func TestTodosControllerCreate(t *testing.T) {
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

			if got.Code != http.StatusBadRequest {
				t.Errorf("expected %d, but got %d", http.StatusBadRequest, got.Code)
			}

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

			if got.Code != http.StatusCreated {
				t.Errorf("expected %d, but got %d", http.StatusCreated, got.Code)
			}

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

func TestTodosControllerIndex(t *testing.T) {
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

func TestTodosControllerShow(t *testing.T) {
	t.Run("when ID is wrong", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		t.Run("responds correct error message", func(t *testing.T) {
			id, _ := uuid.NewV4()
			req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/todos/"+id.String(), nil)
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			data := new(serializers.ErrorResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if got.Code != http.StatusNotFound {
				t.Errorf("expected %d, but got %d", http.StatusNotFound, got.Code)
			}

			if data.Message != "todo is not found." {
				t.Errorf("message is expected to be '%v', but got %v", "todo is not found.", data.Message)
			}
		})
	})

	t.Run("when ID is correct", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		todo := models.Todo{Title: "A sample todo"}
		todo.Insert(context.Background(), db.Conn, boil.Infer())

		t.Run("responds todo detail", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/todos/"+todo.ID, nil)
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			if got.Code != http.StatusOK {
				t.Errorf("expected %d, but got %d", http.StatusOK, got.Code)
			}

			data := new(models.Todo)
			json.Unmarshal(got.Body.Bytes(), data)

			if data.Title != "A sample todo" {
				t.Errorf("expected %v, but got %v", "A sample todo", data.Title)
			}
		})
	})
}

func TestTodosControllerUpdate(t *testing.T) {
	t.Run("when ID is wrong", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		t.Run("responds correct error message", func(t *testing.T) {
			id, _ := uuid.NewV4()

			reqBody := UpdateTodoParams{Title: ""}
			encodedReqBody, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPut, "http://localhost:3000/todos/"+id.String(), bytes.NewBuffer(encodedReqBody))
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			if got.Code != http.StatusNotFound {
				t.Errorf("expected %d, but got %d", http.StatusNotFound, got.Code)
			}

			data := new(serializers.ErrorResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if data.Message != "todo is not found." {
				t.Errorf("expected %+v, but got %+v", "todo is not found.", data.Message)
			}
		})
	})

	t.Run("when title is blank", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		todo := models.Todo{Title: "Old Title"}
		todo.Insert(context.Background(), db.Conn, boil.Infer())

		t.Run("responds correct error message", func(t *testing.T) {
			reqBody := UpdateTodoParams{Title: ""}
			encodedReqBody, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPut, "http://localhost:3000/todos/"+todo.ID, bytes.NewBuffer(encodedReqBody))
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			if got.Code != http.StatusBadRequest {
				t.Errorf("expected %d, but got %d", http.StatusBadRequest, got.Code)
			}

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

		todo := models.Todo{Title: "Old Title"}
		todo.Insert(context.Background(), db.Conn, boil.Infer())

		t.Run("updates title of specified todo", func(t *testing.T) {
			reqBody := UpdateTodoParams{Title: "New Title"}
			encodedReqBody, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPut, "http://localhost:3000/todos/"+todo.ID, bytes.NewBuffer(encodedReqBody))
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			if got.Code != http.StatusOK {
				t.Errorf("expected %d, but got %d", http.StatusOK, got.Code)
			}

			data := new(models.Todo)
			json.Unmarshal(got.Body.Bytes(), data)

			if data.Title != "New Title" {
				t.Errorf("expected %+v, but got %+v", "New Title", data.Title)
			}
		})
	})
}

func TestTodosControllerDestroy(t *testing.T) {
	t.Run("when ID is wrong", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		t.Run("responds correct error message", func(t *testing.T) {
			id, _ := uuid.NewV4()
			req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/todos/"+id.String(), nil)
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			if got.Code != http.StatusNotFound {
				t.Errorf("expected %d, but got %d", http.StatusNotFound, got.Code)
			}

			data := new(serializers.ErrorResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if data.Message != "todo is not found." {
				t.Errorf("expected %+v, but got %+v", "todo is not found.", data.Message)
			}
		})
	})

	t.Run("when ID is correct", func(t *testing.T) {
		t.Cleanup(func() {
			db.Reset()
		})

		todo := models.Todo{Title: "The Destroyer"}
		todo.Insert(context.Background(), db.Conn, boil.Infer())

		t.Run("deletes specified todo successfully", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/todos/"+todo.ID, nil)
			got := httptest.NewRecorder()

			HandleTodoDetail(got, req)

			if got.Code != http.StatusOK {
				t.Errorf("expected %d, but got %d", http.StatusOK, got.Code)
			}

			data := new(DeleteTodoResponse)
			json.Unmarshal(got.Body.Bytes(), data)

			if !data.Success {
				t.Errorf("expected %+v, but got %+v", true, data.Success)
			}
		})
	})
}
