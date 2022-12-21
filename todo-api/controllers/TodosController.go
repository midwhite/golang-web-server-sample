package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/midwhite/golang-web-server-sample/todo-api/models"
)

func HandleTodos(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET", "":
		getTodoList(w, req)
	}
}

func getTodoList(w http.ResponseWriter, _ *http.Request) {
	todo := models.Todo{Id: "1", Title: "Todo Sample 1"}
	todos := []models.Todo{todo}
	response := map[string][]models.Todo{"todos": todos}

	body, err := json.Marshal(response)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		w.Write(body)
	} else {
		w.WriteHeader(400)

		res := models.ErrorResponse{Message: "unexpected error occured."}
		errorRes, _ := json.Marshal(res)

		w.Write(errorRes)
	}
}
