package controllers

import (
	"io"
	"net/http"
)

func GetUserDetail(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "{\"user\":{\"id\":1,\"name\":\"midwhite\"}}")
}
