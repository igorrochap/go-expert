package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil)
	if err != nil {
		panic(err)
	}
	result, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	io.Copy(os.Stdout, result.Body)
}
