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
	template := template.Must(template.New("template.html").ParseFiles("template.html"))
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
