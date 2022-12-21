package utils

import (
	"net/http"
	"strconv"
)

func ReadRequestBody(req *http.Request) ([]byte, error) {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))

	if err != nil {
		return nil, err
	}

	body := make([]byte, length)
	req.Body.Read(body)

	return body, nil
}
