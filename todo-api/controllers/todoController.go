package controllers

import (
	"io"
	"net/http"
)

func GetTodos(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "{\"todos\":[]}")
}
