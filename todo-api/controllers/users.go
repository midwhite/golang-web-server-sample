package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/midwhite/golang-web-server-sample/todo-api/userservice"
)

type GetUserDetailResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func GetUserDetail(w http.ResponseWriter, req *http.Request) {
	user := userservice.GetUserDetail("1")
	data := GetUserDetailResponse{ID: user.Id, Name: user.Name, Age: user.Age}

	body, _ := json.Marshal(data)
	w.Write(body)
}
