package main

import (
	"io"
	"net/http"
)

func main() {
	request, error := http.Get("https://www.google.com.br")
	if error != nil {
		panic(error)
	}

	result, error := io.ReadAll(request.Body)
	if error != nil {
		panic(error)
	}

	println(string(result))

	request.Body.Close()
}
