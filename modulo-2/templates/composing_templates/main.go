package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	http.HandleFunc("/", parseCoursesForTemplate)
	http.ListenAndServe(":8181", nil)
}

func parseCoursesForTemplate(writer http.ResponseWriter, request *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	template := template.Must(template.New("content.html").ParseFiles(templates...))
	error := template.Execute(writer, Courses{
		{"Go", 40},
		{"Java", 80},
		{"Python", 20},
		{"PHP", 30},
	})

	if error != nil {
		panic(error)
	}
}
