package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	defer cancel()

	request, error := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if error != nil {
		panic(error)
	}

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	println(string(body))
}
