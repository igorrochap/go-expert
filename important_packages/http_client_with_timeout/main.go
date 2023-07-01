package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Microsecond}
	response, error := client.Get("https://google.com")
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
