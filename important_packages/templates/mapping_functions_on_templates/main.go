package main

import (
	"html/template"
	"net/http"
	"strings"
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

	myTemplate := template.New("content.html")
	myTemplate.Funcs(template.FuncMap{"ToUpper": ToUpper})
	myTemplate = template.Must(myTemplate.ParseFiles(templates...))

	error := myTemplate.Execute(writer, Courses{
		{"Go", 40},
		{"Java", 80},
		{"Python", 20},
		{"PHP", 30},
	})

	if error != nil {
		panic(error)
	}
}

func ToUpper(element string) string {
	return strings.ToUpper(element)
}
