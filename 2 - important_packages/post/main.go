package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	payload := bytes.NewBuffer([]byte(`"name": "Igor"`))
	response, error := client.Post("https://google.com", "application/json", payload)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	io.CopyBuffer(os.Stdout, response.Body, nil)
}
