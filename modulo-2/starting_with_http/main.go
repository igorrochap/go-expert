package main

import "net/http"

func main() {
	http.HandleFunc("/", searchByCEP)
	http.ListenAndServe(":8080", nil)
}

func searchByCEP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	cep := request.URL.Query().Get("cep")
	if cep == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hello, World!"))
}
