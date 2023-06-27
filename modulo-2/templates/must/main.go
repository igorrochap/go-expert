package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go for noobs", 40}
	template := template.Must(template.New("course_template.html").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))
	error := template.Execute(os.Stdout, course)
	if error != nil {
		panic(error)
	}
}
