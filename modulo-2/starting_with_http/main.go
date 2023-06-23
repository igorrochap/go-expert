package main

import "net/http"

func main() {
	http.HandleFunc("/", searchByCEP)
	http.ListenAndServe(":8080", nil)
}

func searchByCEP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World!"))
}
