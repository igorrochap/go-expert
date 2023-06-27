package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	template := template.Must(template.New("template.html").ParseFiles("template.html"))
	error := template.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 80},
		{"Python", 20},
		{"PHP", 30},
	})
	if error != nil {
		panic(error)
	}
}
