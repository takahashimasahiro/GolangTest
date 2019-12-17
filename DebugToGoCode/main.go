package main

import (
	"net/http"
)

func main() {
	// net/httpの中身を見る用
	_, err := http.Get("http://example.com/")
	if err != nil {
		return
	}
}