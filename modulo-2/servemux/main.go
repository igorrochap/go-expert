package main

import "net/http"

type Blog struct {
	title string
}

func (blog Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(blog.title))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", Blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)

	anotherMux := http.NewServeMux()
	anotherMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello my another Mux!"))
	})
	http.ListenAndServe(":8081", anotherMux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
